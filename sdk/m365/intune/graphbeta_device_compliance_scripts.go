// graphbeta_device_compliance_scripts.go
// Graph Beta Api - Intune: Compliance
// Documentation: https://learn.microsoft.com/en-us/mem/intune/protect/compliance-custom-script
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesMenu/~/compliance
// API reference: https://learn.microsoft.com/en-us/graph/api/intune-devices-devicecompliancescript-list?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriBetaDeviceComplianceScripts                  = "/beta/deviceManagement/deviceComplianceScripts"
	ODataTypeDeviceComplianceScript                 = "#microsoft.graph.deviceComplianceScript"
	ODataTypeDeviceComplianceScriptStringParameter  = "microsoft.graph.deviceHealthScriptStringParameter"
	ODataTypeDeviceComplianceScriptAssignment       = "#microsoft.graph.deviceHealthScriptAssignment"
	ODataTypeDeviceComplianceScriptDailySchedule    = "microsoft.graph.deviceHealthScriptDailySchedule"
	ODataTypeDeviceComplianceScriptAssignmentTarget = "microsoft.graph.groupAssignmentTarget"
)

// ResponseDeviceComplianceScripts represents a response containing a list of Device Compliance Scripts.
type ResponseDeviceComplianceScriptsList struct {
	ODataContext       string                                    `json:"@odata.context"`
	ODataCount         int                                       `json:"@odata.count"`
	MicrosoftGraphTips string                                    `json:"@microsoft.graph.tips"`
	Value              []ResponseDeviceComplianceScriptsListItem `json:"value"`
}

// ResponseDeviceComplianceScriptsListItem represents a single Device Compliance Script item.
type ResponseDeviceComplianceScriptsListItem struct {
	ID                     string                                          `json:"id"`
	Publisher              string                                          `json:"publisher"`
	Version                string                                          `json:"version"`
	DisplayName            string                                          `json:"displayName"`
	Description            string                                          `json:"description"`
	DetectionScriptContent *string                                         `json:"detectionScriptContent"`
	CreatedDateTime        time.Time                                       `json:"createdDateTime"`
	LastModifiedDateTime   time.Time                                       `json:"lastModifiedDateTime"`
	RunAsAccount           string                                          `json:"runAsAccount"`
	EnforceSignatureCheck  bool                                            `json:"enforceSignatureCheck"`
	RunAs32Bit             bool                                            `json:"runAs32Bit"`
	RoleScopeTagIds        []string                                        `json:"roleScopeTagIds"`
	Assignments            []ResponseDeviceComplianceScriptsListAssignment `json:"assignments"`
}

// DeviceComplianceAssignment represents an assignment for a Device Compliance Script.
type ResponseDeviceComplianceScriptsListAssignment struct {
	ID                   string                                          `json:"id"`
	RunRemediationScript bool                                            `json:"runRemediationScript"`
	Target               ResponseDeviceComplianceScriptsListTarget       `json:"target"`
	RunSchedule          *ResponseDeviceComplianceScriptsListRunSchedule `json:"runSchedule,omitempty"` // Can be null, so pointer type is used
}

// ResponseDeviceComplianceScriptsListTarget represents the target of a compliance script assignment.
type ResponseDeviceComplianceScriptsListTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	GroupID                                    string `json:"groupId"`
}

// ResponseDeviceComplianceScriptsListRunSchedule represents the schedule for running a compliance script.
type ResponseDeviceComplianceScriptsListRunSchedule struct {
	ODataType string `json:"@odata.type"`
	Interval  int    `json:"interval"`
}

// ResponseDeviceComplianceScript represents the detailed information of a single device compliance script.
type ResponseDeviceComplianceScript struct {
	ODataContext            string                       `json:"@odata.context"`
	MicrosoftGraphTips      string                       `json:"@microsoft.graph.tips"`
	ID                      string                       `json:"id"`
	Publisher               string                       `json:"publisher"`
	Version                 string                       `json:"version"`
	DisplayName             string                       `json:"displayName"`
	Description             string                       `json:"description"`
	DetectionScriptContent  string                       `json:"detectionScriptContent"`
	CreatedDateTime         time.Time                    `json:"createdDateTime"`
	LastModifiedDateTime    time.Time                    `json:"lastModifiedDateTime"`
	RunAsAccount            string                       `json:"runAsAccount"`
	EnforceSignatureCheck   bool                         `json:"enforceSignatureCheck"`
	RunAs32Bit              bool                         `json:"runAs32Bit"`
	RoleScopeTagIds         []string                     `json:"roleScopeTagIds"`
	AssignmentsODataContext string                       `json:"assignments@odata.context"`
	Assignments             []DeviceComplianceAssignment `json:"assignments"`
}

// DeviceComplianceAssignment represents an assignment for a Device Compliance Script.
type DeviceComplianceAssignment struct {
	ID                   string                            `json:"id"`
	RunRemediationScript bool                              `json:"runRemediationScript"`
	Target               DeviceComplianceAssignmentTarget  `json:"target"`
	RunSchedule          DeviceComplianceScriptRunSchedule `json:"runSchedule"`
}

// DeviceComplianceAssignmentTarget represents the target of a compliance script assignment.
type DeviceComplianceAssignmentTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	GroupID                                    string `json:"groupId"`
}

// DeviceComplianceScriptRunSchedule represents the schedule for running a compliance script.
type DeviceComplianceScriptRunSchedule struct {
	ODataType string `json:"@odata.type"`
	Interval  int    `json:"interval"`
	// Add other fields as per the requirement
}

// ResourceDeviceComplianceScript represents the request structure for creating a device compliance script.
type ResourceDeviceComplianceScript struct {
	ODataType              string   `json:"@odata.type,omitempty"`
	Publisher              string   `json:"publisher,omitempty"`
	Version                string   `json:"version,omitempty"`
	DisplayName            string   `json:"displayName,omitempty"`
	Description            string   `json:"description,omitempty"`
	DetectionScriptContent string   `json:"detectionScriptContent,omitempty"`
	RunAsAccount           string   `json:"runAsAccount,omitempty"`
	EnforceSignatureCheck  bool     `json:"enforceSignatureCheck,omitempty"`
	RunAs32Bit             bool     `json:"runAs32Bit,omitempty"`
	RoleScopeTagIds        []string `json:"roleScopeTagIds,omitempty"`
}

// GetDeviceComplianceScripts retrieves a list of device compliance scripts from Microsoft Graph API.
func (c *Client) GetDeviceComplianceScripts() (*ResponseDeviceComplianceScriptsList, error) {
	endpoint := uriBetaDeviceComplianceScripts + "?$expand=assignments"

	var responseDeviceComplianceScripts ResponseDeviceComplianceScriptsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseDeviceComplianceScripts)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "proactive remediations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseDeviceComplianceScripts, nil
}

// GetDeviceComplianceScriptByID retrieves a Device Compliance Script by its ID.
func (c *Client) GetDeviceComplianceScriptByID(id string) (*ResponseDeviceComplianceScript, error) {
	endpoint := fmt.Sprintf("%s/%s?$expand=assignments", uriBetaDeviceComplianceScripts, id)

	var responseDeviceComplianceScript ResponseDeviceComplianceScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseDeviceComplianceScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "proactive remediation", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseDeviceComplianceScript, nil
}

// GetProactiveRemediationByDisplayName retrieves a specific Proactive Remediation by its name along with its assignments.
func (c *Client) GetDeviceComplianceScriptByDisplayName(displayName string) (*ResponseDeviceComplianceScript, error) {
	remediations, err := c.GetDeviceComplianceScripts()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "proactive remediations", err)
	}

	var deviceComplianceScriptID string
	for _, remediation := range remediations.Value {
		if remediation.DisplayName == displayName {
			deviceComplianceScriptID = remediation.ID
			break
		}
	}

	if deviceComplianceScriptID == "" {
		return nil, fmt.Errorf("proactive remediation with name '%s' not found", displayName)
	}

	// Get full details of the remediation using its ID
	return c.GetDeviceComplianceScriptByID(deviceComplianceScriptID)
}

// CreateDeviceComplianceScript creates a new device compliance script in Microsoft Graph API.
func (c *Client) CreateDeviceComplianceScript(request *ResourceDeviceComplianceScript) (*ResponseDeviceComplianceScript, error) {
	endpoint := uriBetaDeviceComplianceScripts

	// Set the ODataType for the request
	request.ODataType = ODataTypeDeviceComplianceScript

	var createdScript ResponseDeviceComplianceScript
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &createdScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device compliance script", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdScript, nil
}

// UpdateDeviceComplianceScriptByID updates a Device compliance Script by its ID using the PATCH method.
func (c *Client) UpdateDeviceComplianceScriptByID(scriptID string, request *ResourceDeviceComplianceScript) (*ResponseDeviceComplianceScript, error) {
	// Construct the endpoint URL
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceComplianceScripts, scriptID)

	// Set the request OData type
	request.ODataType = ODataTypeDeviceComplianceScript

	var updatedScript ResponseDeviceComplianceScript
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, &updatedScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedUpdateByID, "device compliance script", scriptID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedScript, nil
}

// UpdateDeviceComplianceScriptByDisplayName updates an existing Device Compliance script by its display name.
// Since there is no dedicated endpoint for this, it first retrieves the script by name to get its ID,
// then updates it using the UpdateDeviceComplianceScriptByID function.
func (c *Client) UpdateDeviceComplianceScriptByDisplayName(displayName string, updateRequest *ResourceDeviceComplianceScript) (*ResponseDeviceComplianceScript, error) {
	// Retrieve the script by display name to get its ID
	scripts, err := c.GetDeviceComplianceScripts()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device Compliance scripts", err)
	}

	var scriptID string
	for _, script := range scripts.Value {
		if script.DisplayName == displayName {
			scriptID = script.ID
			break
		}
	}

	if scriptID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device Compliance script", displayName, "script not found")
	}

	// Update the script by its ID using the provided updateRequest
	updatedScript, err := c.UpdateDeviceComplianceScriptByID(scriptID, updateRequest)
	if err != nil {
		return nil, err
	}

	return updatedScript, nil
}

// DeleteDeviceComplianceScriptByID deletes an existing device compliance script by its ID.
func (c *Client) DeleteDeviceComplianceScriptByID(scriptID string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceComplianceScripts, scriptID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedDeleteByID, "device compliance script", scriptID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDeviceComplianceScriptByDisplayName deletes an existing device Shell script by its display name.
func (c *Client) DeleteDeviceComplianceScriptByDisplayName(displayName string) error {
	script, err := c.GetDeviceComplianceScriptByDisplayName(displayName)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGetByName, "device compliance script", displayName, err)
	}

	return c.DeleteDeviceComplianceScriptByID(script.ID)
}
