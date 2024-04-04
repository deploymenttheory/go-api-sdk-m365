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

	// Define scriptID and assignmentID for the device compliance script
	scriptID := "your-device-compliance-script-id"
	assignmentID := "your-assignment-id"

	// Call the GetDeviceComplianceScriptAssignmentByID function
	assignment, err := intune.GetDeviceComplianceScriptAssignmentByID(scriptID, assignmentID)
	if err != nil {
		log.Fatalf("Failed to get device compliance script assignment by ID: %v", err)
	}

	// Pretty print the assignment
	jsonData, err := json.MarshalIndent(assignment, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal assignment: %v", err)
	}
	fmt.Println(string(jsonData))
}
