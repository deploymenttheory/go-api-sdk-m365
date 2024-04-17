package main

import (
	"encoding/json"
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

	// Call ListCloudPCs function to get a list of Cloud PCs
	cloudPCs, err := client.CloudPC.ListCloudPCs()
	if err != nil {
		log.Fatalf("Error listing Cloud PCs: %v", err)
	}

	// Pretty print the Cloud PCs in JSON format
	cloudPCsJSON, err := json.MarshalIndent(cloudPCs, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Cloud PCs to JSON: %v", err)
	}
	fmt.Println("Cloud PCs:", string(cloudPCsJSON))
}
