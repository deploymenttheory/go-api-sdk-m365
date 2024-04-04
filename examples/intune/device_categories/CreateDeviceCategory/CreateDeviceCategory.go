package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Prepare the device category data
	newCategory := intuneSDK.ResourceDeviceCategory{
		DisplayName: "Test Category",
		Description: "Description of Test Category",
	}

	// Call CreateDeviceCategory to create a new category
	createdCategory, err := client.CreateDeviceCategory(&newCategory)
	if err != nil {
		log.Fatalf("Failed to create device category: %v", err)
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(createdCategory, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}
