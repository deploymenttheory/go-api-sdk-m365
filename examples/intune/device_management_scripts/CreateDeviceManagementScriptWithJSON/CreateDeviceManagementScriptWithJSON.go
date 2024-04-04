// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicy",
// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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

	// Read the JSON file
	byteValue, err := os.ReadFile("/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/examples/intune/device_management_scripts/CreateDeviceManagementScriptWithJSON/payload.json") // Replace with your JSON file path
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal the JSON data into the struct
	var powershellScriptRequest intune.ResourceDeviceManagementScript
	err = json.Unmarshal(byteValue, &powershellScriptRequest)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Create the new policy
	createdPolicy, err := client.CreateDeviceManagementScript(&powershellScriptRequest)
	if err != nil {
		fmt.Printf("Error creating policy: %s\n", err)
		return
	}

	fmt.Printf("Created Policy: %+v\n", createdPolicy)
}
