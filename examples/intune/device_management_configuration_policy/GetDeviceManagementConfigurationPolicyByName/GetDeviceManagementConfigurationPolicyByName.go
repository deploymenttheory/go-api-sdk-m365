package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/msgraph/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration
	client, err := intune.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example policy name to get
	policyName := "[Base] Dev | Windows - Settings Catalog | Microsoft Teams ver0.1"

	// Use the Intune client to perform operations
	deviceManagementPolicy, err := client.GetDeviceManagementConfigurationPolicyByName(policyName)
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
