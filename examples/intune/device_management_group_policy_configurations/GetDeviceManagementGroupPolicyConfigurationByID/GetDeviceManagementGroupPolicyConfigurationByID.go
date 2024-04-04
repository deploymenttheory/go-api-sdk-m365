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

	// Example policy ID to get
	groupPolicyConfigurationID := "6f9ba788-f719-46a7-b7c5-d566963d5999" // "7f774f0f-2f2d-4dc3-a76f-6d45af51019e" / "7f774f0f-2f2d-4dc3-a76f-6d45af51019e" / "6f9ba788-f719-46a7-b7c5-d566963d5999"

	// Use the Intune client to perform operations
	deviceManagementGroupPolicyConfiguration, err := client.GetDeviceManagementGroupPolicyConfigurationByID(groupPolicyConfigurationID)
	if err != nil {
		log.Fatalf("Failed to get device configuration policy: %v", err)
	}

	// Pretty print the device configuration policy
	jsonData, err := json.MarshalIndent(deviceManagementGroupPolicyConfiguration, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device configuration policy: %v", err)
	}
	fmt.Println(string(jsonData))
}
