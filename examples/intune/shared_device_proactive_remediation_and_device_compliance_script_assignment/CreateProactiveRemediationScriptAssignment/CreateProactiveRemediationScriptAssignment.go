package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
	// Import http_client for logging
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/msgraph/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration
	client, err := intune.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace with the actual script ID
	scriptID := "ffd8de7a-e0aa-4f14-b917-f644f781c1fc"

	// Prepare the assignment data
	assignment := intune.ResourceDeviceHealthScriptAssignment{
		// Fill in the necessary fields...
		Target: intune.ResourceDeviceHealthScriptAssignmentTarget{
			DeviceAndAppManagementAssignmentFilterID:   "99b2823d-a05c-4316-9a82-3efa40ff482d",
			DeviceAndAppManagementAssignmentFilterType: "include", // include / exclude
			CollectionID: "", // used for Config Mgr only
		},
		RunRemediationScript: false,
		RunSchedule: intune.ResourceDeviceHealthScriptAssignmentSchedule{
			Interval: 1,
			UseUTC:   false,
			Time:     "01:00:00.0000000",
		},
	}

	// Create the proactive remediation script assignment
	response, err := client.CreateProactiveRemediationScriptAssignment(scriptID, assignment)
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
