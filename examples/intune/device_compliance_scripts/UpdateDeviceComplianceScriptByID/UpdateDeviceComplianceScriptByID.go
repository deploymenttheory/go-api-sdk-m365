package main

import (
	"encoding/json"
	"fmt"
	"log"

	// Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create an update request for the Device Shell Script
	updateRequestBody := &intuneSDK.ResourceDeviceComplianceScript{
		Publisher:              "Publisher value",
		Version:                "Version value",
		DisplayName:            "intuneSDK - Updated Device Compliance Script",
		Description:            "Description value",
		DetectionScriptContent: "ZGV0ZWN0aW9uU2NyaXB0Q29udGVudA==", // Base64 encoded script content
		RunAsAccount:           "user",
		EnforceSignatureCheck:  true,
		RunAs32Bit:             true,
		RoleScopeTagIds:        []string{"0"},
	}

	// Replace "scriptID" with the actual ID of the Device Shell Script you want to update
	scriptID := "da992c34-ce76-4275-b336-56af95c14988"

	// Update the Device Shell Script by its ID
	updatedShellScript, err := intune.UpdateDeviceComplianceScriptByID(scriptID, updateRequestBody)
	if err != nil {
		log.Fatalf("Failed to update Device Shell Script by ID: %v", err)
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(updatedShellScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}
