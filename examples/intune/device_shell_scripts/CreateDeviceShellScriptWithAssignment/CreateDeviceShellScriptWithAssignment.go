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

	// Define the new script details
	newScriptDetails := intuneSDK.ResourceDeviceShellScript{
		ExecutionFrequency:          "PT15M",
		RetryCount:                  3,
		BlockExecutionNotifications: true,
		DisplayName:                 "intuneSDK shell script test with assignment",
		Description:                 "Description value",
		ScriptContent:               "c2NyaXB0Q29udGVudA==", // Must be base64 encoded.
		RunAsAccount:                "user",
		FileName:                    "NewScript.sh",
		RoleScopeTagIds:             []string{"0"},
	}

	// Define the assignment details
	assignmentDetails := intuneSDK.RequestDeviceManagementScriptAssignment{
		ResourceDeviceManagementScriptGroupAssignments: []intuneSDK.ResourceDeviceManagementScriptGroupAssignment{
			{
				TargetGroupID: "ea8e2fb8-e909-44e6-bae7-56757cf6f347",
			},
		},
	}

	// Create the device management script and assignment
	createdScript, err := intune.CreateDeviceShellScriptWithAssignment(&newScriptDetails, &assignmentDetails)
	if err != nil {
		log.Fatalf("Failed to create device shell script and assignment: %v", err)
	} else {
		log.Printf("Created device shell script ID: %s", createdScript.ID)
	}

	// Pretty print the created device shell script
	jsonData, err := json.MarshalIndent(createdScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal created device shell script: %v", err)
	}
	fmt.Println("Created Device Shell Script:")
	fmt.Println(string(jsonData))

	// Pretty print the created assignment
	jsonData, err = json.MarshalIndent(assignmentDetails, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal created assignment: %v", err)
	}
	fmt.Println("Created Assignment:")
	fmt.Println(string(jsonData))
}
