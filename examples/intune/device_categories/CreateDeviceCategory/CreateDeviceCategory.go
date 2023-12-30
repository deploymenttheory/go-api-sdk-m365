package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/http_client"
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

	// Prepare the device category data
	newCategory := intuneSDK.ResourceDeviceCategory{
		DisplayName: "Test Category",
		Description: "Description of Test Category",
	}

	// Call CreateDeviceCategory to create a new category
	createdCategory, err := intune.CreateDeviceCategory(&newCategory)
	if err != nil {
		log.Fatalf("Failed to create device category: %v", err)
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(createdCategory, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}