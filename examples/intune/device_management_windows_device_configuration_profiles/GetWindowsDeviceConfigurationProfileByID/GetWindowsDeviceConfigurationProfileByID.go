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

	// Example profile ID to get
	deviceConfigurationProfileID := "12035cf9-156f-46f0-9b80-47749d5e9c16" // 6f511f91-33ba-471a-a2da-6c467c0874cd // 18900079-f55a-4c0d-bc36-bfa292231714

	// Use the Intune client to perform operations
	deviceConfigurationProfile, err := client.GetWindowsDeviceConfigurationProfileByID(deviceConfigurationProfileID)
	if err != nil {
		log.Fatalf("Failed to get device configuration profile: %v", err)
	}

	// Pretty print the device configuration profile
	jsonData, err := json.MarshalIndent(deviceConfigurationProfile, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device configuration profile: %v", err)
	}
	fmt.Println(string(jsonData))
}
