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

	// Define the new script details
	updatedScriptDetails := intuneSDK.ResourceDeviceManagementScriptRequest{
		OdataType:             "#microsoft.graph.deviceManagementScript", // Update this according to the Intune API requirements
		DisplayName:           "New name of Script",
		Description:           "This is a new script created for demonstration purposes.",
		ScriptContent:         "c2NyaXB0Q29udGVudA==", // Must be base64 encoded.
		RunAsAccount:          "system",               // or "user"
		EnforceSignatureCheck: false,
		FileName:              "NewScript.ps1",
		RoleScopeTagIds:       []string{"0"},
		RunAs32Bit:            false,
	}

	deviceManagementScriptName := "New Script"

	// Create the new device management script
	newScript, err := intune.UpdateDeviceManagementScriptByDisplayName(deviceManagementScriptName, &updatedScriptDetails)
	if err != nil {
		log.Fatalf("Failed to create device management script: %v", err)
	}

	// Pretty print the created device management script
	jsonData, err := json.MarshalIndent(newScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal created device management script: %v", err)
	}
	fmt.Println(string(jsonData))
}
