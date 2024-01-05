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

	// Define scriptID and assignmentID for the proactive remediation script
	scriptID := "ffd8de7a-e0aa-4f14-b917-f644f781c1fc"
	assignmentID := "1c4f3adf-ebe8-422c-97b1-f174632d7538"

	// Call the GetDeviceComplianceScriptAssignmentByID function
	assignment, err := intune.GetProactiveRemediationScriptAssignmentByID(scriptID, assignmentID)
	if err != nil {
		log.Fatalf("Failed to get proactive remediation script assignment by ID: %v", err)
	}

	// Pretty print the assignment
	jsonData, err := json.MarshalIndent(assignment, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal assignment: %v", err)
	}
	fmt.Println(string(jsonData))
}
