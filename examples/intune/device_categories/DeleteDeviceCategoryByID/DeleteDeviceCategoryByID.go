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

	// Example category ID to delete
	categoryID := "acb435dc-7e6d-4eaa-a987-7ada867be594"

	// Call the function to delete the device category by ID
	err = client.DeleteDeviceCategoryByID(categoryID)
	if err != nil {
		log.Fatalf("Error deleting device category: %v", err)
	}

	fmt.Println("Device category deleted successfully")
}
