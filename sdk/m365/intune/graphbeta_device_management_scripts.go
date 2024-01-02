// graphbeta_device_management_scripts.go
// Graph Beta Api - Intune: Windows (powershell) Scripts
// Documentation: https://learn.microsoft.com/en-us/mem/intune/apps/intune-management-extension
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/powershell
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-shared-devicemanagementscript?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriBetaDeviceManagementScripts                   = "/beta/deviceManagement/deviceManagementScripts"
	uriBetaDeviceManagementScriptAssignment          = "/beta/deviceManagement/deviceManagementScripts"
	odataTypeDeviceManagementScript                  = "#microsoft.graph.deviceManagementScript"
	odataTypeCreateDeviceManagementScriptAssign      = "#microsoft.graph.deviceManagementScriptAssignment"
	odataTypeCreateDeviceManagementScriptGroupAssign = "#microsoft.graph.deviceManagementScriptGroupAssignment"
)

// ResponseDeviceManagementScriptsList represents a list of Device Management Scripts.
type ResponseDeviceManagementScriptsList struct {
	ODataContext string                                   `json:"@odata.context"`
	Value        []ResponseDeviceManagementScriptListItem `json:"value"`
}

// ResponseDeviceManagementScriptListItem represents a Device Management Script resource.
type ResponseDeviceManagementScriptListItem struct {
	ExecutionFrequency          string                                         `json:"executionFrequency"`
	RetryCount                  int                                            `json:"retryCount"`
	BlockExecutionNotifications bool                                           `json:"blockExecutionNotifications"`
	ID                          string                                         `json:"id"`
	DisplayName                 string                                         `json:"displayName"`
	Description                 string                                         `json:"description"`
	ScriptContent               string                                         `json:"scriptContent"`
	CreatedDateTime             time.Time                                      `json:"createdDateTime"`
	LastModifiedDateTime        time.Time                                      `json:"lastModifiedDateTime"`
	RunAsAccount                string                                         `json:"runAsAccount"`
	FileName                    string                                         `json:"fileName"`
	RoleScopeTagIds             []string                                       `json:"roleScopeTagIds"`
	Assignments                 []ResponseDeviceManagementScriptListAssignment `json:"assignments,omitempty"`
}

// ResponseDeviceManagementScriptListAssignment represents an assignment of a Device Management Script.
type ResponseDeviceManagementScriptListAssignment struct {
	ID     string                                   `json:"id"`
	Target ResponseDeviceManagementScriptListTarget `json:"target"`
}

// ResponseDeviceManagementScriptListTarget represents the target of a script assignment.
type ResponseDeviceManagementScriptListTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	CollectionId                               string `json:"collectionId"`
}

// ResponseDeviceManagementScript represents a Device Management Script resource.
type ResponseDeviceManagementScript struct {
	ODataContext          string                                     `json:"@odata.context"`
	Tips                  string                                     `json:"@microsoft.graph.tips"`
	EnforceSignatureCheck bool                                       `json:"enforceSignatureCheck"`
	RunAs32Bit            bool                                       `json:"runAs32Bit"`
	ID                    string                                     `json:"id"`
	DisplayName           string                                     `json:"displayName"`
	Description           string                                     `json:"description"`
	ScriptContent         string                                     `json:"scriptContent"`
	CreatedDateTime       string                                     `json:"createdDateTime"`
	LastModifiedDateTime  string                                     `json:"lastModifiedDateTime"`
	RunAsAccount          string                                     `json:"runAsAccount"`
	FileName              string                                     `json:"fileName"`
	RoleScopeTagIds       []string                                   `json:"roleScopeTagIds"`
	AssignmentsContext    string                                     `json:"assignments@odata.context"`
	Assignments           []ResponseDeviceManagementScriptAssignment `json:"assignments"`
}

// ResponseDeviceManagementScriptAssignment represents an assignment of a Device Management Script.
type ResponseDeviceManagementScriptAssignment struct {
	ID     string                               `json:"id"`
	Target ResponseDeviceManagementScriptTarget `json:"target"`
}

// ResponseDeviceManagementScriptTarget represents the target of a script assignment.
type ResponseDeviceManagementScriptTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	GroupId                                    string `json:"groupId"`
}

// ResourceDeviceManagementScript represents the request payload for creating and updating a new Device Management Script.
type ResourceDeviceManagementScript struct {
	ODataType             string   `json:"@odata.type,omitempty"`
	DisplayName           string   `json:"displayName,omitempty"`
	Description           string   `json:"description,omitempty"`
	ScriptContent         string   `json:"scriptContent,omitempty"`
	RunAsAccount          string   `json:"runAsAccount,omitempty"`
	EnforceSignatureCheck bool     `json:"enforceSignatureCheck,omitempty"`
	FileName              string   `json:"fileName,omitempty"`
	RoleScopeTagIds       []string `json:"roleScopeTagIds,omitempty"`
	RunAs32Bit            bool     `json:"runAs32Bit,omitempty"`
}

// GetDeviceManagementScripts gets a list of all Intune Device Management Scripts
// with expanded information on assignments.
func (c *Client) GetDeviceManagementScripts() (*ResponseDeviceManagementScriptsList, error) {
	endpoint := uriBetaDeviceManagementScripts + "?$expand=assignments"

	var deviceManagementScripts ResponseDeviceManagementScriptsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceManagementScripts)

	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management scripts", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceManagementScripts, nil
}

// GetDeviceManagementScriptByID retrieves a Device Management Script by its ID.
func (c *Client) GetDeviceManagementScriptByID(id string) (*ResponseDeviceManagementScript, error) {
	endpoint := fmt.Sprintf("%s/%s?$expand=assignments", uriBetaDeviceManagementScripts, id)

	var deviceManagementScript ResponseDeviceManagementScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceManagementScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device management script", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceManagementScript, nil
}

// GetDeviceManagementScriptByDisplayName retrieves a device management script by its display name.
func (c *Client) GetDeviceManagementScriptByDisplayName(displayName string) (*ResponseDeviceManagementScript, error) {
	scripts, err := c.GetDeviceManagementScripts()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management scripts", err)
	}

	var scriptID string
	for _, script := range scripts.Value {
		if script.DisplayName == displayName {
			scriptID = script.ID
			break
		}
	}

	if scriptID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device management script", displayName, err)
	}

	return c.GetDeviceManagementScriptByID(scriptID)
}

// CreateDeviceManagementScript creates a new device management script.
func (c *Client) CreateDeviceManagementScript(request *ResourceDeviceManagementScript) (*ResponseDeviceManagementScript, error) {
	request.ODataType = odataTypeDeviceManagementScript
	endpoint := uriBetaDeviceManagementScripts

	var createdScript ResponseDeviceManagementScript
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &createdScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdScript, nil
}

// CreateDeviceManagementScriptAssignment creates a new device management script assignment.
func (c *Client) CreateDeviceManagementScriptAssignment(scriptID string, assignment *AssignmentDeviceManagementScript) (*ResourceDeviceManagementScriptGroupAssignment, error) {
	// Set graph metadata values
	for i := range assignment.ResourceDeviceManagementScriptAssignments {
		assignment.ResourceDeviceManagementScriptAssignments[i].OdataType = odataTypeCreateDeviceManagementScriptAssign
	}

	for i := range assignment.ResourceDeviceManagementScriptGroupAssignments {
		assignment.ResourceDeviceManagementScriptGroupAssignments[i].OdataType = odataTypeCreateDeviceManagementScriptGroupAssign
	}

	// Construct endpoint
	endpoint := fmt.Sprintf("%s/%s/assign", uriBetaDeviceManagementScripts, scriptID)

	var createdAssignment ResourceDeviceManagementScriptGroupAssignment
	resp, err := c.HTTP.DoRequest("POST", endpoint, assignment, &createdAssignment)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script assignment", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdAssignment, nil
}

// CreateDeviceManagementScriptWithAssignment creates a new device management script and assigns it.
func (c *Client) CreateDeviceManagementScriptWithAssignment(request *ResourceDeviceManagementScript, assignment *AssignmentDeviceManagementScript) (*ResponseDeviceManagementScript, error) {
	// Create the device management script
	createdScript, err := c.CreateDeviceManagementScript(request)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script", err)
	}

	// Assign the script
	_, err = c.CreateDeviceManagementScriptAssignment(createdScript.ID, assignment)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script assignment", err)
	}

	return createdScript, nil
}

// UpdateDeviceManagementScriptByID updates a Device Management Script by its ID using the PATCH method.
func (c *Client) UpdateDeviceManagementScriptByID(scriptID string, request *ResourceDeviceManagementScript) (*ResponseDeviceManagementScript, error) {
	// Construct the endpoint URL
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementScripts, scriptID)

	// Set the request OData type
	request.ODataType = odataTypeDeviceManagementScript

	var updatedScript ResponseDeviceManagementScript
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, &updatedScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedUpdateByID, "device shell script", scriptID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedScript, nil
}

// UpdateDeviceManagementScriptByDisplayName updates an existing device management script by its display name.
// Since there is no dedicated endpoint for this, it first retrieves the script by name to get its ID,
// then updates it using the UpdateDeviceManagementScriptByID function.
func (c *Client) UpdateDeviceManagementScriptByDisplayName(displayName string, updateRequest *ResourceDeviceManagementScript) (*ResponseDeviceManagementScript, error) {
	scripts, err := c.GetDeviceManagementScripts()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management scripts", err)
	}

	var scriptID string
	for _, script := range scripts.Value {
		if script.DisplayName == displayName {
			scriptID = script.ID
			break
		}
	}

	if scriptID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device management script", displayName, "script not found")
	}

	// Update the script by its ID
	return c.UpdateDeviceManagementScriptByID(scriptID, updateRequest)
}

// DeleteDeviceManagementScriptByID deletes an existing device management script by its ID.
func (c *Client) DeleteDeviceManagementScriptByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementScripts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedDeleteByID, "device management script", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDeviceManagementScriptByDisplayName deletes an existing device management script by its display name.
func (c *Client) DeleteDeviceManagementScriptByDisplayName(displayName string) error {
	script, err := c.GetDeviceManagementScriptByDisplayName(displayName)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGetByName, "device management script", displayName, err)
	}

	return c.DeleteDeviceManagementScriptByID(script.ID)
}
