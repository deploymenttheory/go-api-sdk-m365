package main

import (
	"fmt"
	"log"
	"os"

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

	// Define the script ID and group ID you want to assign
	scriptID := "2fa362c5-22cd-42a6-b7d3-2bdde4850356" // Replace with your actual script ID
	groupID := "f15ab091-95be-4ff2-bc28-7dd10f5b6829"  // Replace with your actual group ID

	// Call the assign function
	err = intune.AssignDeviceManagementScriptByGroupID(scriptID, groupID)
	if err != nil {
		fmt.Printf("Error assigning device management script: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Device management script assigned successfully.")
}
