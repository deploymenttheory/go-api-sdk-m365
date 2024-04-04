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
	scriptID := "3b28afa8-01d6-41dd-a116-243caf29c57d"

	// Call the function to delete the device shell script by ID
	err = intune.DeleteDeviceShellScriptByID(scriptID)
	if err != nil {
		log.Fatalf("Error deleting device shell script: %v", err)
	}

	fmt.Println("Device shell script deleted successfully")
}
