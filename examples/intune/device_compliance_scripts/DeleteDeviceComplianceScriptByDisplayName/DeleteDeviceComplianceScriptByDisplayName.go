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
	// Example script display name to delete
	scriptName := "intuneSDK - Device Compliance Script"

	// Call the function to delete the device management script by tName
	err = client.DeleteDeviceComplianceScriptByDisplayName(scriptName)
	if err != nil {
		log.Fatalf("Error deleting device management script: %v", err)
	}

	fmt.Println("Device management script deleted successfully")
}
