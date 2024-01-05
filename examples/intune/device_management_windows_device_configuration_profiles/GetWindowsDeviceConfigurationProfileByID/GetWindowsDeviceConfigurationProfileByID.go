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

	// Example profile ID to get
	deviceConfigurationProfileID := "12035cf9-156f-46f0-9b80-47749d5e9c16" // 6f511f91-33ba-471a-a2da-6c467c0874cd // 18900079-f55a-4c0d-bc36-bfa292231714

	// Use the Intune client to perform operations
	deviceConfigurationProfile, err := intune.GetWindowsDeviceConfigurationProfileByID(deviceConfigurationProfileID)
	if err != nil {
		log.Fatalf("Failed to get device configuration profile: %v", err)
	}

	// Pretty print the device configuration profile
	jsonData, err := json.MarshalIndent(deviceConfigurationProfile, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device configuration profile: %v", err)
	}
	fmt.Println(string(jsonData))
}
