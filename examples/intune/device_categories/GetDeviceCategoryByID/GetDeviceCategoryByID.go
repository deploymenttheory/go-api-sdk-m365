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

	deviceCategoryID := "018cfd5d-992f-4780-a557-468e98888537"

	// Use the Intune client to perform operations
	deviceCategory, err := client.GetDeviceCategoryByID(deviceCategoryID)
	if err != nil {
		log.Fatalf("Failed to get device category: %v", err)
	}

	// Pretty print the device category
	jsonData, err := json.MarshalIndent(deviceCategory, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device category: %v", err)
	}
	fmt.Println(string(jsonData))
}
