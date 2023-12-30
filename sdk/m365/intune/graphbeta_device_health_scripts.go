// graphbeta_device_health_Scripts.go
// Graph Beta Api - Intune: Proactive Remediations
// Documentation: https://learn.microsoft.com/en-us/mem/intune/fundamentals/remediations
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/powershell
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-devicehealthscript?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriBetaProactiveRemediations                  = "/beta/deviceManagement/deviceHealthScripts"
	ODataTypeDeviceHealthScript                   = "#microsoft.graph.deviceHealthScript"
	ODataTypeDeviceHealthScriptStringParameter    = "microsoft.graph.deviceHealthScriptStringParameter"
	ODataTypeDeviceHealthScriptAssignment         = "#microsoft.graph.deviceHealthScriptAssignment"
	ODataTypeConfigurationManagerCollectionTarget = "microsoft.graph.configurationManagerCollectionAssignmentTarget"
	ODataTypeDeviceHealthScriptDailySchedule      = "microsoft.graph.deviceHealthScriptDailySchedule"
	ODataTypeGroupAssignmentTarget                = "microsoft.graph.groupAssignmentTarget"
)

/* Struct hierarchy using embedded anonymous structs for reference

// ResourceDeviceHealthScript represents an individual Proactive Remediation (Device Health Script) resource from Microsoft Graph API.
type ResourceDeviceHealthScript struct {
	OdataType                   string `json:"@odata.type"`
	ID                          string `json:"id"`
	Publisher                   string `json:"publisher"`
	Version                     string `json:"version"`
	DisplayName                 string `json:"displayName"`
	Description                 string `json:"description"`
	DetectionScriptContent      []byte `json:"detectionScriptContent"`
	RemediationScriptContent    []byte `json:"remediationScriptContent"`
	CreatedDateTime             string `json:"createdDateTime"`
	LastModifiedDateTime        string `json:"lastModifiedDateTime"`
	RunAsAccount                string `json:"runAsAccount"`
	EnforceSignatureCheck       bool   `json:"enforceSignatureCheck"`
	RunAs32Bit                  bool   `json:"runAs32Bit"`
	RoleScopeTagIds             []string `json:"roleScopeTagIds"`
	IsGlobalScript              bool     `json:"isGlobalScript"`
	HighestAvailableVersion     string   `json:"highestAvailableVersion"`
	DeviceHealthScriptType      string   `json:"deviceHealthScriptType"`
	DetectionScriptParameters   []struct {
		OdataType                        string `json:"@odata.type"`
		Name                             string `json:"name"`
		Description                      string `json:"description"`
		IsRequired                       bool   `json:"isRequired"`
		ApplyDefaultValueWhenNotAssigned bool   `json:"applyDefaultValueWhenNotAssigned"`
		DefaultValue                     string `json:"defaultValue"`
	} `json:"detectionScriptParameters"`
	RemediationScriptParameters []struct {
		OdataType                        string `json:"@odata.type"`
		Name                             string `json:"name"`
		Description                      string `json:"description"`
		IsRequired                       bool   `json:"isRequired"`
		ApplyDefaultValueWhenNotAssigned bool   `json:"applyDefaultValueWhenNotAssigned"`
		DefaultValue                     string `json:"defaultValue"`
	} `json:"remediationScriptParameters"`
	Assignments []struct {
		OdataType            string `json:"@odata.type"`
		ID                   string `json:"id"`
		Target               struct {
			OdataType                                  string `json:"@odata.type"`
			DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
			DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
			CollectionId                               string `json:"collectionId"`
		} `json:"target"`
		RunRemediationScript bool `json:"runRemediationScript"`
		RunSchedule          struct {
			OdataType string `json:"@odata.type"`
			Interval  int    `json:"interval"`
			UseUtc    bool   `json:"useUtc"`
			Time      string `json:"time"`
		} `json:"runSchedule"`
	} `json:"assignments"`
}
*/

// ResponseProactiveRemediationsList is used to parse the list response of Proactive Remediations from Microsoft Graph API.
type ResponseProactiveRemediationsList struct {
	ODataContext string                       `json:"@odata.context"`
	Value        []ResourceDeviceHealthScript `json:"value"`
}

// ResourceDeviceHealthScript represents an individual Proactive Remediation (Device Health Script) resource from Microsoft Graph API.
type ResourceDeviceHealthScript struct {
	OdataType                   string                         `json:"@odata.type"`
	ID                          string                         `json:"id"`
	Publisher                   string                         `json:"publisher"`
	Version                     string                         `json:"version"`
	DisplayName                 string                         `json:"displayName"`
	Description                 string                         `json:"description"`
	DetectionScriptContent      string                         `json:"detectionScriptContent"`
	RemediationScriptContent    string                         `json:"remediationScriptContent"`
	CreatedDateTime             string                         `json:"createdDateTime"`
	LastModifiedDateTime        string                         `json:"lastModifiedDateTime"`
	RunAsAccount                string                         `json:"runAsAccount"`
	EnforceSignatureCheck       bool                           `json:"enforceSignatureCheck"`
	RunAs32Bit                  bool                           `json:"runAs32Bit"`
	RoleScopeTagIds             []string                       `json:"roleScopeTagIds"`
	IsGlobalScript              bool                           `json:"isGlobalScript"`
	HighestAvailableVersion     string                         `json:"highestAvailableVersion"`
	DeviceHealthScriptType      string                         `json:"deviceHealthScriptType"`
	DetectionScriptParameters   []DeviceHealthScriptParameter  `json:"detectionScriptParameters"`
	RemediationScriptParameters []DeviceHealthScriptParameter  `json:"remediationScriptParameters"`
	Assignments                 []AssignmentDeviceHealthScript `json:"assignments"`
}

// DeviceHealthScriptParameter represents a parameter for a Device Health Script in Microsoft Graph API.
type DeviceHealthScriptParameter struct {
	OdataType                        string `json:"@odata.type"`
	Name                             string `json:"name"`
	Description                      string `json:"description"`
	IsRequired                       bool   `json:"isRequired"`
	ApplyDefaultValueWhenNotAssigned bool   `json:"applyDefaultValueWhenNotAssigned"`
	DefaultValue                     string `json:"defaultValue"`
}

// Assignment represents an assignment of a Proactive Remediation to a target.
type AssignmentDeviceHealthScript struct {
	OdataType            string                             `json:"@odata.type"`
	ID                   string                             `json:"id"`
	Target               AssignmentTargetDeviceHealthScript `json:"target"`
	RunRemediationScript bool                               `json:"runRemediationScript"`
	RunSchedule          RunScheduleDeviceHealthScript      `json:"runSchedule"`
}

// AssignmentTarget represents the target of an Assignment.
type AssignmentTargetDeviceHealthScript struct {
	OdataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	CollectionId                               string `json:"collectionId"`
}

// RunSchedule represents the schedule for running a device health script.
type RunScheduleDeviceHealthScript struct {
	OdataType string `json:"@odata.type"`
	Interval  int    `json:"interval"`
	UseUtc    bool   `json:"useUtc"`
	Time      string `json:"time"`
}

// GetProactiveRemediations retrieves a list of Proactive Remediations (Device Health Scripts) from Microsoft Graph API.
func (c *Client) GetProactiveRemediations() (*ResponseProactiveRemediationsList, error) {
	endpoint := uriBetaProactiveRemediations

	var responseProactiveRemediations ResponseProactiveRemediationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseProactiveRemediations)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "proactive remediations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProactiveRemediations, nil
}

// GetProactiveRemediationByID retrieves a specific Proactive Remediation by its ID along with its assignments.
func (c *Client) GetProactiveRemediationByID(remediationID string) (*ResourceDeviceHealthScript, error) {
	// Endpoint to get the proactive remediation details
	endpoint := fmt.Sprintf("%s/%s", uriBetaProactiveRemediations, remediationID)
	var remediation ResourceDeviceHealthScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &remediation)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "proactive remediation", remediationID, err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Endpoint to get the assignments for the proactive remediation
	assignmentsEndpoint := fmt.Sprintf("%s/%s/assignments", uriBetaProactiveRemediations, remediationID)
	var assignmentsResponse struct {
		Value []AssignmentDeviceHealthScript `json:"value"`
	}

	resp, err = c.HTTP.DoRequest("GET", assignmentsEndpoint, nil, &assignmentsResponse)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "proactive remediation assignments", remediationID, err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Append the assignments to the remediation object
	remediation.Assignments = assignmentsResponse.Value

	return &remediation, nil
}

// GetProactiveRemediationByDisplayName retrieves a specific Proactive Remediation by its name along with its assignments.
func (c *Client) GetProactiveRemediationByDisplayName(remediationName string) (*ResourceDeviceHealthScript, error) {
	// Retrieve all proactive remediations
	remediations, err := c.GetProactiveRemediations()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve proactive remediations: %v", err)
	}

	// Search for the remediation with the matching name
	var remediationID string
	for _, remediation := range remediations.Value {
		if remediation.DisplayName == remediationName {
			remediationID = remediation.ID
			break
		}
	}

	if remediationID == "" {
		return nil, fmt.Errorf("proactive remediation with name '%s' not found", remediationName)
	}

	// Get full details of the remediation using its ID
	return c.GetProactiveRemediationByID(remediationID)
}

// CreateProactiveRemediation creates a new Device Health Script in Microsoft Graph API.
func (c *Client) CreateProactiveRemediation(remediationData *ResourceDeviceHealthScript) (*ResourceDeviceHealthScript, error) {
	// Set the OdataType for the detection and remediation scripts
	remediationData.OdataType = ODataTypeDeviceHealthScript
	for i := range remediationData.DetectionScriptParameters {
		remediationData.DetectionScriptParameters[i].OdataType = ODataTypeDeviceHealthScriptStringParameter
	}
	for i := range remediationData.RemediationScriptParameters {
		remediationData.RemediationScriptParameters[i].OdataType = ODataTypeDeviceHealthScriptStringParameter
	}

	// Endpoint to create the device health script
	endpoint := uriBetaProactiveRemediations

	var createdRemediation ResourceDeviceHealthScript
	resp, err := c.HTTP.DoRequest("POST", endpoint, remediationData, &createdRemediation)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "proactive remediations", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdRemediation, nil
}

// CreateProactiveRemediationAssignment creates an assignment for a Device Health Script in Microsoft Graph API.
func (c *Client) CreateProactiveRemediationAssignment(scriptID string, assignmentData *AssignmentDeviceHealthScript) (*AssignmentDeviceHealthScript, error) {
	// Set the OdataType for the assignment and its properties
	assignmentData.OdataType = ODataTypeDeviceHealthScriptAssignment
	assignmentData.Target.OdataType = ODataTypeConfigurationManagerCollectionTarget
	assignmentData.RunSchedule.OdataType = ODataTypeDeviceHealthScriptDailySchedule

	assignmentEndpoint := fmt.Sprintf("%s/%s/assignments", uriBetaProactiveRemediations, scriptID)

	var createdAssignment AssignmentDeviceHealthScript
	resp, err := c.HTTP.DoRequest("POST", assignmentEndpoint, assignmentData, &createdAssignment)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "proactive remediations", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdAssignment, nil
}
