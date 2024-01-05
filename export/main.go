package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	complianceScripts "github.com/deploymenttheory/go-api-sdk-m365/export/intune/compliance_scripts"
	powershellScripts "github.com/deploymenttheory/go-api-sdk-m365/export/intune/powershell_scripts"
	proactiveRemediations "github.com/deploymenttheory/go-api-sdk-m365/export/intune/proactive_remediations"
	shellScripts "github.com/deploymenttheory/go-api-sdk-m365/export/intune/shell_scripts"
	"github.com/deploymenttheory/go-api-sdk-m365/sdk/http_client"
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

// Config represents the configuration items for intune go.
type Config struct {
	TenantName         string `json:"tenantName"`
	TenantID           string `json:"tenantID"`
	ClientID           string `json:"clientID"`
	ClientSecret       string `json:"clientSecret"`
	OutputPath         string `json:"outputPath"`
	OutputFormat       string `json:"outputFormat"`       // "json" or "yaml"
	ExcludeAssignments bool   `json:"excludeAssignments"` // true or false
	Prefix             string `json:"prefix"`             // Prefix for the backups
	AppendID           bool   `json:"appendID"`           // true or false
}

func main() {
	fmt.Println(`
  __    _____  ___  ___________  ____  ____  _____  ___    _______       _______     ______    
 |" \  (\"   \|"  \("     _   ")("  _||_ " |(\"   \|"  \  /"     "|     /" _   "|   /    " \   
 ||  | |.\\   \    |)__/  \\__/ |   (  ) : ||.\\   \    |(: ______)    (: ( \___)  // ____  \  
 |:  | |: \.   \\  |   \\_ /    (:  |  | . )|: \.   \\  | \/    |       \/ \      /  /    ) :) 
 |.  | |.  \    \. |   |.  |     \\ \__/ // |.  \    \. | // ___)_      //  \ ___(: (____/ //  
 /\  |\|    \    \ |   \:  |     /\\ __ //\ |    \    \ |(:      "|    (:   _(  _|\        /   
(__\_|_)\___|\____\)    \__|    (__________) \___|\____\) \_______)     \_______)  \"_____/    
    `)

	// Define command-line flags
	var configFilePath, outputPath string
	flag.StringVar(&configFilePath, "config", "config.json", "Path to the configuration file")
	flag.StringVar(&outputPath, "output", "", "Output path for backups")
	flag.Parse()

	// Load or create config
	config := loadOrCreateConfig(configFilePath)
	if outputPath != "" {
		config.OutputPath = outputPath
	}
	saveConfig(configFilePath, config)

	// Initialize the HTTP client with config values
	httpClient, err := initializeHTTPClient(config)
	if err != nil {
		log.Fatalf("Failed to initialize HTTP client: %v", err)
	}

	// Create an Intune client with the HTTP client
	intuneClient := &intuneSDK.Client{HTTP: httpClient}

	// Call backup functions based on user input or configuration
	performBackups(intuneClient, config)
}

// loadOrCreateConfig loads the configuration from the specified path.
// If the file does not exist, it prompts the user to enter configuration details,
// creates a new Config struct with these details, and returns it.
func loadOrCreateConfig(path string) *Config {
	config := &Config{}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		config = &Config{
			TenantName:         promptForInput("Enter Tenant Name: "),
			TenantID:           promptForInput("Enter Tenant ID: "),
			ClientID:           promptForInput("Enter Client ID: "),
			ClientSecret:       promptForInput("Enter Client Secret: "),
			OutputPath:         promptForInput("Enter Output Path: "),
			OutputFormat:       promptForInput("Enter Output Format (json/yaml): "),
			ExcludeAssignments: promptForBool("Exclude Assignments? (true/false): "),
			Prefix:             promptForInput("Enter Backup Prefix: "),
			AppendID:           promptForBool("Append ID to backups? (true/false): "),
		}
	} else {
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("Failed to open config file: %v", err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(config); err != nil {
			log.Fatalf("Failed to parse config file: %v", err)
		}
	}
	return config
}

// saveConfig saves the provided Config struct to a file at the specified path.
func saveConfig(path string, config *Config) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Failed to create config file: %v", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(config); err != nil {
		log.Fatalf("Failed to write to config file: %v", err)
	}
}

// promptForInput prompts the user with the provided string and returns the input.
// It repeats the prompt until a non-empty input is received.
func promptForInput(prompt string) string {
	var input string
	for input == "" {
		fmt.Print(prompt)
		fmt.Scanln(&input)
	}
	return input
}

// promptForBool prompts the user with the provided string and expects 'true' or 'false' as input.
// It returns a boolean value corresponding to the input. The prompt is repeated until
// a valid 'true' or 'false' input is received.
func promptForBool(prompt string) bool {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)
		if input == "true" {
			return true
		} else if input == "false" {
			return false
		}
		fmt.Println("Please enter 'true' or 'false'.")
	}
}

func initializeHTTPClient(config *Config) (*http_client.Client, error) {

	// Prepare the ClientAuthConfig using the configuration
	clientAuthConfig := &http_client.ClientAuthConfig{
		TenantID:     config.TenantID,
		TenantName:   config.TenantName,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		// Include other necessary fields such as CertificatePath, CertificateKeyPath, CertThumbprint if needed
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logger.SetLevel(http_client.LogLevelDebug) // Adjust the log level as needed

	// Configuration for the HTTP client
	httpClientConfig := http_client.Config{
		LogLevel:                  http_client.LogLevelDebug,
		MaxRetryAttempts:          3,
		EnableDynamicRateLimiting: true,
		Logger:                    logger,
		MaxConcurrentRequests:     5,
	}

	// Initialize HTTP client instance
	return http_client.NewClient(httpClientConfig, clientAuthConfig, logger)
}

func performBackups(intuneClient *intuneSDK.Client, config *Config) {

	if err := complianceScripts.Backup(intuneClient, config.OutputPath, config.OutputFormat, config.ExcludeAssignments, config.Prefix, config.AppendID); err != nil {
		fmt.Println("Error during compliance script backup:", err)
	}
	if err := shellScripts.Backup(intuneClient, config.OutputPath, config.OutputFormat, config.ExcludeAssignments, config.Prefix, config.AppendID); err != nil {
		fmt.Println("Error during macOS shell script backup:", err)
	}
	if err := powershellScripts.Backup(intuneClient, config.OutputPath, config.OutputFormat, config.ExcludeAssignments, config.Prefix, config.AppendID); err != nil {
		fmt.Println("Error during Powershell Script backup:", err)
	}
	if err := proactiveRemediations.Backup(intuneClient, config.OutputPath, config.OutputFormat, config.ExcludeAssignments, config.Prefix, config.AppendID); err != nil {
		fmt.Println("Error during proactive remediation script backup:", err)
	}
	// Similarly, call other backup functions as required.
}
