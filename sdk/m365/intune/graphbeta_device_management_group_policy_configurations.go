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
	OdataType                        string                       `json:"@odata.type"`
	ID                               string                       `json:"id"`
	DisplayName                      string                       `json:"displayName"`
	Description                      string                       `json:"description"`
	RoleScopeTagIds                  []string                     `json:"roleScopeTagIds"`
	PolicyConfigurationIngestionType string                       `json:"policyConfigurationIngestionType"`
	CreatedDateTime                  time.Time                    `json:"createdDateTime"`
	LastModifiedDateTime             time.Time                    `json:"lastModifiedDateTime"`
	DefinitionValues                 []GroupPolicyDefinitionValue `json:"definitionValues,omitempty"`
	Assignments                      []Assignment                 `json:"assignments,omitempty"`
}

// Struct for holding a list of Group Policy Definition Values
type ResponseGroupPolicyDefinitionValuesList struct {
	Value []GroupPolicyDefinitionValue `json:"value"`
}

// GroupPolicyDefinitionValue represents a single Group Policy Definition Value.
type GroupPolicyDefinitionValue struct {
	ID                   string                         `json:"id"`
	Enabled              bool                           `json:"enabled"`
	ConfigurationType    string                         `json:"configurationType"`
	CreatedDateTime      time.Time                      `json:"createdDateTime"`
	LastModifiedDateTime time.Time                      `json:"lastModifiedDateTime"`
	Definition           *GroupPolicyDefinition         `json:"definition,omitempty"`
	PresentationValues   []GroupPolicyPresentationValue `json:"presentationValues,omitempty"`
}

// GroupPolicyDefinition represents the definition of a Group Policy.
type GroupPolicyDefinition struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	// Add other relevant fields as required
}

type ResponsePresentationValuesList struct {
	Value []GroupPolicyPresentationValue `json:"value"`
}

type GroupPolicyPresentationValue struct {
	ID                   string                  `json:"id"`
	LastModifiedDateTime time.Time               `json:"lastModifiedDateTime"`
	CreatedDateTime      time.Time               `json:"createdDateTime"`
	Label                string                  `json:"label"`
	Description          string                  `json:"description"`
	ValueType            string                  `json:"valueType"`
	Value                DynamicValue            `json:"value"`
	Presentation         GroupPolicyPresentation `json:"presentation,omitempty"`
}

type GroupPolicyPresentation struct {
	Label       string             `json:"label"`
	ID          string             `json:"id"`
	Required    bool               `json:"required"`
	DefaultItem PresentationItem   `json:"defaultItem"`
	Items       []PresentationItem `json:"items"`
}

type PresentationItem struct {
	DisplayName string `json:"displayName"`
	Value       string `json:"value"`
}

type DynamicValue struct {
	// Use an interface to hold the actual value.
	Value interface{} `json:"Value"`
}

// Struct for holding a list of Assignments
type ResponseAssignmentsList struct {
	Value []Assignment `json:"value"`
}

// Assignment represents an assignment of a Group Policy Configuration
type Assignment struct {
	ID                   string           `json:"id"`
	LastModifiedDateTime time.Time        `json:"lastModifiedDateTime"`
	Target               AssignmentTarget `json:"target"`
	// Add other relevant fields as required
}

// AssignmentTarget represents the target of an Assignment (User, Group, etc.)
type AssignmentTarget struct {
	ID                                         string `json:"id"`
	Type                                       string `json:"@odata.type"` // e.g., "#microsoft.graph.groupTarget", etc.
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	CollectionId                               string `json:"collectionId"`
	// Add other relevant fields as required
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

// GetDeviceManagementGroupPolicyConfigurationByID retrieves a specific Group Policy Configuration by its ID with expanded details.
func (c *Client) GetDeviceManagementGroupPolicyConfigurationByID(policyConfigurationId string) (*ResourceDeviceManagementGroupPolicyConfiguration, error) {
	// Retrieve the base Group Policy Configuration
	baseEndpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementGroupPolicyConfigurations, policyConfigurationId)
	var baseConfig ResourceDeviceManagementGroupPolicyConfiguration
	_, err := c.HTTP.DoRequest("GET", baseEndpoint, nil, &baseConfig)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "group policy configuration", policyConfigurationId, err)
	}

	// Retrieve Definition Values and expand each definition
	defValuesEndpoint := fmt.Sprintf("%s/definitionValues?$expand=definition", baseEndpoint)
	var defValuesList ResponseGroupPolicyDefinitionValuesList
	_, err = c.HTTP.DoRequest("GET", defValuesEndpoint, nil, &defValuesList)
	if err != nil {
		return nil, fmt.Errorf("failed to get definition values: %v", err)
	}

	// For each Definition Value, retrieve and expand Presentation Values
	for i, defValue := range defValuesList.Value {
		presentationEndpoint := fmt.Sprintf("%s/definitionValues/%s/presentationValues?$expand=presentation", baseEndpoint, defValue.ID)
		var presentationList ResponsePresentationValuesList
		_, err = c.HTTP.DoRequest("GET", presentationEndpoint, nil, &presentationList)
		if err != nil {
			return nil, fmt.Errorf("failed to get presentation values: %v", err)
		}
		defValuesList.Value[i].PresentationValues = presentationList.Value
	}

	// Attach expanded Definition Values to the base configuration
	baseConfig.DefinitionValues = defValuesList.Value

	// Retrieve Assignments
	assignmentsEndpoint := fmt.Sprintf("%s/assignments", baseEndpoint)
	var assignmentsList ResponseAssignmentsList
	_, err = c.HTTP.DoRequest("GET", assignmentsEndpoint, nil, &assignmentsList)
	if err != nil {
		return nil, fmt.Errorf("failed to get assignments: %v", err)
	}

	// Attach Assignments to the base configuration
	baseConfig.Assignments = assignmentsList.Value

	return &baseConfig, nil
}
