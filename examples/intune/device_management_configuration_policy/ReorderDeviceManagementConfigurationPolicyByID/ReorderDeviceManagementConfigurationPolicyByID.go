package main

import (
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

	// Replace with your actual policy ID and desired new priority
	policyID := "57ddf6b9-29d0-43d6-9a1d-5b9688da0487"
	newPriority := 8

	// Call ReorderDeviceManagementConfigurationPolicyByID
	reorderedPolicy, err := intune.ReorderDeviceManagementConfigurationPolicyByID(policyID, newPriority)
	if err != nil {
		log.Fatalf("Error reordering policy: %v", err)
	}

	// Output the result
	fmt.Printf("Reordered Policy: %+v\n", reorderedPolicy)
}
