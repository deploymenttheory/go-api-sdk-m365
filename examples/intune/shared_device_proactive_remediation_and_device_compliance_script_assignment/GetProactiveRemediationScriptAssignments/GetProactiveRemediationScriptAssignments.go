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

	// Replace 'scriptID' with the actual ID of the Device Health Script
	scriptID := "ffd8de7a-e0aa-4f14-b917-f644f781c1fc"

	// Retrieve the Device Health Script Assignments
	assignments, err := intune.GetProactiveRemediationScriptAssignments(scriptID)
	if err != nil {
		log.Fatalf("Failed to get proactive remediation assignments: %v", err)
	}

	// Pretty print the list of device health scripts
	jsonData, err := json.MarshalIndent(assignments, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal proactive remediations: %v", err)
	}
	fmt.Println(string(jsonData))
}
