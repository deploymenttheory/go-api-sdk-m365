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
	scriptID := "61966ecb-29f1-469f-9cbc-8b3e664f8d96"

	// Call the function to delete the device management script by ID
	err = intune.DeleteDeviceManagementScriptByID(scriptID)
	if err != nil {
		log.Fatalf("Error deleting device management script: %v", err)
	}

	fmt.Println("Device management script deleted successfully")
}
