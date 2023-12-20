// graphbeta_device_management_configuration_policy.go
// Graph Beta Api - Intune: Configuration Profiles
// Documentation: https://learn.microsoft.com/en-us/mem/intune/configuration/device-profile-create
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/configProfiles
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-devicemanagementscriptgroupassignment?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const uriBetaDeviceManagementConfigurationPolicies = "/beta/deviceManagement/configurationPolicies"
const uriDeviceManagementConfigurationPolicies = "/deviceManagement/configurationPolicies"

// ResourceDeviceManagementConfigurationPoliciesList represents the response structure for configuration policies.
type ResourceDeviceManagementConfigurationPoliciesList struct {
	Value []DeviceManagementConfigurationPolicy `json:"value"`
}

// DeviceManagementConfigurationPolicy represents a configuration policy.
type DeviceManagementConfigurationPolicy struct {
	OdataType            string                                               `json:"@odata.type"`
	ID                   string                                               `json:"id"`
	Name                 string                                               `json:"name"`
	Description          string                                               `json:"description"`
	Platforms            string                                               `json:"platforms"`
	Technologies         string                                               `json:"technologies"`
	CreatedDateTime      time.Time                                            `json:"createdDateTime"`
	LastModifiedDateTime time.Time                                            `json:"lastModifiedDateTime"`
	SettingCount         int                                                  `json:"settingCount"`
	CreationSource       string                                               `json:"creationSource"`
	RoleScopeTagIds      []string                                             `json:"roleScopeTagIds"`
	IsAssigned           bool                                                 `json:"isAssigned"`
	TemplateReference    DeviceManagementConfigurationPolicyTemplateReference `json:"templateReference"`
	PriorityMetaData     *DeviceManagementPriorityMetaData                    `json:"priorityMetaData,omitempty"`
	Settings             []DeviceManagementConfigurationSetting               `json:"settings"`
}

// DeviceManagementConfigurationPolicyTemplateReference represents the template reference in a configuration policy.
type DeviceManagementConfigurationPolicyTemplateReference struct {
	OdataType              string `json:"@odata.type"`
	TemplateId             string `json:"templateId"`
	TemplateFamily         string `json:"templateFamily"`
	TemplateDisplayName    string `json:"templateDisplayName,omitempty"`
	TemplateDisplayVersion string `json:"templateDisplayVersion,omitempty"`
}

// DeviceManagementPriorityMetaData represents the priority metadata in a configuration policy.
type DeviceManagementPriorityMetaData struct {
	OdataType string `json:"@odata.type"`
	Priority  int    `json:"priority"`
}

// DeviceManagementConfigurationSetting represents a configuration setting within a configuration policy.
type DeviceManagementConfigurationSetting struct {
	ID              string                                       `json:"id"`
	SettingInstance DeviceManagementConfigurationSettingInstance `json:"settingInstance"`
}

// DeviceManagementConfigurationSettingInstance represents an instance of a configuration setting.
type DeviceManagementConfigurationSettingInstance struct {
	OdataType                        string                                                 `json:"@odata.type"`
	SettingDefinitionId              string                                                 `json:"settingDefinitionId"`
	SettingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceReference `json:"settingInstanceTemplateReference,omitempty"`
	ChoiceSettingValue               DeviceManagementConfigurationChoiceSettingValue        `json:"choiceSettingValue"`
}

// DeviceManagementConfigurationSettingInstanceReference represents a reference to a setting instance.
type DeviceManagementConfigurationSettingInstanceReference struct {
	// Define fields if needed
}

// DeviceManagementConfigurationChoiceSettingValue represents the value of a choice setting.
type DeviceManagementConfigurationChoiceSettingValue struct {
	Value    string                                         `json:"value"`
	Children []DeviceManagementConfigurationSettingInstance `json:"children"`
}

// GetDeviceManagementConfigurationPolicies retrieves a list of all device management configuration policies.
func (c *Client) GetDeviceManagementConfigurationPolicies() (*ResourceDeviceManagementConfigurationPoliciesList, error) {
	endpoint := uriBetaDeviceManagementConfigurationPolicies

	var configurationPolicies ResourceDeviceManagementConfigurationPoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &configurationPolicies)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management configuration policies", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &configurationPolicies, nil
}

// GetDeviceManagementConfigurationPolicyByID retrieves a specific device management configuration policy by its ID.
func (c *Client) GetDeviceManagementConfigurationPolicyByID(policyId string) (*DeviceManagementConfigurationPolicy, error) {
	endpoint := fmt.Sprintf("%s('%s')?$expand=settings", uriBetaDeviceManagementConfigurationPolicies, policyId)

	var policy DeviceManagementConfigurationPolicy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policy)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device management configuration policy", policyId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policy, nil
}
