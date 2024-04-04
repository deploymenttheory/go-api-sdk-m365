package main

import (
	"encoding/json"
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

	// Call GetDeviceProactiveRemediationScripts to fetch the list of device health scripts
	remediations, err := intune.GetDeviceProactiveRemediationScripts()
	if err != nil {
		log.Fatalf("Failed to get proactive remediations: %v", err)
	}

	// Pretty print the list of device health scripts
	jsonData, err := json.MarshalIndent(remediations, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal proactive remediations: %v", err)
	}
	fmt.Println(string(jsonData))
}
