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

	// Create an update request for the Device Shell Script
	updateRequest := &intuneSDK.ResourceDeviceShellScript{
		ExecutionFrequency:          "PT15M",
		RetryCount:                  2,
		BlockExecutionNotifications: true,
		DisplayName:                 "Updated Display Name",
		Description:                 "Updated Description",
		ScriptContent:               "c2NyaXB0Q29udGVudA==",
		RunAsAccount:                "user",
		FileName:                    "Updated File Name",
		RoleScopeTagIds:             []string{"0"},
		Assignments:                 []intuneSDK.ResponseDeviceShellScriptAssignment{},
	}

	// Replace "scriptID" with the actual ID of the Device Shell Script you want to update
	scriptID := "3b28afa8-01d6-41dd-a116-243caf29c57d"

	// Update the Device Shell Script by its ID
	updatedShellScript, err := intune.UpdateDeviceShellScriptByID(scriptID, updateRequest)
	if err != nil {
		log.Fatalf("Failed to update Device Shell Script by ID: %v", err)
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(updatedShellScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}
