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
	updatedDeviceCategory := intuneSDK.ResourceDeviceCategory{
		DisplayName: "Updated category name 3",
		Description: "Description of Test Category",
	}

	categoryName := "Updated category name 2"

	// Call UpdateDeviceCategoryByDisplayName to update a new category
	updatedCategory, err := intune.UpdateDeviceCategoryByDisplayName(categoryName, &updatedDeviceCategory)
	if err != nil {
		log.Fatalf("Failed to update device category: %v", err)
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(updatedCategory, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}
