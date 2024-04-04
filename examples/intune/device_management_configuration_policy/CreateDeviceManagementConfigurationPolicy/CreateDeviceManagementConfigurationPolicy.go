// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicy",
// "@odata.type": "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
package main

import (
	"fmt"
	"log"
	"time"

	// Import http_client for logging

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

	// Define the new settings
	policySettings := []intune.DeviceManagementConfigurationSubsetSetting{
		{
			ID: "0",
			SettingInstance: intune.DeviceManagementConfigurationSubsetSettingInstance{
				OdataType:           "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance",
				SettingDefinitionId: "user_vendor_msft_policy_config_teamsv2~policy~l_teams_teams_preventfirstlaunchafterinstall_policy",
				ChoiceSettingValue: &intune.DeviceManagementConfigurationSubsetChoiceSettingValue{
					Value:    "user_vendor_msft_policy_config_teamsv2~policy~l_teams_teams_preventfirstlaunchafterinstall_policy_1",
					Children: []intune.DeviceManagementConfigurationSubsetSettingInstance{},
				},
			},
		},
		{
			ID: "1",
			SettingInstance: intune.DeviceManagementConfigurationSubsetSettingInstance{
				OdataType:           "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance",
				SettingDefinitionId: "user_vendor_msft_policy_config_teamsv3~policy~l_teams_string_teams_signinrestriction_policy",
				ChoiceSettingValue: &intune.DeviceManagementConfigurationSubsetChoiceSettingValue{
					Value: "user_vendor_msft_policy_config_teamsv3~policy~l_teams_string_teams_signinrestriction_policy_1",
					Children: []intune.DeviceManagementConfigurationSubsetSettingInstance{
						{
							OdataType:           "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstance",
							SettingDefinitionId: "user_vendor_msft_policy_config_teamsv3~policy~l_teams_string_teams_signinrestriction_policy_restrictteamssignintoaccountsfromtenantlist",
							SimpleSettingValue: &intune.DeviceManagementConfigurationSubsetSimpleSettingValue{
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
	policyRequest := intune.ResourceDeviceManagementConfigurationPolicy{
		OdataType:            "#microsoft.graph.deviceManagementConfigurationPolicy",
		ID:                   "24b18871-d50f-43fb-b0dc-8c88b7901cdf",
		Name:                 "intune | Windows - Settings Catalog | Microsoft Teams ver0.1",
		Description:          "09.06.2022\nContext: User",
		Platforms:            "windows10",
		Technologies:         "mdm",
		CreatedDateTime:      time.Now(),
		LastModifiedDateTime: time.Now(),
		SettingCount:         2,
		CreationSource:       "",
		RoleScopeTagIds:      []string{"0"},
		IsAssigned:           false,
		TemplateReference: intune.DeviceManagementConfigurationPolicySubsetTemplateReference{
			OdataType:      "#microsoft.graph.deviceManagementConfigurationPolicyTemplateReference",
			TemplateId:     "",
			TemplateFamily: "none",
		},
		Settings: policySettings,
	}

	// Create the new policy
	createdPolicy, err := client.CreateDeviceManagementConfigurationPolicy(&policyRequest)
	if err != nil {
		fmt.Printf("Error creating policy: %s\n", err)
		return
	}

	fmt.Printf("Created Policy: %+v\n", createdPolicy)
}
