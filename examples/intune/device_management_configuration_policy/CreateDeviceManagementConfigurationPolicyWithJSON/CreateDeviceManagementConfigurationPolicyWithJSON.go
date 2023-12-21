// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicy",
// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
package main

import (
	"encoding/json"
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

	// Read the JSON file
	byteValue, err := os.ReadFile("/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/examples/intune/device_management_configuration_policy/CreateDeviceManagementConfigurationPolicyWithJSON/payload.json") // Replace with your JSON file path
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal the JSON data into the struct
	var policyRequest intuneSDK.ResourceDeviceManagementConfigurationPolicy
	err = json.Unmarshal(byteValue, &policyRequest)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Create the new policy
	createdPolicy, err := intune.CreateDeviceManagementConfigurationPolicy(&policyRequest)
	if err != nil {
		fmt.Printf("Error creating policy: %s\n", err)
		return
	}

	fmt.Printf("Created Policy: %+v\n", createdPolicy)
}
