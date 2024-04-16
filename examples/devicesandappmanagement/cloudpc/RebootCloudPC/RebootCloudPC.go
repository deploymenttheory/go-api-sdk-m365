package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/client"
)

func main() {
	// Path to the JSON configuration file for initializing the msgraph client
	configFilePath := "/Users/dafyddwatkins/localtesting/msgraph/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration from the config file
	client, err := client.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize msgraph client: %v", err)
	}

	// Cloud PC ID to reboot
	cloudPCID := "12345678-1234-1234-1234-123456789012"

	// Call RebootCloudPC function to reboot the specified Cloud PC
	err = client.CloudPC.RebootCloudPC(cloudPCID)
	if err != nil {
		log.Fatalf("Error rebooting Cloud PC: %v", err)
	}

	fmt.Printf("Reboot command sent successfully to Cloud PC ID: %s\n", cloudPCID)
}
