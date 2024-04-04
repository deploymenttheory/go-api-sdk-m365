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

	// Replace 'scriptID' with the actual ID of the Device compliance Script
	scriptID := "ebba8690-c32d-4073-b44b-8a00f4487ae7"

	// Retrieve the Device compliance Script Assignments
	assignments, err := intune.GetDeviceComplianceScriptAssignments(scriptID)
	if err != nil {
		log.Fatalf("Failed to get device compliance script assignments: %v", err)
	}

	// Pretty print the list of device health scripts
	jsonData, err := json.MarshalIndent(assignments, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device compliance script  remediations: %v", err)
	}
	fmt.Println(string(jsonData))
}
