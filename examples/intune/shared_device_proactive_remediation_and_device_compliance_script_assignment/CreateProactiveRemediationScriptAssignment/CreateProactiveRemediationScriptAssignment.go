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

	// Replace with the actual script ID
	scriptID := "ffd8de7a-e0aa-4f14-b917-f644f781c1fc"

	// Prepare the assignment data
	assignment := intuneSDK.ResourceDeviceHealthScriptAssignment{
		// Fill in the necessary fields...
		Target: intuneSDK.ResourceDeviceHealthScriptAssignmentTarget{
			DeviceAndAppManagementAssignmentFilterID:   "99b2823d-a05c-4316-9a82-3efa40ff482d",
			DeviceAndAppManagementAssignmentFilterType: "include", // include / exclude
			CollectionID: "", // used for Config Mgr only
		},
		RunRemediationScript: false,
		RunSchedule: intuneSDK.ResourceDeviceHealthScriptAssignmentSchedule{
			Interval: 1,
			UseUTC:   false,
			Time:     "01:00:00.0000000",
		},
	}

	// Create the proactive remediation script assignment
	response, err := intune.CreateProactiveRemediationScriptAssignment(scriptID, assignment)
	if err != nil {
		log.Fatalf("Failed to create proactive remediation script assignment: %v", err)
	}

	// Pretty print the response
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	fmt.Println(string(jsonData))
}
