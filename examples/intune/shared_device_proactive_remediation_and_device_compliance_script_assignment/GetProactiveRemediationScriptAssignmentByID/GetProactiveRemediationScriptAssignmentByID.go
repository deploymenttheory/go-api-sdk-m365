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

	// Define scriptID and assignmentID for the proactive remediation script
	scriptID := "ffd8de7a-e0aa-4f14-b917-f644f781c1fc"
	assignmentID := "1c4f3adf-ebe8-422c-97b1-f174632d7538"

	// Call the GetDeviceComplianceScriptAssignmentByID function
	assignment, err := client.GetProactiveRemediationScriptAssignmentByID(scriptID, assignmentID)
	if err != nil {
		log.Fatalf("Failed to get proactive remediation script assignment by ID: %v", err)
	}

	// Pretty print the assignment
	jsonData, err := json.MarshalIndent(assignment, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal assignment: %v", err)
	}
	fmt.Println(string(jsonData))
}
