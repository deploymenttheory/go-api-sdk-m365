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

	// Specify the ID of the Proactive Remediation you want to retrieve
	remediationDisplayName := "intuneSDK - proactive remediation created from JSON"

	// Call GetDeviceProactiveRemediationScriptByDisplayName to fetch the details of the specified remediation
	remediation, err := intune.GetDeviceProactiveRemediationScriptByDisplayName(remediationDisplayName)
	if err != nil {
		log.Fatalf("Failed to get proactive remediation by Name: %v", err)
	}

	// Pretty print the details of the proactive remediation
	jsonData, err := json.MarshalIndent(remediation, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal proactive remediation: %v", err)
	}
	fmt.Println(string(jsonData))
}
