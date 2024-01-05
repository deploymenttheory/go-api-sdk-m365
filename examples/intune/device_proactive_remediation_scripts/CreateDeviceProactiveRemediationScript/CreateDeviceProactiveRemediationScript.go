package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/http_client" // Import http_client for logging
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
	"github.com/deploymenttheory/go-api-sdk-m365/sdk/utils"
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

	// Base64 encode the detection script content
	detectionScriptBase64Encoded, err := utils.Base64Encode("/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/examples/intune/device_proactive_remediations/CreateProactiveRemediation/Template/Get-TemplateDetection.ps1")
	if err != nil {
		log.Fatalf("Failed to encode detection script: %v", err)
	}
	log.Printf("Detection Script Content (base64 encoded):\n%s\n", detectionScriptBase64Encoded)

	// Base64 encode the remediation script content
	remediationScriptBase64Encoded, err := utils.Base64Encode("/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-m365/examples/intune/device_proactive_remediations/CreateProactiveRemediation/Template/Get-TemplateRemediaton.ps1")
	if err != nil {
		log.Fatalf("Failed to encode remediation script: %v", err)
	}
	log.Printf("Remediation Script Content (base64 encoded):\n%s\n", remediationScriptBase64Encoded)

	// Define detection and remediation script parameters
	detectionParams := []intuneSDK.DeviceHealthScriptParameter{
		{
			Name:         "DetectionParam1",
			Description:  "Description for detection parameter 1",
			IsRequired:   true,
			DefaultValue: "false",
		},
	}

	remediationParams := []intuneSDK.DeviceHealthScriptParameter{
		{
			Name:         "RemediationParam1",
			Description:  "Description for remediation parameter 1",
			IsRequired:   true,
			DefaultValue: "true",
		},
	}

	// Example data for the Proactive Remediation Script
	remediationData := &intuneSDK.ResourceProactiveRemediation{
		Publisher:                   "Example Publisher",
		Version:                     "1.0",
		DisplayName:                 "intuneSDK - Example Proactive Remediation Script",
		Description:                 "This is a test script",
		DetectionScriptContent:      detectionScriptBase64Encoded,
		RemediationScriptContent:    remediationScriptBase64Encoded,
		RunAsAccount:                "system",
		EnforceSignatureCheck:       false,
		RunAs32Bit:                  false,
		RoleScopeTagIds:             []string{"0"},
		IsGlobalScript:              false,
		DeviceHealthScriptType:      "deviceHealthScript",
		DetectionScriptParameters:   detectionParams,
		RemediationScriptParameters: remediationParams,
	}

	// Create the Device Health Script
	createdRemediation, err := intune.CreateDeviceProactiveRemediationScript(remediationData)
	if err != nil {
		log.Fatalf("Error creating device health script: %v", err)
	}
	fmt.Printf("Created device health script: %+v\n", createdRemediation)
}
