// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicy",
// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
package main

import (
	"fmt"
	"log"

	// Import http_client for logging

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
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

	scriptID := "8c3d2ec3-3e63-4df3-8265-69bbba1e53e5"

	// Create a Device Compliance Script Assignment
	assignmentRequest := &intuneSDK.AssignDeviceComplianceScript{
		DeviceComplianceScriptAssignments: []intuneSDK.DeviceComplianceAssignmentItem{
			{
				ID: "1148a469-8a3a-43a9-9e2f-1687cbe56538", // Replace with actual assignment ID

				Target: intuneSDK.DeviceComplianceAssignmentItemTarget{
					DeviceAndAppManagementAssignmentFilterID:   "",
					DeviceAndAppManagementAssignmentFilterType: "",
					CollectionID: "",
				},
				RunRemediationScript: true,
				RunSchedule: intuneSDK.DeviceComplianceAssignmentItemRunSchedule{
					Interval: 8,
					UseUTC:   true,
					Time:     "11:58:36.2550000",
				},
			},
		},
	}

	// Call CreateDeviceComplianceScriptAssignment
	err = client.CreateDeviceComplianceScriptAssignment(scriptID, assignmentRequest)
	if err != nil {
		log.Fatalf("Failed to create device compliance script assignment: %v", err)
	}

	fmt.Println("Device Compliance Script Assignment created successfully.")
}
