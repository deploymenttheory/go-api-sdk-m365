// http_client.go
/* The `http_client` package provides a configurable HTTP client tailored for interacting with specific APIs.
It supports different authentication methods, including "bearer" and "oauth". The client is designed with a
focus on concurrency management, structured error handling, and flexible configuration options.
The package offers a default timeout, custom backoff strategies, dynamic rate limiting,
and detailed logging capabilities. The main `Client` structure encapsulates all necessary components,
like the baseURL, authentication details, and an embedded standard HTTP client. */
package http_client

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const DefaultTimeout = 10 * time.Second

// Client represents an HTTP client to interact with a specific API.
type Client struct {
	TenantID           string           // M365 tenant ID
	TenantName         string           // M365 tenant name
	AuthMethod         string           // Specifies the authentication method: "clientApp" or "clientCertificate"
	Token              string           // Authentication Token
	OverrideBaseDomain string           // Base domain override used when the default in the api handler isn't suitable
	OAuthCredentials   OAuthCredentials // ClientID / Client Secret
	Expiry             time.Time        // Expiry time set for the auth token
	httpClient         *http.Client
	config             Config
	logger             Logger
	ConcurrencyMgr     *ConcurrencyManager
	PerfMetrics        ClientPerformanceMetrics
}

// OAuthCredentials contains the client ID and client secret required for OAuth authentication.
type OAuthCredentials struct {
	ClientID           string
	ClientSecret       string
	CertificatePath    string
	CertificateKeyPath string
	CertThumbprint     string
}

// Config holds configuration options for the HTTP Client.
type Config struct {
	LogLevel                  LogLevel // Field for defining tiered logging level.
	MaxRetryAttempts          int      // Config item defines the max number of retry request attempts for retryable HTTP methods.
	EnableDynamicRateLimiting bool
	Logger                    Logger // Field for the packages initailzed logger
	MaxConcurrentRequests     int    // Field for defining the maximum number of concurrent requests allowed in the semaphore
	TokenLifespan             time.Duration
	TokenRefreshBufferPeriod  time.Duration
	TotalRetryDuration        time.Duration
}

// ClientPerformanceMetrics captures various metrics related to the client's
// interactions with the API, providing insights into its performance and behavior.
type ClientPerformanceMetrics struct {
	TotalRequests        int64
	TotalRetries         int64
	TotalRateLimitErrors int64
	TotalResponseTime    time.Duration
	TokenWaitTime        time.Duration
	lock                 sync.Mutex
}

// ClientAuthConfig represents the structure to read authentication details from a JSON configuration file.
type ClientAuthConfig struct {
	TenantID           string `json:"tenantID,omitempty"`
	TenantName         string `json:"tenantName,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	ClientID           string `json:"clientID,omitempty"`
	ClientSecret       string `json:"clientSecret,omitempty"`
	CertificatePath    string `json:"certificatePath,omitempty"`    // Path to the certificate file
	CertificateKeyPath string `json:"certificateKeyPath,omitempty"` // Path to the certificate key file
	CertThumbprint     string `json:"certThumbprint,omitempty"`     // Certificate thumbprint
}

// StructuredError represents a structured error response from the API.
/*
type StructuredError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
*/
// StructuredError represents a detailed API error response.
type StructuredError struct {
	Error struct {
		Code       string `json:"code"`
		Message    string `json:"message"`
		InnerError struct {
			Date            string `json:"date"`
			RequestID       string `json:"request-id"`
			ClientRequestID string `json:"client-request-id"`
		} `json:"innerError"`
	} `json:"error"`
}

// ClientOption defines a function type for modifying client properties during initialization.
type ClientOption func(*Client)

// NewClient initializes a new http client instance with the given baseURL, logger, concurrency manager and client configuration
/*
If TokenLifespan and BufferPeriod aren't set in the config, they default to 30 minutes and 5 minutes, respectively.
If TotalRetryDuration isn't set in the config, it defaults to 1 minute.
If no logger is provided, a default logger will be used.
Any additional options provided will be applied to the client during initialization.
Detect authentication method based on supplied credential type
*/
func NewClient(config Config, clientAuthConfig *ClientAuthConfig, logger Logger, options ...ClientOption) (*Client, error) {

	// Validate MaxRetryAttempts
	if config.MaxRetryAttempts < 0 {
		return nil, fmt.Errorf("MaxRetryAttempts cannot be negative")
	}

	// Validate LogLevel
	if config.LogLevel < LogLevelNone || config.LogLevel > LogLevelDebug {
		return nil, fmt.Errorf("invalid LogLevel")
	}

	// Validate MaxConcurrentRequests
	if config.MaxConcurrentRequests < 0 {
		return nil, fmt.Errorf("MaxConcurrentRequests cannot be negative")
	}

	// Validate TokenLifespan
	if config.TokenLifespan < 0 {
		return nil, fmt.Errorf("TokenLifespan cannot be negative")
	}

	// Validate TokenRefreshBufferPeriod
	if config.TokenRefreshBufferPeriod < 0 {
		return nil, fmt.Errorf("TokenRefreshBufferPeriod cannot be negative")
	}

	// Validate TotalRetryDuration
	if config.TotalRetryDuration < 0 {
		return nil, fmt.Errorf("TotalRetryDuration cannot be negative")
	}

	// Default settings if not supplied
	if config.TokenLifespan == 0 {
		config.TokenLifespan = 30 * time.Minute
	}
	if config.TokenRefreshBufferPeriod == 0 {
		config.TokenRefreshBufferPeriod = 5 * time.Minute
	}
	if config.TotalRetryDuration == 0 {
		config.TotalRetryDuration = 60 * time.Second
	}

	if logger == nil {
		logger = NewDefaultLogger()
	}

	// Set the log level of the logger
	logger.SetLevel(config.LogLevel)

	client := &Client{
		TenantName:     clientAuthConfig.TenantName,
		TenantID:       clientAuthConfig.TenantID,
		httpClient:     &http.Client{Timeout: DefaultTimeout},
		config:         config,
		logger:         logger,
		ConcurrencyMgr: NewConcurrencyManager(config.MaxConcurrentRequests, logger, config.LogLevel >= LogLevelDebug),
		PerfMetrics:    ClientPerformanceMetrics{},
	}

	// Set authentication credentials and determine AuthMethod
	client.SetGraphAuthenticationMethod(map[string]string{
		"clientID":           clientAuthConfig.ClientID,
		"clientSecret":       clientAuthConfig.ClientSecret,
		"certificatePath":    clientAuthConfig.CertificatePath,
		"certificateKeyPath": clientAuthConfig.CertificateKeyPath,
		"certThumbprint":     clientAuthConfig.CertThumbprint,
	})

	// Apply any additional client options provided during initialization
	for _, opt := range options {
		opt(client)
	}

	// Start the periodic metric evaluation for adjusting concurrency.
	go client.StartMetricEvaluation()

	if client.config.LogLevel >= LogLevelDebug {
		client.logger.Debug(
			"New client initialized with the following details:",
			"TenantName", client.TenantName,
			"AuthMethod", client.AuthMethod,
			"Timeout", client.httpClient.Timeout,
			"TokenLifespan", client.config.TokenLifespan,
			"TokenRefreshBufferPeriod", client.config.TokenRefreshBufferPeriod,
			"TotalRetryDuration", client.config.TotalRetryDuration,
			"MaxRetryAttempts", client.config.MaxRetryAttempts,
			"MaxConcurrentRequests", client.config.MaxConcurrentRequests,
			"EnableDynamicRateLimiting", client.config.EnableDynamicRateLimiting,
		)
	}

	return client, nil
}
