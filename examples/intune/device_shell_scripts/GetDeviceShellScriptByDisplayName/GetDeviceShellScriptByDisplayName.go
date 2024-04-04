package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
	"github.com/deploymenttheory/go-api-sdk-m365/sdk/utils"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/msgraph/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration
	client, err := intune.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	deviceShellScriptName := "macOS-shell_script-update_SSH_public_key"

	// Use the Intune client to perform operations
	deviceShellScript, err := client.GetDeviceShellScriptByDisplayName(deviceShellScriptName)
	if err != nil {
		log.Fatalf("Failed to get device management scripts: %v", err)
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(deviceShellScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))

	// Base64 decode the scriptContent field
	decodedContent, err := utils.Base64Decode(deviceShellScript.ScriptContent)
	if err != nil {
		log.Fatalf("Failed to Base64 decode the script content: %v", err)
	}

	// Assuming the decoded content is a string, print it
	fmt.Println("Decoded Intune Script Content:")
	fmt.Println(string(decodedContent))
}
