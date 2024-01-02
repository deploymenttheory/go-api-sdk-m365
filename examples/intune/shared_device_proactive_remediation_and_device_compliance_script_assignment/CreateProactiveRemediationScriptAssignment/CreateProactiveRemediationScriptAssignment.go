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

	// Replace 'scriptID' with the actual ID of the Device Health Script
	scriptID := "your_device_health_script_id"

	// Define the assignment request
	assignmentRequest := &intuneSDK.DeviceHealthScriptAssignmentItem{
		Target: intuneSDK.ConfigurationManagerCollectionTarget{
			DeviceAndAppManagementAssignmentFilterID:   "your_filter_id",
			DeviceAndAppManagementAssignmentFilterType: "include",
			CollectionID: "your_collection_id",
		},
		RunRemediationScript: true,
		RunSchedule: intuneSDK.DeviceHealthScriptAssignmentSchedule{
			Interval: 8,
			UseUTC:   true,
			Time:     "11:58:36.2550000",
		},
	}

	// Create Device Health Script Assignment
	response, err := intune.CreateProactiveRemediationScriptAssignment(scriptID, assignmentRequest)
	if err != nil {
		log.Fatalf("Failed to create device health script assignment: %v", err)
	}

	// Output the response
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	fmt.Println(string(jsonData))
}
