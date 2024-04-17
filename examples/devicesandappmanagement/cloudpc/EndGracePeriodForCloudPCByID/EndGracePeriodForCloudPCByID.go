package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/msgraphclient"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/msgraph/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration from the config file
	client, err := msgraphclient.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize msgraph client: %v", err)
	}

	// Cloud PC ID to end the grace period for
	id := "12345678-1234-1234-1234-123456789012"

	// Attempt to end the grace period for the specified Cloud PC
	err = client.CloudPC.EndGracePeriodForCloudPCByID(id)
	if err != nil {
		log.Fatalf("Error ending grace period for Cloud PC: %v", err)
	}

	fmt.Println("Grace period ended successfully for Cloud PC ID:", id)
}
