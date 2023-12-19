package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-intune/sdk/http_client"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-intune/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Debug print statement to check the loaded configuration
	fmt.Printf("Loaded Config: %+v\n", authConfig)

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	loggingConfig := http_client.Config{
		LogLevel: logLevel,
		Logger:   logger,
	}

	// Create a new client
	client, err := http_client.NewClient(loggingConfig, authConfig, logger)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

	// Call the ValidAuthTokenCheck function to ensure that a valid token is set in the client
	isTokenValid, err := client.ValidAuthTokenCheck()
	if err != nil {
		log.Fatalf("Error while validating token: %v", err)
	}
	if !isTokenValid {
		fmt.Println("Error obtaining or refreshing token.")
		return
	}

	// Print the obtained token
	fmt.Println("Successfully obtained token:", client.Token)
}
