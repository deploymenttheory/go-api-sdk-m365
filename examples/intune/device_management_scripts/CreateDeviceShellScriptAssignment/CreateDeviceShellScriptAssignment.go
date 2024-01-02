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

	// Create a request for the script assignment
	assignmentRequest := &intune.RequestDeviceManagementScriptAssignment{
		ResourceDeviceManagementScriptGroupAssignments: []intune.ResourceDeviceManagementScriptGroupAssignment{
			{
				OdataType:     "#microsoft.graph.deviceManagementScriptGroupAssignment", // Use the correct type
				ID:            "assignment1",
				TargetGroupID: "target_group_id", // Replace with the actual target group ID
			},
			// Add more assignments as needed
		},
		ResourceDeviceManagementScriptAssignments: []intune.ResourceDeviceManagementScriptAssignment{
			{
				OdataType: "#microsoft.graph.deviceManagementScriptAssignment", // Use the correct type
				ID:        "assignment2",
				Target: intune.ResponseDeviceShellScriptListTarget{
					ODataType:                                  "#microsoft.graph.configurationManagerCollectionAssignmentTarget", // Use the correct type
					DeviceAndAppManagementAssignmentFilterID:   "filter_id",
					DeviceAndAppManagementAssignmentFilterType: "include",
					CollectionID:                               "collection_id",
				},
			},
			// Add more assignments as needed
		},
	}

	// Replace with the actual script ID you want to assign
	scriptID := "your_script_id"

	// Create the assignment
	createdAssignment, err := intune.CreateDeviceShellScriptAssignment(scriptID, assignmentRequest)
	if err != nil {
		log.Fatalf("Error creating device management script assignment: %v", err)
	}

	// Handle the assignment result
	if createdAssignment != nil {
		fmt.Printf("Device management script assignment created successfully.\n")
		fmt.Printf("Assignment ID: %s\n", createdAssignment.ID)
		fmt.Printf("Script ID: %s\n", scriptID)
	} else {
		fmt.Printf("Device management script assignment creation failed.\n")
	}
}
