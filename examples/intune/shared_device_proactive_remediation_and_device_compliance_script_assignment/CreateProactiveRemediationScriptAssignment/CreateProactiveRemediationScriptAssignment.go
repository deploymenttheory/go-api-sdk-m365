package main

import (
	"encoding/json"
	"fmt"
	"log"

	// Import http_client for logging
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

	// Replace with the actual script ID
	scriptID := "ffd8de7a-e0aa-4f14-b917-f644f781c1fc"

	// Prepare the assignment data
	assignment := intuneSDK.ResourceDeviceHealthScriptAssignment{
		// Fill in the necessary fields...
		Target: intuneSDK.ResourceDeviceHealthScriptAssignmentTarget{
			DeviceAndAppManagementAssignmentFilterID:   "99b2823d-a05c-4316-9a82-3efa40ff482d",
			DeviceAndAppManagementAssignmentFilterType: "include", // include / exclude
			CollectionID: "", // used for Config Mgr only
		},
		RunRemediationScript: false,
		RunSchedule: intuneSDK.ResourceDeviceHealthScriptAssignmentSchedule{
			Interval: 1,
			UseUTC:   false,
			Time:     "01:00:00.0000000",
		},
	}

	// Create the proactive remediation script assignment
	response, err := intune.CreateProactiveRemediationScriptAssignment(scriptID, assignment)
	if err != nil {
		log.Fatalf("Failed to create proactive remediation script assignment: %v", err)
	}

	// Pretty print the response
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	fmt.Println(string(jsonData))
}
