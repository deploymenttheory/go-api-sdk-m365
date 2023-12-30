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
		DisplayName: "Updated category name 2",
		Description: "Description of Test Category",
	}

	categoryId := "acb435dc-7e6d-4eaa-a987-7ada867be594"

	// Call UpdateDeviceCategoryByID to update a new category
	updatedCategory, err := intune.UpdateDeviceCategoryByID(categoryId, &newCategory)
	if err != nil {
		log.Fatalf("Failed to update device category: %v", err)
	}

	// Pretty print the device category
	jsonData, err := json.MarshalIndent(updatedCategory, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device category: %v", err)
	}
	fmt.Println(string(jsonData))
}
