package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	// Import http_client for logging
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example policy name to get
	policyName := "[Base] Dev | Windows - Settings Catalog | Microsoft Teams ver0.1"

	// Use the Intune client to perform operations
	deviceManagementPolicy, err := intune.GetDeviceManagementConfigurationPolicyByName(policyName)
	if err != nil {
		log.Fatalf("Failed to get device configuration policy: %v", err)
	}

	// Pretty print the device configuration policy
	jsonData, err := json.MarshalIndent(deviceManagementPolicy, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device configuration policy: %v", err)
	}
	fmt.Println(string(jsonData))
}
