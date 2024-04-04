// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicy",
// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
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

	// Construct the request body
	requestBody := &intune.ResourceDeviceComplianceScript{
		ODataType:              "#microsoft.graph.deviceComplianceScript",
		Publisher:              "Publisher value",
		Version:                "Version value",
		DisplayName:            "intune - Device Compliance Script",
		Description:            "Description value",
		DetectionScriptContent: "ZGV0ZWN0aW9uU2NyaXB0Q29udGVudA==", // Base64 encoded script content
		RunAsAccount:           "user",
		EnforceSignatureCheck:  true,
		RunAs32Bit:             true,
		RoleScopeTagIds:        []string{"0"},
	}

	// Create the new policy
	createdPolicy, err := client.CreateDeviceComplianceScript(requestBody)
	if err != nil {
		fmt.Printf("Error creating policy: %s\n", err)
		return
	}

	// Pretty print the device compliance script
	jsonData, err := json.MarshalIndent(createdPolicy, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device compliance script: %v", err)
	}
	fmt.Println(string(jsonData))
}
