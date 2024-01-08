package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/http_client" // Import http_client for logging
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/clientauth.json"

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
	intune := &intuneSDK.Client{HTTP: httpClient}

	// Define the new script details
	newCustomConfigurationProfile := intuneSDK.ResourceiOSConfigurationProfile_CustomTemplate{
		DisplayName:       "IntuneSDK - Example custom iOS profile",
		Description:       "This is a test profile",
		PayloadName:       "ExamplePayloadName",
		PayloadFileName:   "ExamplePayloadFile.mobileconfig",
		Payload:           "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPCFET0NUWVBFIHBsaXN0IFBVQkxJQyAiLS8vQXBwbGUvL0RURCBQTElTVCAxLjAvL0VOIiAiaHR0cDovL3d3dy5hcHBsZS5jb20vRFREcy9Qcm9wZXJ0eUxpc3QtMS4wLmR0ZCI+CjxwbGlzdCB2ZXJzaW9uPSIxIj4KICAgIDxkaWN0PgogICAgICAgIDxrZXk+UGF5bG9hZFR5cGU8L2tleT4KICAgICAgICA8c3RyaW5nPkNvbmZpZ3VyYXRpb248L3N0cmluZz4KICAgICAgICA8a2V5PlBheWxvYWREaXNwbGF5TmFtZTwva2V5PgogICAgICAgIDxzdHJpbmc+YWNjZXNzaWJpbGl0eTwvc3RyaW5nPgogICAgICAgIDxrZXk+UGF5bG9hZERlc2NyaXB0aW9uPC9rZXk+CiAgICAgICAgPHN0cmluZy8+CiAgICAgICAgPGtleT5QYXlsb2FkVmVyc2lvbjwva2V5PgogICAgICAgIDxpbnRlZ2VyPjE8L2ludGVnZXI+CiAgICAgICAgPGtleT5QYXlsb2FkRW5hYmxlZDwva2V5PgogICAgICAgIDx0cnVlLz4KICAgICAgICA8a2V5PlBheWxvYWRSZW1vdmFsRGlzYWxsb3dlZDwva2V5PgogICAgICAgIDx0cnVlLz4KICAgICAgICA8a2V5PlBheWxvYWRTY29wZTwva2V5PgogICAgICAgIDxzdHJpbmc+U3lzdGVtPC9zdHJpbmc+CiAgICAgICAgPGtleT5QYXlsb2FkQ29udGVudDwva2V5PgogICAgICAgIDxhcnJheT4KICAgICAgICAgICAgPGRpY3Q+CiAgICAgICAgICAgICAgICA8a2V5PlBheWxvYWRUeXBlPC9rZXk+CiAgICAgICAgICAgICAgICA8c3RyaW5nPmNvbS5hcHBsZS51bml2ZXJzYWxhY2Nlc3M8L3N0cmluZz4KICAgICAgICAgICAgICAgIDxrZXk+UGF5bG9hZERpc3BsYXlOYW1lPC9rZXk+CiAgICAgICAgICAgICAgICA8c3RyaW5nPkFjY2Vzc2liaWxpdHk8L3N0cmluZz4KICAgICAgICAgICAgICAgIDxrZXk+UGF5bG9hZERlc2NyaXB0aW9uPC9rZXk+CiAgICAgICAgICAgICAgICA8c3RyaW5nLz4KICAgICAgICAgICAgICAgIDxrZXk+UGF5bG9hZFZlcnNpb248L2tleT4KICAgICAgICAgICAgICAgIDxpbnRlZ2VyPjE8L2ludGVnZXI+CiAgICAgICAgICAgICAgICA8a2V5PlBheWxvYWRFbmFibGVkPC9rZXk+CiAgICAgICAgICAgICAgICA8dHJ1ZS8+CiAgICAgICAgICAgICAgICA8a2V5PmNsb3NlVmlld1Njcm9sbFdoZWVsVG9nZ2xlPC9rZXk+CiAgICAgICAgICAgICAgICA8dHJ1ZS8+CiAgICAgICAgICAgICAgICA8a2V5PmNsb3NlVmlld0hvdGtleXNFbmFibGVkPC9rZXk+CiAgICAgICAgICAgICAgICA8dHJ1ZS8+CiAgICAgICAgICAgICAgICA8a2V5PmNsb3NlVmlld05lYXJQb2ludDwva2V5PgogICAgICAgICAgICAgICAgPGludGVnZXI+MTA8L2ludGVnZXI+CiAgICAgICAgICAgICAgICA8a2V5PmNsb3NlVmlld0ZhclBvaW50PC9rZXk+CiAgICAgICAgICAgICAgICA8aW50ZWdlcj4xPC9pbnRlZ2VyPgogICAgICAgICAgICAgICAgPGtleT5jbG9zZVZpZXdTaG93UHJldmlldzwva2V5PgogICAgICAgICAgICAgICAgPHRydWUvPgogICAgICAgICAgICAgICAgPGtleT5jbG9zZVZpZXdTbW9vdGhJbWFnZXM8L2tleT4KICAgICAgICAgICAgICAgIDx0cnVlLz4KICAgICAgICAgICAgICAgIDxrZXk+d2hpdGVPbkJsYWNrPC9rZXk+CiAgICAgICAgICAgICAgICA8dHJ1ZS8+CiAgICAgICAgICAgICAgICA8a2V5PmdyYXlzY2FsZTwva2V5PgogICAgICAgICAgICAgICAgPHRydWUvPgogICAgICAgICAgICAgICAgPGtleT5jb250cmFzdDwva2V5PgogICAgICAgICAgICAgICAgPGludGVnZXI+MDwvaW50ZWdlcj4KICAgICAgICAgICAgICAgIDxrZXk+bW91c2VEcml2ZXJDdXJzb3JTaXplPC9rZXk+CiAgICAgICAgICAgICAgICA8aW50ZWdlcj4xPC9pbnRlZ2VyPgogICAgICAgICAgICAgICAgPGtleT52b2ljZU92ZXJPbk9mZktleTwva2V5PgogICAgICAgICAgICAgICAgPHRydWUvPgogICAgICAgICAgICAgICAgPGtleT5mbGFzaFNjcmVlbjwva2V5PgogICAgICAgICAgICAgICAgPHRydWUvPgogICAgICAgICAgICAgICAgPGtleT5zdGVyZW9Bc01vbm88L2tleT4KICAgICAgICAgICAgICAgIDx0cnVlLz4KICAgICAgICAgICAgICAgIDxrZXk+c3RpY2t5S2V5PC9rZXk+CiAgICAgICAgICAgICAgICA8dHJ1ZS8+CiAgICAgICAgICAgICAgICA8a2V5PnN0aWNreUtleUJlZXBPbk1vZGlmaWVyPC9rZXk+CiAgICAgICAgICAgICAgICA8dHJ1ZS8+CiAgICAgICAgICAgICAgICA8a2V5PnN0aWNreUtleVNob3dXaW5kb3c8L2tleT4KICAgICAgICAgICAgICAgIDx0cnVlLz4KICAgICAgICAgICAgICAgIDxrZXk+c2xvd0tleTwva2V5PgogICAgICAgICAgICAgICAgPHRydWUvPgogICAgICAgICAgICAgICAgPGtleT5zbG93S2V5QmVlcE9uPC9rZXk+CiAgICAgICAgICAgICAgICA8dHJ1ZS8+CiAgICAgICAgICAgICAgICA8a2V5PnNsb3dLZXlEZWxheTwva2V5PgogICAgICAgICAgICAgICAgPGludGVnZXI+MDwvaW50ZWdlcj4KICAgICAgICAgICAgICAgIDxrZXk+bW91c2VEcml2ZXI8L2tleT4KICAgICAgICAgICAgICAgIDx0cnVlLz4KICAgICAgICAgICAgICAgIDxrZXk+bW91c2VEcml2ZXJJbml0aWFsRGVsYXk8L2tleT4KICAgICAgICAgICAgICAgIDxyZWFsPjEuMDwvcmVhbD4KICAgICAgICAgICAgICAgIDxrZXk+bW91c2VEcml2ZXJNYXhTcGVlZDwva2V5PgogICAgICAgICAgICAgICAgPGludGVnZXI+MzwvaW50ZWdlcj4KICAgICAgICAgICAgICAgIDxrZXk+bW91c2VEcml2ZXJJZ25vcmVUcmFja3BhZDwva2V5PgogICAgICAgICAgICAgICAgPGZhbHNlLz4KICAgICAgICAgICAgPC9kaWN0PgogICAgICAgIDwvYXJyYXk+CiAgICA8L2RpY3Q+CjwvcGxpc3Q+Cg==", // This should be your actual configuration payload
		Version:           1,
		RoleScopeTagIds:   []string{"0"},
		SupportsScopeTags: true,
	}

	// Create the new iOS configuration profile
	newScript, err := intune.CreateiOSConfigurationProfile_CustomTemplate(&newCustomConfigurationProfile)
	if err != nil {
		log.Fatalf("Failed to create custom iOS configuration profile: %v", err)
	}

	// Pretty print the created iOS configuration profile
	jsonData, err := json.MarshalIndent(newScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal created iOS configuration profile: %v", err)
	}
	fmt.Println(string(jsonData))
}
