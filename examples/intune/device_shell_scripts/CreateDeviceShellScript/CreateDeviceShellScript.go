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

	// Define the new script details
	newScriptDetails := intuneSDK.ResourceDeviceShellScript{
		ExecutionFrequency:          "PT15M",
		RetryCount:                  3,
		BlockExecutionNotifications: true,
		DisplayName:                 "intune SDK macOS shell script creation test",
		Description:                 "Description value",
		ScriptContent:               "c2NyaXB0Q29udGVudA==", // Must be base64 encoded.
		RunAsAccount:                "user",
		FileName:                    "NewScript.sh",
		RoleScopeTagIds:             []string{"0"},
	}

	// Create the new device shell script
	newScript, err := intune.CreateDeviceShellScript(&newScriptDetails)
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
