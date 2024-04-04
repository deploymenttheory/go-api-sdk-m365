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
	// Replace with your actual policy ID and desired new priority
	policyID := "57ddf6b9-29d0-43d6-9a1d-5b9688da0487"
	newPriority := 8

	// Call ReorderDeviceManagementConfigurationPolicyByID
	reorderedPolicy, err := intune.ReorderDeviceManagementConfigurationPolicyByID(policyID, newPriority)
	if err != nil {
		log.Fatalf("Error reordering policy: %v", err)
	}

	// Output the result
	fmt.Printf("Reordered Policy: %+v\n", reorderedPolicy)
}
