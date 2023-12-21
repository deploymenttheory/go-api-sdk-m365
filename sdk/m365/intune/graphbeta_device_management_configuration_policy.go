// graphbeta_device_management_configuration_policy.go
// Graph Beta Api - Intune: Configuration Profiles
// Documentation: https://learn.microsoft.com/en-us/mem/intune/configuration/device-profile-create
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/configProfiles
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-devicemanagementscriptgroupassignment?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
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
type ResponseDeviceManagementConfigurationPoliciesList struct {
	ODataContext string                                        `json:"@odata.context"`
	ODataCount   int                                           `json:"@odata.count"`
	Value        []ResourceDeviceManagementConfigurationPolicy `json:"value"`
}

// ResourceDeviceManagementConfigurationPolicy represents a device management configuration policy.
type ResourceDeviceManagementConfigurationPolicy struct {
	OdataType            string                                                     `json:"@odata.type"`
	ID                   string                                                     `json:"id"`
	Name                 string                                                     `json:"name"`
	Description          string                                                     `json:"description"`
	Platforms            string                                                     `json:"platforms"`
	Technologies         string                                                     `json:"technologies"`
	CreatedDateTime      time.Time                                                  `json:"createdDateTime"`
	LastModifiedDateTime time.Time                                                  `json:"lastModifiedDateTime"`
	SettingCount         int                                                        `json:"settingCount"`
	CreationSource       string                                                     `json:"creationSource"`
	RoleScopeTagIds      []string                                                   `json:"roleScopeTagIds"`
	IsAssigned           bool                                                       `json:"isAssigned"`
	TemplateReference    DeviceManagementConfigurationPolicySubsetTemplateReference `json:"templateReference"`
	PriorityMetaData     *DeviceManagementSubsetPriorityMetaData                    `json:"priorityMetaData,omitempty"`
	Settings             []DeviceManagementConfigurationSubsetSetting               `json:"settings"`
}

// DeviceManagementConfigurationPolicyTemplateReference represents the template reference in a configuration policy.
type DeviceManagementConfigurationPolicySubsetTemplateReference struct {
	OdataType              string `json:"@odata.type"`
	TemplateId             string `json:"templateId"`
	TemplateFamily         string `json:"templateFamily"`
	TemplateDisplayName    string `json:"templateDisplayName,omitempty"`
	TemplateDisplayVersion string `json:"templateDisplayVersion,omitempty"`
}

// DeviceManagementPriorityMetaData represents the priority metadata in a configuration policy.
type DeviceManagementSubsetPriorityMetaData struct {
	OdataType string `json:"@odata.type"`
	Priority  int    `json:"priority"`
}

// DeviceManagementConfigurationSetting represents a configuration settings within a configuration policy.
type DeviceManagementConfigurationSubsetSetting struct {
	ID              string                                             `json:"id"`
	SettingInstance DeviceManagementConfigurationSubsetSettingInstance `json:"settingInstance"`
}

// DeviceManagementConfigurationSettingInstance represents an instance of a configuration setting.
type DeviceManagementConfigurationSubsetSettingInstance struct {
	OdataType                        string                                                       `json:"@odata.type"`
	SettingDefinitionId              string                                                       `json:"settingDefinitionId"`
	SettingInstanceTemplateReference *DeviceManagementConfigurationSubsetSettingInstanceReference `json:"settingInstanceTemplateReference,omitempty"`
	ChoiceSettingValue               *DeviceManagementConfigurationSubsetChoiceSettingValue       `json:"choiceSettingValue,omitempty"`
	SimpleSettingValue               *DeviceManagementConfigurationSubsetSimpleSettingValue       `json:"simpleSettingValue,omitempty"`
}

// DeviceManagementConfigurationSettingInstanceReference represents a reference to a setting instance.
type DeviceManagementConfigurationSubsetSettingInstanceReference struct {
	SettingInstanceTemplateId string `json:"settingInstanceTemplateId,omitempty"`
}

// DeviceManagementConfigurationChoiceSettingValue represents the value of a choice setting.
type DeviceManagementConfigurationSubsetChoiceSettingValue struct {
	Value                         string                                               `json:"value"`
	Children                      []DeviceManagementConfigurationSubsetSettingInstance `json:"children"`
	SettingValueTemplateReference *DeviceManagementSettingValueTemplateReference       `json:"settingValueTemplateReference,omitempty"`
}

// DeviceManagementConfigurationSimpleSettingValue represents the value of a simple setting.
type DeviceManagementConfigurationSubsetSimpleSettingValue struct {
	OdataType                     string                                         `json:"@odata.type"`
	SettingValueTemplateReference *DeviceManagementSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`
	Value                         interface{}                                    `json:"value"`
}

type DeviceManagementSettingValueTemplateReference struct {
	SettingValueTemplateId string `json:"settingValueTemplateId,omitempty"`
	UseTemplateDefault     bool   `json:"useTemplateDefault,omitempty"`
}

// GetDeviceManagementConfigurationPolicies retrieves a list of all device management configuration policies.
func (c *Client) GetDeviceManagementConfigurationPolicies() (*ResponseDeviceManagementConfigurationPoliciesList, error) {
	endpoint := uriBetaDeviceManagementConfigurationPolicies

	var responseConfigurationPolicies ResponseDeviceManagementConfigurationPoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseConfigurationPolicies)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management configuration policies", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseConfigurationPolicies, nil
}

// GetDeviceManagementConfigurationPolicyByID retrieves a specific device management configuration policy by its ID.
func (c *Client) GetDeviceManagementConfigurationPolicyByID(policyId string) (*ResourceDeviceManagementConfigurationPolicy, error) {
	endpoint := fmt.Sprintf("%s('%s')?$expand=settings", uriBetaDeviceManagementConfigurationPolicies, policyId)

	var responseDeviceManagementConfigurationPolicy ResourceDeviceManagementConfigurationPolicy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseDeviceManagementConfigurationPolicy)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device management configuration policy", policyId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseDeviceManagementConfigurationPolicy, nil
}

// GetDeviceManagementConfigurationPolicyByName retrieves a specific device management configuration policy by its name.
func (c *Client) GetDeviceManagementConfigurationPolicyByName(policyName string) (*ResourceDeviceManagementConfigurationPolicy, error) {
	// Retrieve all policies
	policiesList, err := c.GetDeviceManagementConfigurationPolicies()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management configuration policies", err)
	}

	// Search for the policy with the matching name
	var policyID string
	for _, policy := range policiesList.Value {
		if policy.Name == policyName {
			policyID = policy.ID
			break
		}
	}

	if policyID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device management configuration policy", policyName, "Policy not found")
	}
	// Retrieve the full details of the policy using its ID
	return c.GetDeviceManagementConfigurationPolicyByID(policyID)
}

// CreateDeviceManagementConfigurationPolicy creates a new device management configuration policy.
func (c *Client) CreateDeviceManagementConfigurationPolicy(request *ResourceDeviceManagementConfigurationPolicy) (*ResourceDeviceManagementConfigurationPolicy, error) {
	endpoint := uriBetaDeviceManagementConfigurationPolicies

	var responseCreatedPolicy ResourceDeviceManagementConfigurationPolicy
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &responseCreatedPolicy)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management configuration policy", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseCreatedPolicy, nil
}

// CreateCopyOfDeviceManagementConfigurationPolicyByID creates a copy of an existing device management configuration policy.
func (c *Client) CreateCopyOfDeviceManagementConfigurationPolicyByID(sourcePolicyId string, copyDisplayName string, copyDescription string) (*ResourceDeviceManagementConfigurationPolicy, error) {
	// Construct the endpoint URL using the existing constant
	endpoint := fmt.Sprintf("%s/%s/createCopy", uriBetaDeviceManagementConfigurationPolicies, sourcePolicyId)

	// Create the request body
	requestBody := map[string]string{
		"displayName": copyDisplayName,
		"description": copyDescription,
	}

	var responseCopyPolicy ResourceDeviceManagementConfigurationPolicy
	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, &responseCopyPolicy)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreateCopy, "device management configuration policy", sourcePolicyId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseCopyPolicy, nil
}

// CreateCopyOfDeviceManagementConfigurationPolicyByName creates a copy of an existing device management configuration policy by its name.
func (c *Client) CreateCopyOfDeviceManagementConfigurationPolicyByName(sourcePolicyName string, copyDisplayName string, copyDescription string) (*ResourceDeviceManagementConfigurationPolicy, error) {
	// Retrieve all policies
	policiesList, err := c.GetDeviceManagementConfigurationPolicies()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management configuration policies", err)
	}

	// Search for the policy with the matching name
	var policyID string
	for _, policy := range policiesList.Value {
		if policy.Name == sourcePolicyName {
			policyID = policy.ID
			break
		}
	}

	if policyID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device management configuration policy", sourcePolicyName, "Policy not found")
	}

	// Create a copy of the policy using its ID
	return c.CreateCopyOfDeviceManagementConfigurationPolicyByID(policyID, copyDisplayName, copyDescription)
}

// UpdateDeviceManagementConfigurationPolicyByID updates an existing device management configuration policy by its ID.
func (c *Client) UpdateDeviceManagementConfigurationPolicyByID(policyId string, request *ResourceDeviceManagementConfigurationPolicy) (*ResourceDeviceManagementConfigurationPolicy, error) {
	// Construct the endpoint URL using the existing constant
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementConfigurationPolicies, policyId)

	// Exclude 'settings' from the request object. Cannot apply PATCH to navigation property 'settings'
	// on entity type 'microsoft.management.services.api.deviceManagementConfigurationPolicy'
	request.Settings = nil

	var responseUpdatedPolicy ResourceDeviceManagementConfigurationPolicy
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, &responseUpdatedPolicy)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedUpdate, "device management configuration policy", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseUpdatedPolicy, nil
}

// ReorderDeviceManagementConfigurationPolicyByID updates the priority of a device management configuration policy.
func (c *Client) ReorderDeviceManagementConfigurationPolicyByID(policyId string, newPriority int) (*ResourceDeviceManagementConfigurationPolicy, error) {
	endpoint := fmt.Sprintf("%s/%s/reorder", uriBetaDeviceManagementConfigurationPolicies, policyId)

	// Create the request body using a map
	requestBody := map[string]int{
		"priority": newPriority,
	}

	var reorderedPolicy ResourceDeviceManagementConfigurationPolicy
	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, &reorderedPolicy)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedReorder, "device management configuration policy", policyId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &reorderedPolicy, nil
}

// DeleteDeviceManagementConfigurationPolicyByID deletes a device management configuration policy by its ID.
func (c *Client) DeleteDeviceManagementConfigurationPolicyByID(policyId string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementConfigurationPolicies, policyId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedDeleteByID, "device management configuration policy", policyId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDeviceManagementConfigurationPolicyByName deletes a device management configuration policy by its name.
func (c *Client) DeleteDeviceManagementConfigurationPolicyByName(policyName string) error {
	policiesList, err := c.GetDeviceManagementConfigurationPolicies()
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGet, "device management configuration policies", err)
	}

	var policyID string
	for _, policy := range policiesList.Value {
		if policy.Name == policyName {
			policyID = policy.ID
			break
		}
	}

	if policyID == "" {
		return fmt.Errorf(shared.ErrorMsgFailedGetByName, "device management configuration policy", policyName, "Policy not found")
	}

	err = c.DeleteDeviceManagementConfigurationPolicyByID(policyID)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedDeleteByName, "device management configuration policy", policyName, err)
	}

	return nil
}
