// graphbeta_device_management_iOS_template_custom_configuration_profiles.go
// Graph Beta Api - Intune: iOS configuration profiles (Custom)
// Documentation: https://learn.microsoft.com/en-us/mem/intune/configuration/custom-settings-ios
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesIosMenu/~/configProfiles
// API reference: https://learn.microsoft.com/en-us/graph/api/intune-shared-deviceconfiguration-list?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"strings"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriBetaGraphiOSDeviceManagementDeviceConfiguration = "/beta/deviceManagement/deviceConfigurations"
	//odataTypeMacOSCustomConfigurationProfile           = "#microsoft.graph.macOSCustomConfiguration"
	odataTypeiOSCustomConfigurationProfile = "#microsoft.graph.iosCustomConfiguration"
	//odataTypeiOSTemplateConfigurationProfile           = "#microsoft.graph.iosGeneralDeviceConfiguration"
)

// ResourceiOSConfigurationProfile_CustomTemplateList represents a response containing a list of Device Configuration Profiles.
type ResourceiOSConfigurationProfile_CustomTemplateList struct {
	ODataContext string                                           `json:"@odata.context,omitempty"`
	Value        []ResourceiOSConfigurationProfile_CustomTemplate `json:"value,omitempty"`
}

// ResourceiOSConfigurationProfile_CustomTemplate represents a single windows device configuration profile template.
type ResourceiOSConfigurationProfile_CustomTemplate struct {
	ODataType                                   string                                        `json:"@odata.type,omitempty"`
	ID                                          string                                        `json:"id,omitempty"`
	CreatedDateTime                             string                                        `json:"createdDateTime,omitempty"`
	LastModifiedDateTime                        string                                        `json:"lastModifiedDateTime,omitempty"`
	Description                                 string                                        `json:"description,omitempty"`
	DisplayName                                 string                                        `json:"displayName,omitempty"`
	Version                                     int                                           `json:"version,omitempty"`
	RoleScopeTagIds                             []string                                      `json:"roleScopeTagIds,omitempty"`
	SupportsScopeTags                           bool                                          `json:"supportsScopeTags,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition   `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion   `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode  `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	PayloadName                                 string                                        `json:"payloadName,omitempty"`
	PayloadFileName                             string                                        `json:"payloadFileName,omitempty"`
	Payload                                     string                                        `json:"payload,omitempty"`
	AssignmentsOdataContext                     string                                        `json:"assignments@odata.context,omitempty"`
	Assignments                                 []CustomConfigurationProfileSubsetAssignments `json:"assignments,omitempty"`
}

// CustomConfigurationProfileSubsetAssignments represents an assignment of a Device Management Script.
type CustomConfigurationProfileSubsetAssignments struct {
	ID     string                  `json:"id,omitempty"`
	Target AssignmentsSubsetTarget `json:"target,omitempty"`
}

// AssignmentsSubsetTarget represents the target of a script assignment.
type AssignmentsSubsetTarget struct {
	ODataType                                  string `json:"@odata.type,omitempty"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId,omitempty"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType,omitempty"`
	CollectionId                               string `json:"collectionId,omitempty"`
}

// GetiOSConfigurationProfiles_CustomTemplates retrieves a list of iOS device configuration profiles from Microsoft Graph API.
// Because this is a shared endpoint, an OdataType match is used to filter the response so that only windows configuration
// profiles are returned.
func (c *Client) GetiOSConfigurationProfiles_CustomTemplates() (*ResourceiOSConfigurationProfile_CustomTemplateList, error) {
	endpoint := uriBetaGraphiOSDeviceManagementDeviceConfiguration + "?$expand=assignments"

	var responseiOSDeviceConfigurationProfiles ResourceiOSConfigurationProfile_CustomTemplateList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseiOSDeviceConfigurationProfiles)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "iOS configuration profiles", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Filter to include only iOS device profiles
	var iOSCustomConfigurationProfiles []ResourceiOSConfigurationProfile_CustomTemplate
	for _, profile := range responseiOSDeviceConfigurationProfiles.Value {
		if strings.HasPrefix(profile.ODataType, odataTypeiOSCustomConfigurationProfile) {
			iOSCustomConfigurationProfiles = append(iOSCustomConfigurationProfiles, profile)
		}
	}

	responseiOSDeviceConfigurationProfiles.Value = iOSCustomConfigurationProfiles
	return &responseiOSDeviceConfigurationProfiles, nil
}

// GetiOSConfigurationProfileByID_CustomTemplate retrieves a Device Management Script by its ID.
func (c *Client) GetiOSConfigurationProfileByID_CustomTemplate(id string) (*ResourceiOSConfigurationProfile_CustomTemplate, error) {
	endpoint := fmt.Sprintf("%s/%s?$expand=assignments", uriBetaGraphiOSDeviceManagementDeviceConfiguration, id)

	var responseiOSDeviceConfigurationProfile ResourceiOSConfigurationProfile_CustomTemplate
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseiOSDeviceConfigurationProfile)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "iOS configuration profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseiOSDeviceConfigurationProfile, nil
}

// GetiOSConfigurationProfileByDisplayName_CustomTemplate retrieves an iOS configuration profile with a custom template by its display name.
func (c *Client) GetiOSConfigurationProfileByDisplayName_CustomTemplate(displayName string) (*ResourceiOSConfigurationProfile_CustomTemplate, error) {
	// First, get the list of all iOS configuration profiles
	profiles, err := c.GetiOSConfigurationProfiles_CustomTemplates()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "iOS configuration profiles", displayName, err)
	}

	var profileID string
	// Search for a profile with the matching display name
	for _, profile := range profiles.Value {
		if profile.DisplayName == displayName {
			profileID = profile.ID
			break
		}
	}

	// Check if a profile with the given display name was found
	if profileID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "iOS configuration profile", displayName, err)
	}

	// Retrieve the specific profile by its ID
	return c.GetiOSConfigurationProfileByID_CustomTemplate(profileID)
}

// CreateiOSConfigurationProfile_CustomTemplate creates a iOS configuration profile using a custom template.
func (c *Client) CreateiOSConfigurationProfile_CustomTemplate(request *ResourceiOSConfigurationProfile_CustomTemplate) (*ResourceiOSConfigurationProfile_CustomTemplate, error) {
	request.ODataType = odataTypeiOSCustomConfigurationProfile
	endpoint := uriBetaGraphiOSDeviceManagementDeviceConfiguration

	var responseCreatedCustomConfigurationProfile ResourceiOSConfigurationProfile_CustomTemplate
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &responseCreatedCustomConfigurationProfile)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "iOS configuration profile", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseCreatedCustomConfigurationProfile, nil
}
