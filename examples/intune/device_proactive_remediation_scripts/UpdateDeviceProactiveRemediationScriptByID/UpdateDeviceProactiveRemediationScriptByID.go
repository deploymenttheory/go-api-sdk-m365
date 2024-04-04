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

	// Example: Updating a Proactive Remediation script with a given ID
	scriptID := "fcb4e658-f2e4-440b-95a8-80e9430717fe" // Replace with the actual ID of the script you want to update

	// Prepare the updated data for the Proactive Remediation script
	updateRequest := &intuneSDK.ResourceProactiveRemediation{
		Publisher:                   "admin_d.watkins@deploymenttheory.com ",
		Version:                     "1",
		DisplayName:                 "Remediate Office Click-to-Run updater tool if not Run in 3 Days",
		Description:                 "The Office Click-to-Run updater tool is often always lagging behind on updates. This proactive remediation will check the registry for when the machine last checked for updates, and if more than 3 days ago, clear the reg key and start the Scheduled task.\n\nThe UpdateDetectionLastRunTime key value is in LDAP/Win32 FILETIME which needs to be converted into date/time which is human readable. Can do this in PowerShell which I found from the link below\n\nhttps://www.epochconverter.com/ldap\n\nRef: https://github.com/markkerry/Proactive-Remediations",
		DetectionScriptContent:      "ZGV0ZWN0aW9uU2NyaXB0Q29udGVudA==",
		RemediationScriptContent:    "cmVtZWRpYXRpb25TY3JpcHRDb250ZW50",
		RunAsAccount:                "system",
		EnforceSignatureCheck:       false,
		RunAs32Bit:                  false,
		RoleScopeTagIds:             []string{"0"},
		IsGlobalScript:              false,
		HighestAvailableVersion:     "null", // null if not set.
		DeviceHealthScriptType:      "deviceHealthScript",
		DetectionScriptParameters:   []intuneSDK.DeviceHealthScriptParameter{}, // Empty slice as per JSON
		RemediationScriptParameters: []intuneSDK.DeviceHealthScriptParameter{}, // Empty slice as per JSON
	}

	// Call the UpdateDeviceProactiveRemediationScriptByID function
	updatedScript, err := intune.UpdateDeviceProactiveRemediationScriptByID(scriptID, updateRequest)
	if err != nil {
		fmt.Printf("Error updating Proactive Remediation: %v\n", err)
		return
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(updatedScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}
