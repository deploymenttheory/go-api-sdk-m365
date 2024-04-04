package main

import (
	"encoding/json"
	"fmt"
	"log"

	// Import http_client for logging

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
	updatedScriptDetails := intune.ResourceDeviceManagementScript{
		DisplayName:           "intune - Updated Script by display name",
		Description:           "This is a new script created for demonstration purposes.",
		ScriptContent:         "c2NyaXB0Q29udGVudA==", // Must be base64 encoded.
		RunAsAccount:          "system",               // or "user"
		EnforceSignatureCheck: false,
		FileName:              "NewScript.ps1",
		RoleScopeTagIds:       []string{"0"},
		RunAs32Bit:            false,
	}

	deviceManagementScriptName := "intune - Updated Script"

	// Create the new device management script
	newScript, err := client.UpdateDeviceManagementScriptByDisplayName(deviceManagementScriptName, &updatedScriptDetails)
	if err != nil {
		log.Fatalf("Failed to create device management script: %v", err)
	}

	// Pretty print the created device management script
	jsonData, err := json.MarshalIndent(newScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal created device management script: %v", err)
	}
	fmt.Println(string(jsonData))
}
