package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/client"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/msgraph/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration from the config file
	client, err := client.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize msgraph client: %v", err)
	}

	id := "12345678-1234-1234-1234-123456789012"

	// Call ListCloudPCs function to get a list of Cloud PCs
	cloudPCs, err := client.CloudPC.GetCloudPCByID(id)
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
