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
	scriptName := "Display Name value"

	// Call the function to delete the device shell script by tName
	err = intune.DeleteDeviceShellScriptByDisplayName(scriptName)
	if err != nil {
		log.Fatalf("Error deleting device shell script: %v", err)
	}

	fmt.Println("Device shell script deleted successfully")
}
