// graphbeta_device_management_group_policy_configurations.go
// Graph Beta Api - Intune: Group Policy Configuration Profiles
// Documentation: https://learn.microsoft.com/en-us/mem/intune/configuration/group-policy-analytics
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/configProfiles
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-grouppolicy-grouppolicyconfiguration?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

// Constant for the endpoint URL
const uriBetaDeviceManagementGroupPolicyConfigurations = "/beta/deviceManagement/groupPolicyConfigurations"

// Struct for the list response
type ResponseDeviceManagementGroupPolicyConfigurationsList struct {
	ODataContext string                                             `json:"@odata.context"`
	Value        []ResourceDeviceManagementGroupPolicyConfiguration `json:"value"`
}

// Struct for individual Group Policy Configuration
type ResourceDeviceManagementGroupPolicyConfiguration struct {
	OdataType                        string    `json:"@odata.type"`
	ID                               string    `json:"id"`
	DisplayName                      string    `json:"displayName"`
	Description                      string    `json:"description"`
	RoleScopeTagIds                  []string  `json:"roleScopeTagIds"`
	PolicyConfigurationIngestionType string    `json:"policyConfigurationIngestionType"`
	CreatedDateTime                  time.Time `json:"createdDateTime"`
	LastModifiedDateTime             time.Time `json:"lastModifiedDateTime"`
}

// Function to get the list of Group Policy Configurations
func (c *Client) GetDeviceManagementGroupPolicyConfigurations() ([]ResourceDeviceManagementGroupPolicyConfiguration, error) {
	endpoint := uriBetaDeviceManagementGroupPolicyConfigurations

	var responseGroupPolicyConfigurations ResponseDeviceManagementGroupPolicyConfigurationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseGroupPolicyConfigurations)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management group policy configurations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return responseGroupPolicyConfigurations.Value, nil
}

// GetDeviceManagementGroupPolicyConfigurationByID retrieves a specific Group Policy Configuration by its ID.
func (c *Client) GetDeviceManagementGroupPolicyConfigurationByID(policyConfigurationId string) (*ResourceDeviceManagementGroupPolicyConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementGroupPolicyConfigurations, policyConfigurationId)

	var responseGroupPolicyConfiguration ResourceDeviceManagementGroupPolicyConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseGroupPolicyConfiguration)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device management group policy configuration", policyConfigurationId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseGroupPolicyConfiguration, nil
}
