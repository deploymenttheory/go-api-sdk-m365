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

	deviceName := "Device Category | Integration Device"

	// Use the Intune client to perform operations
	deviceCategory, err := client.GetDeviceCategoryByDisplayName(deviceName)
	if err != nil {
		log.Fatalf("Failed to get device category scripts: %v", err)
	}

	// Pretty print the device category scripts
	jsonData, err := json.MarshalIndent(deviceCategory, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device category scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}
