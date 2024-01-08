// graphbeta_device_management_reuseable_policy_settings.go
// Graph Beta Api - Intune: Configuration Profiles with reuseable settings
// Documentation: https://learn.microsoft.com/en-us/mem/intune/protect/reusable-settings-groups
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/configProfiles
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfigv2-devicemanagementreusablepolicysetting?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const uriBetaDeviceManagementReusablePolicySettings = "/beta/deviceManagement/reusablePolicySettings"

// ResourceDeviceManagementReusablePolicySetting represents a list of reusable policy settings in device management.
type ResponseDeviceManagementReusablePolicySettingsList struct {
	ODataContext string                                          `json:"@odata.context"`
	ODataCount   int                                             `json:"@odata.count"`
	Value        []ResourceDeviceManagementReusablePolicySetting `json:"value"`
}

// ResourceDeviceManagementReusablePolicySetting represents a reusable policy setting resource in device management.
type ResourceDeviceManagementReusablePolicySetting struct {
	OdataType                           string                                              `json:"@odata.type"`
	ID                                  string                                              `json:"id"`
	DisplayName                         string                                              `json:"displayName"`
	Description                         string                                              `json:"description"`
	SettingDefinitionId                 string                                              `json:"settingDefinitionId"`
	SettingInstance                     *DeviceManagementConfigurationChoiceSettingInstance `json:"settingInstance,omitempty"`
	CreatedDateTime                     time.Time                                           `json:"createdDateTime"`
	LastModifiedDateTime                time.Time                                           `json:"lastModifiedDateTime"`
	Version                             int                                                 `json:"version"`
	ReferencingConfigurationPolicyCount int                                                 `json:"referencingConfigurationPolicyCount"`
}

// DeviceManagementConfigurationChoiceSettingInstance represents an instance of a choice setting.
type DeviceManagementConfigurationChoiceSettingInstance struct {
	OdataType                        string                                                         `json:"@odata.type"`
	SettingDefinitionId              string                                                         `json:"settingDefinitionId"`
	SettingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference `json:"settingInstanceTemplateReference,omitempty"`
	ChoiceSettingValue               *DeviceManagementConfigurationChoiceSettingValue               `json:"choiceSettingValue,omitempty"`
}

// DeviceManagementConfigurationSettingInstanceTemplateReference represents a reference to a setting instance template.
type DeviceManagementConfigurationSettingInstanceTemplateReference struct {
	OdataType                 string `json:"@odata.type"`
	SettingInstanceTemplateId string `json:"settingInstanceTemplateId"`
}

// DeviceManagementConfigurationChoiceSettingValue represents the value of a choice setting.
type DeviceManagementConfigurationChoiceSettingValue struct {
	OdataType                     string                                                      `json:"@odata.type"`
	SettingValueTemplateReference *DeviceManagementConfigurationSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`
	Value                         string                                                      `json:"value"`
	Children                      []*DeviceManagementConfigurationChoiceSettingInstance       `json:"children,omitempty"`
}

// DeviceManagementConfigurationSettingValueTemplateReference represents a template reference for a setting value.
type DeviceManagementConfigurationSettingValueTemplateReference struct {
	OdataType              string `json:"@odata.type"`
	SettingValueTemplateId string `json:"settingValueTemplateId"`
	UseTemplateDefault     bool   `json:"useTemplateDefault"`
}

// GetResourceDeviceManagementReusablePolicySettings retrieves a list of all device management reusable policy settings.
func (c *Client) GetResourceDeviceManagementReusablePolicySettings() ([]ResourceDeviceManagementReusablePolicySetting, error) {
	endpoint := uriBetaDeviceManagementReusablePolicySettings

	var responseReusablePolicySettings ResponseDeviceManagementReusablePolicySettingsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseReusablePolicySettings)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management reusable policy settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return responseReusablePolicySettings.Value, nil
}

// GetDeviceManagementReusablePolicySettingByID retrieves a specific device management Reusable Policy Setting by its ID.
func (c *Client) GetDeviceManagementReusablePolicySettingByID(policySettingId string) (*ResourceDeviceManagementReusablePolicySetting, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementReusablePolicySettings, policySettingId)

	var responseReusablePolicySetting ResourceDeviceManagementReusablePolicySetting
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseReusablePolicySetting)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device management reusable policy setting", policySettingId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseReusablePolicySetting, nil
}

//TODO - rest of the CRUD functions
