// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicy",
// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
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

	// Construct the request body
	requestBody := &intuneSDK.ResourceDeviceComplianceScript{
		ODataType:              "#microsoft.graph.deviceComplianceScript",
		Publisher:              "Publisher value",
		Version:                "Version value",
		DisplayName:            "intuneSDK - Device Compliance Script",
		Description:            "Description value",
		DetectionScriptContent: "ZGV0ZWN0aW9uU2NyaXB0Q29udGVudA==", // Base64 encoded script content
		RunAsAccount:           "user",
		EnforceSignatureCheck:  true,
		RunAs32Bit:             true,
		RoleScopeTagIds:        []string{"0"},
	}

	// Create the new policy
	createdPolicy, err := intune.CreateDeviceComplianceScript(requestBody)
	if err != nil {
		fmt.Printf("Error creating policy: %s\n", err)
		return
	}

	// Pretty print the device compliance script
	jsonData, err := json.MarshalIndent(createdPolicy, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device compliance script: %v", err)
	}
	fmt.Println(string(jsonData))
}
