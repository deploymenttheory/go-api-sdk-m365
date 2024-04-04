package main

import (
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
	sourcePolicyID := "17436f8b-a93c-45d6-a204-6a80d3d43155"
	copyDisplayName := "New Policy Display Name"
	copyDescription := "New Policy Description"

	// Create a copy of the policy
	copiedPolicy, err := intune.CreateCopyOfDeviceManagementConfigurationPolicyByID(sourcePolicyID, copyDisplayName, copyDescription)
	if err != nil {
		log.Fatalf("Failed to create a copy of the policy: %v", err)
	}

	// Output the details of the copied policy
	fmt.Printf("Copied Policy: %+v\n", copiedPolicy)
}
