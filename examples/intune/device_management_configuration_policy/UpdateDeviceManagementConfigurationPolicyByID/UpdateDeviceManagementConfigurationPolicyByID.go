// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicy",
// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
package main

import (
	"fmt"
	"log"
	"time"

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

	// Define the new settings
	policySettings := []intuneSDK.DeviceManagementConfigurationSubsetSetting{
		{
			ID: "0",
			SettingInstance: intuneSDK.DeviceManagementConfigurationSubsetSettingInstance{
				OdataType:           "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance",
				SettingDefinitionId: "user_vendor_msft_policy_config_teamsv2~policy~l_teams_teams_preventfirstlaunchafterinstall_policy",
				ChoiceSettingValue: &intuneSDK.DeviceManagementConfigurationSubsetChoiceSettingValue{
					Value:    "user_vendor_msft_policy_config_teamsv2~policy~l_teams_teams_preventfirstlaunchafterinstall_policy_1",
					Children: []intuneSDK.DeviceManagementConfigurationSubsetSettingInstance{},
				},
			},
		},
		{
			ID: "1",
			SettingInstance: intuneSDK.DeviceManagementConfigurationSubsetSettingInstance{
				OdataType:           "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance",
				SettingDefinitionId: "user_vendor_msft_policy_config_teamsv3~policy~l_teams_string_teams_signinrestriction_policy",
				ChoiceSettingValue: &intuneSDK.DeviceManagementConfigurationSubsetChoiceSettingValue{
					Value: "user_vendor_msft_policy_config_teamsv3~policy~l_teams_string_teams_signinrestriction_policy_1",
					Children: []intuneSDK.DeviceManagementConfigurationSubsetSettingInstance{
						{
							OdataType:           "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstance",
							SettingDefinitionId: "user_vendor_msft_policy_config_teamsv3~policy~l_teams_string_teams_signinrestriction_policy_restrictteamssignintoaccountsfromtenantlist",
							SimpleSettingValue: &intuneSDK.DeviceManagementConfigurationSubsetSimpleSettingValue{
								OdataType: "#microsoft.graph.deviceManagementConfigurationStringSettingValue",
								Value:     "2fd6bb84-ad40-4ec5-9369-a215b25c9952",
							},
						},
					},
				},
			},
		},
	}

	// Create the policy request based on the new payload
	policyRequest := intuneSDK.ResourceDeviceManagementConfigurationPolicy{
		OdataType:            "#microsoft.graph.deviceManagementConfigurationPolicy",
		ID:                   "24b18871-d50f-43fb-b0dc-8c88b7901cdf",
		Name:                 "intuneSDK | Windows - Settings Catalog | Microsoft Teams ver0.1",
		Description:          "09.06.2022\nContext: User",
		Platforms:            "windows10",
		Technologies:         "mdm",
		CreatedDateTime:      time.Now(),
		LastModifiedDateTime: time.Now(),
		SettingCount:         2,
		CreationSource:       "",
		RoleScopeTagIds:      []string{"0"},
		IsAssigned:           false,
		TemplateReference: intuneSDK.DeviceManagementConfigurationPolicySubsetTemplateReference{
			OdataType:      "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
			TemplateId:     "",
			TemplateFamily: "none",
		},
		Settings: policySettings,
	}

	// Example policy ID to get
	policyID := "8077bf4b-2677-4521-b839-549396b052b1"

	// Create the new policy
	createdPolicy, err := intune.UpdateDeviceManagementConfigurationPolicyByID(policyID, &policyRequest)
	if err != nil {
		fmt.Printf("Error creating policy: %s\n", err)
		return
	}

	fmt.Printf("Created Policy: %+v\n", createdPolicy)
}
