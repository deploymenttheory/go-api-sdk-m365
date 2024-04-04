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
	// Example policy name to get
	policyName := "intune | [Base] Dev | Windows - Settings Catalog | Delivery Optimization ver0.1"

	// Delete the policy
	err = client.DeleteDeviceManagementConfigurationPolicyByName(policyName)
	if err != nil {
		fmt.Printf("Error deleting policy: %s\n", err)
		return
	}

	fmt.Println("Policy deleted successfully")
}
