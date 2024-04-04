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

	deviceEnrollmentConfigurationID := "be94fc43-03c5-4787-b42e-cfe57a24a7d8_PlatformRestrictions"

	// Use the Intune client to perform operations
	deviceEnrollmentConfigurationAssignments, err := intune.GetDeviceEnrollmentConfigurationAssignmentsByDeviceEnrollmentConfigurationID(deviceEnrollmentConfigurationID)
	if err != nil {
		log.Fatalf("Failed to get device management scripts: %v", err)
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(deviceEnrollmentConfigurationAssignments, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}
