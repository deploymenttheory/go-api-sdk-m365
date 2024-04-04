package main

import (
	"encoding/json"
	"fmt"
	"log"

	// Import http_client for logging

	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/msgraph/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration
	client, err := intune.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	scriptID := "3300b08a-d78e-45db-9a56-ec852464cd5b"

	// Define the assignment details
	assignmentDetails := intuneSDK.RequestDeviceManagementScriptAssignment{
		ResourceDeviceManagementScriptGroupAssignments: []intuneSDK.ResourceDeviceManagementScriptGroupAssignment{
			{
				TargetGroupID: "ea8e2fb8-e909-44e6-bae7-56757cf6f347",
			},
		},
	}
	// Create the new device shell script assignment
	newScript, err := intune.CreateDeviceShellScriptAssignment(scriptID, &assignmentDetails)
	if err != nil {
		log.Fatalf("Failed to create device shell script: %v", err)
	}

	// Pretty print the created device shell script
	jsonData, err := json.MarshalIndent(newScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal created device shell script: %v", err)
	}
	fmt.Println(string(jsonData))
}
