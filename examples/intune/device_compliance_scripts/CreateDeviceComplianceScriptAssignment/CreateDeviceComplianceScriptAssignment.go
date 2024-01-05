// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicy",
// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/http_client" // Import http_client for logging
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	clientAuthConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logger.SetLevel(http_client.LogLevelDebug) // Adjust the log level as needed

	// Configuration for the HTTP client
	httpClientconfig := http_client.Config{
		LogLevel:                  http_client.LogLevelDebug,
		MaxRetryAttempts:          3,
		EnableDynamicRateLimiting: true,
		Logger:                    logger,
		MaxConcurrentRequests:     5,
	}

	// initialize HTTP client instance
	httpClient, err := http_client.NewClient(httpClientconfig, clientAuthConfig, logger)
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}

	// Create an Intune client with the HTTP client
	intune := &intuneSDK.Client{HTTP: httpClient}

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
	err = intune.CreateDeviceComplianceScriptAssignment(scriptID, assignmentRequest)
	if err != nil {
		log.Fatalf("Failed to create device compliance script assignment: %v", err)
	}

	fmt.Println("Device Compliance Script Assignment created successfully.")
}
