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

	scriptID := "3300b08a-d78e-45db-9a56-ec852464cd5b"

	// Define the assignment details
	assignmentDetails := intuneSDK.RequestDeviceManagementScriptAssignment{
		ResourceDeviceManagementScriptGroupAssignments: []intuneSDK.ResourceDeviceManagementScriptGroupAssignment{
			{
				TargetGroupID: "ea8e2fb8-e909-44e6-bae7-56757cf6f347",
			},
		},
	}
	// Create the new device shell script assignment
	newScript, err := intune.CreateDeviceShellScriptAssignment(scriptID, &assignmentDetails)
	if err != nil {
		log.Fatalf("Failed to create device shell script: %v", err)
	}

	// Pretty print the created device shell script
	jsonData, err := json.MarshalIndent(newScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal created device shell script: %v", err)
	}
	fmt.Println(string(jsonData))
}
