package main

import (
	"encoding/json"
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

	// Prepare the device category data
	newCategory := intune.ResourceDeviceCategory{
		DisplayName: "Updated category name 2",
		Description: "Description of Test Category",
	}

	categoryId := "acb435dc-7e6d-4eaa-a987-7ada867be594"

	// Call UpdateDeviceCategoryByID to update a new category
	updatedCategory, err := client.UpdateDeviceCategoryByID(categoryId, &newCategory)
	if err != nil {
		log.Fatalf("Failed to update device category: %v", err)
	}

	// Pretty print the device category
	jsonData, err := json.MarshalIndent(updatedCategory, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device category: %v", err)
	}
	fmt.Println(string(jsonData))
}
