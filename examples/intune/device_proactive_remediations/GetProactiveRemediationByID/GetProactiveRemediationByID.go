package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/http_client" // Import http_client for logging
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	clientAuthConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logger.SetLevel(http_client.LogLevelDebug) // Adjust the log level as needed

	// Configuration for the HTTP client
	httpClientconfig := http_client.Config{
		LogLevel:                  http_client.LogLevelDebug,
		MaxRetryAttempts:          3,
		EnableDynamicRateLimiting: true,
		Logger:                    logger,
		MaxConcurrentRequests:     5,
	}

	// initialize HTTP client instance
	httpClient, err := http_client.NewClient(httpClientconfig, clientAuthConfig, logger)
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}

	// Create an Intune client with the HTTP client
	intune := &intuneSDK.Client{HTTP: httpClient}

	// Specify the ID of the Proactive Remediation you want to retrieve
	remediationID := "9a25df0c-2268-48a9-95ac-45de11f82e2c"

	// Call GetProactiveRemediationByID to fetch the details of the specified remediation
	remediation, err := intune.GetProactiveRemediationByID(remediationID)
	if err != nil {
		log.Fatalf("Failed to get proactive remediation by ID: %v", err)
	}

	// Pretty print the details of the proactive remediation
	jsonData, err := json.MarshalIndent(remediation, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal proactive remediation: %v", err)
	}
	fmt.Println(string(jsonData))
}
