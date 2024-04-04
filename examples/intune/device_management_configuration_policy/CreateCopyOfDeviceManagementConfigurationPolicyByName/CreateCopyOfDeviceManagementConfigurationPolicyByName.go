package main

import (
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

	// Example policy ID to get
	sourcePolicyName := "[Base] Dev | Windows - Settings Catalog | Microsoft Security Baseline | MSFT Windows 11 22H2 - Computer [Device] ver1.0"
	copyDisplayName := "intuneSDK policy copy"
	copyDescription := "New Policy Description"

	// Create a copy of the policy
	copiedPolicy, err := intune.CreateCopyOfDeviceManagementConfigurationPolicyByName(sourcePolicyName, copyDisplayName, copyDescription)
	if err != nil {
		log.Fatalf("Failed to create a copy of the policy: %v", err)
	}

	// Output the details of the copied policy
	fmt.Printf("Copied Policy: %+v\n", copiedPolicy)
}
