package main

import (
	"fmt"
	"log"

	powershellScript "github.com/deploymenttheory/go-api-sdk-m365/export/intune/powershell_scripts"
	proactiveRemediation "github.com/deploymenttheory/go-api-sdk-m365/export/intune/proactive_remediations"
	"github.com/deploymenttheory/go-api-sdk-m365/sdk/http_client"
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {

	// Print ASCII art
	fmt.Println(`
  __    _____  ___  ___________  ____  ____  _____  ___    _______       _______     ______    
 |" \  (\"   \|"  \("     _   ")("  _||_ " |(\"   \|"  \  /"     "|     /" _   "|   /    " \   
 ||  | |.\\   \    |)__/  \\__/ |   (  ) : ||.\\   \    |(: ______)    (: ( \___)  // ____  \  
 |:  | |: \.   \\  |   \\_ /    (:  |  | . )|: \.   \\  | \/    |       \/ \      /  /    ) :) 
 |.  | |.  \    \. |   |.  |     \\ \__/ // |.  \    \. | // ___)_      //  \ ___(: (____/ //  
 /\  |\|    \    \ |   \:  |     /\\ __ //\ |    \    \ |(:      "|    (:   _(  _|\        /   
(__\_|_)\___|\____\)    \__|    (__________) \___|\____\) \_______)     \_______)  \"_____/    

    `)
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/clientauth.json"

	outputPath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/export/backup"
	outputFormat := "json"      // "json" or "yaml"
	excludeAssignments := false // or true, depending on your use case
	prefix := "[Intune]"        // replace with your desired prefix
	appendID := true            // or false

	// Load the client OAuth credentials from the configuration file
	clientAuthConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logger.SetLevel(http_client.LogLevelDebug) // Adjust the log level as needed

	// Configuration for the HTTP client
	httpClientconfig := http_client.Config{
		LogLevel:                  http_client.LogLevelDebug,
		MaxRetryAttempts:          3,
		EnableDynamicRateLimiting: true,
		Logger:                    logger,
		MaxConcurrentRequests:     5,
	}

	// initialize HTTP client instance
	httpClient, err := http_client.NewClient(httpClientconfig, clientAuthConfig, logger)
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}

	// Create an Intune client with the HTTP client
	intuneClient := &intuneSDK.Client{HTTP: httpClient}

	// Call backup functions based on user input or configuration
	if err := proactiveRemediation.Backup(intuneClient, outputPath, outputFormat, excludeAssignments, prefix, appendID); err != nil {
		fmt.Println("Error during Proactive Remediation backup:", err)
	}

	// Call backup functions based on user input or configuration
	if err := powershellScript.Backup(intuneClient, outputPath, outputFormat, excludeAssignments, prefix, appendID); err != nil {
		fmt.Println("Error during Proactive Remediation backup:", err)
	}

	// ... Call other backup functions similarly
}
