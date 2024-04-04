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

	// Specify the ID of the Proactive Remediation you want to retrieve
	remediationID := "9a25df0c-2268-48a9-95ac-45de11f82e2c"

	// Call GetDeviceProactiveRemediationScriptByID to fetch the details of the specified remediation
	remediation, err := client.GetDeviceProactiveRemediationScriptByID(remediationID)
	if err != nil {
		log.Fatalf("Failed to get proactive remediation by ID: %v", err)
	}

	// Pretty print the details of the proactive remediation
	jsonData, err := json.MarshalIndent(remediation, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal proactive remediation: %v", err)
	}
	fmt.Println(string(jsonData))
}
