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
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriBetaProactiveRemediations                  = "/beta/deviceManagement/deviceHealthScripts"
	ODataTypeDeviceHealthScript                   = "#microsoft.graph.deviceHealthScript"
	ODataTypeDeviceHealthScriptStringParameter    = "microsoft.graph.deviceHealthScriptStringParameter"
	ODataTypeConfigurationManagerCollectionTarget = "microsoft.graph.configurationManagerCollectionAssignmentTarget"
	ODataTypeGroupAssignmentTarget                = "microsoft.graph.groupAssignmentTarget"
)

// ResponseProactiveRemediationsList represents a list of Proactive Remediation resources.
type ResponseProactiveRemediationsList struct {
	ODataContext       string                                 `json:"@odata.context"`
	ODataCount         int                                    `json:"@odata.count"`
	MicrosoftGraphTips string                                 `json:"@microsoft.graph.tips"`
	Value              []ResponseProactiveRemediationListItem `json:"value"`
}

// ResponseProactiveRemediationListItem represents a single health script entry from a list.
type ResponseProactiveRemediationListItem struct {
	ID                          string                                  `json:"id"`
	Publisher                   string                                  `json:"publisher"`
	Version                     string                                  `json:"version"`
	DisplayName                 string                                  `json:"displayName"`
	Description                 string                                  `json:"description"`
	DetectionScriptContent      string                                  `json:"detectionScriptContent"`
	RemediationScriptContent    string                                  `json:"remediationScriptContent"`
	CreatedDateTime             string                                  `json:"createdDateTime"`
	LastModifiedDateTime        string                                  `json:"lastModifiedDateTime"`
	RunAsAccount                string                                  `json:"runAsAccount"`
	EnforceSignatureCheck       bool                                    `json:"enforceSignatureCheck"`
	RunAs32Bit                  bool                                    `json:"runAs32Bit"`
	RoleScopeTagIds             []string                                `json:"roleScopeTagIds"`
	IsGlobalScript              bool                                    `json:"isGlobalScript"`
	HighestAvailableVersion     string                                  `json:"highestAvailableVersion"`
	DeviceHealthScriptType      string                                  `json:"deviceHealthScriptType"`
	DetectionScriptParameters   []interface{}                           `json:"detectionScriptParameters"`
	RemediationScriptParameters []interface{}                           `json:"remediationScriptParameters"`
	AssignmentsODataContext     string                                  `json:"assignments@odata.context"`
	Assignments                 []ResponseProactiveRemediatioAssignment `json:"assignments"`
}

// ResponseProactiveRemediatioAssignment represents an assignment for a health script.
type ResponseProactiveRemediatioAssignment struct {
	ID                   string                                      `json:"id"`
	RunRemediationScript bool                                        `json:"runRemediationScript"`
	Target               ResponseProactiveRemediationListTarget      `json:"target"`
	RunSchedule          ResponseProactiveRemediationListRunSchedule `json:"runSchedule"`
}

// ResponseProactiveRemediationListTarget represents the target of a script assignment.
type ResponseProactiveRemediationListTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	GroupID                                    string `json:"groupId"`
}

// ResponseProactiveRemediationListRunSchedule represents the schedule for running a script.
type ResponseProactiveRemediationListRunSchedule struct {
	ODataType string `json:"@odata.type"`
	Interval  int    `json:"interval"`
	UseUTC    bool   `json:"useUtc"`
	Time      string `json:"time"`
}

// ResponseProactiveRemediation represents the response for a single Proactive Remediation resource.
type ResponseProactiveRemediation struct {
	ODataContext                string                                   `json:"@odata.context"`
	MicrosoftGraphTips          string                                   `json:"@microsoft.graph.tips"`
	ID                          string                                   `json:"id"`
	Publisher                   string                                   `json:"publisher"`
	Version                     string                                   `json:"version"`
	DisplayName                 string                                   `json:"displayName"`
	Description                 string                                   `json:"description"`
	DetectionScriptContent      string                                   `json:"detectionScriptContent"`
	RemediationScriptContent    string                                   `json:"remediationScriptContent"`
	CreatedDateTime             time.Time                                `json:"createdDateTime"`
	LastModifiedDateTime        time.Time                                `json:"lastModifiedDateTime"`
	RunAsAccount                string                                   `json:"runAsAccount"`
	EnforceSignatureCheck       bool                                     `json:"enforceSignatureCheck"`
	RunAs32Bit                  bool                                     `json:"runAs32Bit"`
	RoleScopeTagIds             []string                                 `json:"roleScopeTagIds"`
	IsGlobalScript              bool                                     `json:"isGlobalScript"`
	HighestAvailableVersion     *string                                  `json:"highestAvailableVersion"`
	DeviceHealthScriptType      string                                   `json:"deviceHealthScriptType"`
	DetectionScriptParameters   []interface{}                            `json:"detectionScriptParameters"`
	RemediationScriptParameters []interface{}                            `json:"remediationScriptParameters"`
	AssignmentsODataContext     string                                   `json:"assignments@odata.context"`
	Assignments                 []ResponseProactiveRemediationAssignment `json:"assignments"`
}

// ResponseProactiveRemediationAssignment represents an assignment for a health script.
type ResponseProactiveRemediationAssignment struct {
	ID                   string                                            `json:"id"`
	RunRemediationScript bool                                              `json:"runRemediationScript"`
	Target               ResponseProactiveRemediationGroupAssignmentTarget `json:"target"`
	RunSchedule          ResponseProactiveRemediationRunSchedule           `json:"runSchedule"`
}

// ResponseProactiveRemediationGroupAssignmentTarget represents the target of a script assignment.
type ResponseProactiveRemediationGroupAssignmentTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	GroupId                                    string `json:"groupId"`
}

// ResponseProactiveRemediationRunSchedule represents the schedule for running a script.
type ResponseProactiveRemediationRunSchedule struct {
	ODataType string `json:"@odata.type"`
	Interval  int    `json:"interval"`
	UseUtc    bool   `json:"useUtc"`
	Time      string `json:"time"`
}

// ResourceProactiveRemediation represents the structure for a device health script in the Microsoft Graph API.
type ResourceProactiveRemediation struct {
	ODataType                   string                        `json:"@odata.type,omitempty"`
	Publisher                   string                        `json:"publisher,omitempty"`
	Version                     string                        `json:"version,omitempty"`
	DisplayName                 string                        `json:"displayName,omitempty"`
	Description                 string                        `json:"description,omitempty"`
	DetectionScriptContent      string                        `json:"detectionScriptContent,omitempty"`
	RemediationScriptContent    string                        `json:"remediationScriptContent,omitempty"`
	RunAsAccount                string                        `json:"runAsAccount,omitempty"`
	EnforceSignatureCheck       bool                          `json:"enforceSignatureCheck,omitempty"`
	RunAs32Bit                  bool                          `json:"runAs32Bit,omitempty"`
	RoleScopeTagIds             []string                      `json:"roleScopeTagIds,omitempty"`
	IsGlobalScript              bool                          `json:"isGlobalScript,omitempty"`
	HighestAvailableVersion     string                        `json:"highestAvailableVersion,omitempty"`
	DeviceHealthScriptType      string                        `json:"deviceHealthScriptType,omitempty"`
	DetectionScriptParameters   []DeviceHealthScriptParameter `json:"detectionScriptParameters,omitempty"`
	RemediationScriptParameters []DeviceHealthScriptParameter `json:"remediationScriptParameters,omitempty"`
}

// DeviceHealthScriptParameter represents a parameter for a device health script,
// which can be used in either detection or remediation scripts.
type DeviceHealthScriptParameter struct {
	ODataType                        string `json:"@odata.type,omitempty"`
	Name                             string `json:"name,omitempty"`
	Description                      string `json:"description,omitempty"`
	IsRequired                       bool   `json:"isRequired,omitempty"`
	ApplyDefaultValueWhenNotAssigned bool   `json:"applyDefaultValueWhenNotAssigned,omitempty"`
	DefaultValue                     string `json:"defaultValue,omitempty"`
}

// GetProactiveRemediations retrieves a list of Proactive Remediations (Device Health Scripts) from Microsoft Graph API.
func (c *Client) GetProactiveRemediations() (*ResponseProactiveRemediationsList, error) {
	endpoint := uriBetaProactiveRemediations + "?$expand=assignments"

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

// GetProactiveRemediationByID retrieves a Device Shell Script by its ID.
func (c *Client) GetProactiveRemediationByID(id string) (*ResponseProactiveRemediation, error) {
	endpoint := fmt.Sprintf("%s/%s?$expand=assignments", uriBetaProactiveRemediations, id)

	var proactiveRemediationScript ResponseProactiveRemediation
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &proactiveRemediationScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "proactive remediation", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &proactiveRemediationScript, nil
}

// GetProactiveRemediationByDisplayName retrieves a specific Proactive Remediation by its name along with its assignments.
func (c *Client) GetProactiveRemediationByDisplayName(displayName string) (*ResponseProactiveRemediation, error) {
	remediations, err := c.GetProactiveRemediations()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve proactive remediations: %v", err)
	}

	var remediationID string
	for _, remediation := range remediations.Value {
		if remediation.DisplayName == displayName {
			remediationID = remediation.ID
			break
		}
	}

	if remediationID == "" {
		return nil, fmt.Errorf("proactive remediation with name '%s' not found", displayName)
	}

	// Get full details of the remediation using its ID
	return c.GetProactiveRemediationByID(remediationID)
}

// CreateProactiveRemediation creates a new Device Health Script in Microsoft Graph API.
func (c *Client) CreateProactiveRemediation(request *ResourceProactiveRemediation) (*ResponseProactiveRemediation, error) {
	// Endpoint to create the device health script
	endpoint := uriBetaProactiveRemediations

	// Set the OdataType for the detection and remediation scripts
	request.ODataType = ODataTypeDeviceHealthScript
	for i := range request.DetectionScriptParameters {
		request.DetectionScriptParameters[i].ODataType = ODataTypeDeviceHealthScriptStringParameter
	}
	for i := range request.RemediationScriptParameters {
		request.RemediationScriptParameters[i].ODataType = ODataTypeDeviceHealthScriptStringParameter
	}

	var createdRemediation ResponseProactiveRemediation
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &createdRemediation)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "proactive remediations", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdRemediation, nil
}

// UpdateProactiveRemediationByID updates a Device Shell Script by its ID using the PATCH method.
func (c *Client) UpdateProactiveRemediationByID(scriptID string, request *ResourceProactiveRemediation) (*ResponseProactiveRemediation, error) {
	// Construct the endpoint URL
	endpoint := fmt.Sprintf("%s/%s", uriBetaProactiveRemediations, scriptID)

	// Set the OdataType for the detection and remediation scripts
	request.ODataType = ODataTypeDeviceHealthScript
	for i := range request.DetectionScriptParameters {
		request.DetectionScriptParameters[i].ODataType = ODataTypeDeviceHealthScriptStringParameter
	}
	for i := range request.RemediationScriptParameters {
		request.RemediationScriptParameters[i].ODataType = ODataTypeDeviceHealthScriptStringParameter
	}

	var updatedScript ResponseProactiveRemediation
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, &updatedScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedUpdateByID, "proactive remediation", scriptID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedScript, nil
}

// UpdateProactiveRemediationByDisplayName updates an existing Device Shell script by its display name.
// Since there is no dedicated endpoint for this, it first retrieves the script by name to get its ID,
// then updates it using the UpdateProactiveRemediationByID function.
func (c *Client) UpdateProactiveRemediationByDisplayName(displayName string, updateRequest *ResourceProactiveRemediation) (*ResponseProactiveRemediation, error) {
	// Retrieve the script by display name to get its ID
	scripts, err := c.GetProactiveRemediations()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device Shell scripts", err)
	}

	var scriptID string
	for _, script := range scripts.Value {
		if script.DisplayName == displayName {
			scriptID = script.ID
			break
		}
	}

	if scriptID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device Shell script", displayName, "script not found")
	}

	// Update the script by its ID using the provided updateRequest
	updatedScript, err := c.UpdateProactiveRemediationByID(scriptID, updateRequest)
	if err != nil {
		return nil, err
	}

	return updatedScript, nil
}

// DeleteDeviceShellScriptByID deletes an existing proactive remediation by its ID.
func (c *Client) DeleteProactiveRemediationByID(scriptID string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBetaProactiveRemediations, scriptID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedDeleteByID, "proactive remediation", scriptID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteProactiveRemediationByDisplayName deletes an existing device Shell script by its display name.
func (c *Client) DeleteProactiveRemediationByDisplayName(displayName string) error {
	script, err := c.GetProactiveRemediationByDisplayName(displayName)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGetByName, "proactive remediation", displayName, err)
	}

	return c.DeleteProactiveRemediationByID(script.ID)
}
