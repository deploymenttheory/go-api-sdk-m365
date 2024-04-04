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

	// Example script ID to delete
	scriptID := "e065dfe3-55f7-4260-99b4-aa4beb727297"

	// Call the function to delete the device management script by ID
	err = intune.DeleteDeviceComplianceScriptByID(scriptID)
	if err != nil {
		log.Fatalf("Error deleting device management script: %v", err)
	}

	fmt.Println("Device management script deleted successfully")
}
