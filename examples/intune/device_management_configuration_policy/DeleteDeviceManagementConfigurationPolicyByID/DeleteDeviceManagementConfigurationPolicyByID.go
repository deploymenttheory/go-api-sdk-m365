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
	policyID := "1d9cb549-d495-47c7-8f69-9c97783f1318"

	// Delete the policy
	err = intune.DeleteDeviceManagementConfigurationPolicyByID(policyID)
	if err != nil {
		fmt.Printf("Error deleting policy: %s\n", err)
		return
	}

	fmt.Println("Policy deleted successfully")
}
