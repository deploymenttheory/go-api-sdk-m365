// graphbeta_device_shell_scripts.go
// Graph Beta Api - Intune: macOS (shell) Scripts
// Documentation: https://learn.microsoft.com/en-us/mem/intune/apps/macos-shell-scripts
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesMacOsMenu/~/shellScripts
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-deviceshellscript?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.
// device shell scripts utilises 3 distinct sets of structs. A list object strucutre, a individual object structure and a resource object structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriBetaDeviceShellScripts                   = "/beta/deviceManagement/deviceShellScripts"
	uriBetaDeviceShellScriptAssignment          = "/beta/deviceManagement/deviceManagementScripts"
	odataTypeDeviceShellScript                  = "#microsoft.graph.deviceShellScript"
	odataTypeCreateDeviceShellScriptAssign      = "#microsoft.graph.deviceManagementScriptAssignment"
	odataTypeCreateDeviceShellScriptGroupAssign = "#microsoft.graph.deviceManagementScriptGroupAssignment"
)

// ResponseDeviceShellScriptsList represents a list of Device Shell Scripts.
type ResponseDeviceShellScriptsList struct {
	ODataContext string                              `json:"@odata.context"`
	Value        []ResponseDeviceShellScriptListItem `json:"value"`
}

// ResponseDeviceShellScript represents a Device Shell Script resource.
type ResponseDeviceShellScriptListItem struct {
	ExecutionFrequency          string                                    `json:"executionFrequency"`
	RetryCount                  int                                       `json:"retryCount"`
	BlockExecutionNotifications bool                                      `json:"blockExecutionNotifications"`
	ID                          string                                    `json:"id"`
	DisplayName                 string                                    `json:"displayName"`
	Description                 string                                    `json:"description"`
	ScriptContent               string                                    `json:"scriptContent"`
	CreatedDateTime             time.Time                                 `json:"createdDateTime"`
	LastModifiedDateTime        time.Time                                 `json:"lastModifiedDateTime"`
	RunAsAccount                string                                    `json:"runAsAccount"`
	FileName                    string                                    `json:"fileName"`
	RoleScopeTagIds             []string                                  `json:"roleScopeTagIds"`
	Assignments                 []ResponseDeviceShellScriptListAssignment `json:"assignments,omitempty"`
}

// ResponseDeviceShellScriptListAssignment represents an assignment of a Device Shell Script.
type ResponseDeviceShellScriptListAssignment struct {
	ID     string                              `json:"id"`
	Target ResponseDeviceShellScriptListTarget `json:"target"`
}

// ResponseDeviceShellScriptTarget represents the target of a script assignment.
type ResponseDeviceShellScriptListTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	CollectionId                               string `json:"collectionId"`
}

// ResponseDeviceShellScript represents a Device Shell Script by its ID.
type ResponseDeviceShellScript struct {
	ODataContext                string                                `json:"@odata.context"`
	Tips                        string                                `json:"@microsoft.graph.tips"`
	ExecutionFrequency          string                                `json:"executionFrequency"`
	RetryCount                  int                                   `json:"retryCount"`
	BlockExecutionNotifications bool                                  `json:"blockExecutionNotifications"`
	ID                          string                                `json:"id"`
	DisplayName                 string                                `json:"displayName"`
	Description                 string                                `json:"description"`
	ScriptContent               string                                `json:"scriptContent"`
	CreatedDateTime             string                                `json:"createdDateTime"`
	LastModifiedDateTime        string                                `json:"lastModifiedDateTime"`
	RunAsAccount                string                                `json:"runAsAccount"`
	FileName                    string                                `json:"fileName"`
	RoleScopeTagIds             []string                              `json:"roleScopeTagIds"`
	AssignmentsContext          string                                `json:"assignments@odata.context"`
	Assignments                 []ResponseDeviceShellScriptAssignment `json:"assignments"`
}

// ResponseDeviceShellScriptAssignment represents an assignment of a Device Shell Script by its ID.
type ResponseDeviceShellScriptAssignment struct {
	ID     string                          `json:"id"`
	Target ResponseDeviceShellScriptTarget `json:"target"`
}

// ResponseDeviceShellScriptTarget represents the target of a script assignment.
type ResponseDeviceShellScriptTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	GroupId                                    string `json:"groupId"`
}

// ResourceDeviceShellScript represents the request payload for creating and updating a new Device Shell Script.
type ResourceDeviceShellScript struct {
	ODataType                   string   `json:"@odata.type,omitempty"`
	ExecutionFrequency          string   `json:"executionFrequency,omitempty"`
	RetryCount                  int      `json:"retryCount,omitempty"`
	BlockExecutionNotifications bool     `json:"blockExecutionNotifications,omitempty"`
	DisplayName                 string   `json:"displayName,omitempty"`
	Description                 string   `json:"description,omitempty"`
	ScriptContent               string   `json:"scriptContent,omitempty"`
	RunAsAccount                string   `json:"runAsAccount,omitempty"`
	FileName                    string   `json:"fileName,omitempty"`
	RoleScopeTagIds             []string `json:"roleScopeTagIds,omitempty"`
}

// GetDeviceShellScripts gets a list of all Intune Device Shell Scripts
// with expanded information on assignments.
func (c *Client) GetDeviceShellScripts() (*ResponseDeviceShellScriptsList, error) {
	// Append query parameters to the endpoint URL
	endpoint := uriBetaDeviceShellScripts + "?$expand=assignments"

	var deviceShellScripts ResponseDeviceShellScriptsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceShellScripts)

	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device shell scripts", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceShellScripts, nil
}

// GetDeviceShellScriptByID retrieves a Device Shell Script by its ID.
func (c *Client) GetDeviceShellScriptByID(id string) (*ResponseDeviceShellScript, error) {
	endpoint := fmt.Sprintf("%s/%s?$expand=assignments", uriBetaDeviceShellScripts, id)

	var deviceShellScript ResponseDeviceShellScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceShellScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device shell script", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceShellScript, nil
}

// GetDeviceShellScriptByDisplayName retrieves a device shell script by its display name.
func (c *Client) GetDeviceShellScriptByDisplayName(displayName string) (*ResponseDeviceShellScript, error) {
	scripts, err := c.GetDeviceShellScripts()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device shell scripts", err)
	}

	var scriptID string
	for _, script := range scripts.Value {
		if script.DisplayName == displayName {
			scriptID = script.ID
			break
		}
	}

	if scriptID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device shell script", displayName, err)
	}

	return c.GetDeviceShellScriptByID(scriptID)
}

// CreateDeviceShellScript creates a new device management script.
func (c *Client) CreateDeviceShellScript(request *ResourceDeviceShellScript) (*ResponseDeviceShellScript, error) {
	request.ODataType = odataTypeDeviceShellScript
	endpoint := uriBetaDeviceShellScripts

	var createdScript ResponseDeviceShellScript
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &createdScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdScript, nil
}

// CreateDeviceShellScriptAssignment creates a new device management script assignment.
func (c *Client) CreateDeviceShellScriptAssignment(scriptID string, assignment *AssignmentDeviceManagementScript) (*ResourceDeviceManagementScriptGroupAssignment, error) {
	// Set graph metadata values
	for i := range assignment.ResourceDeviceManagementScriptAssignments {
		assignment.ResourceDeviceManagementScriptAssignments[i].OdataType = odataTypeCreateDeviceShellScriptAssign
	}

	for i := range assignment.ResourceDeviceManagementScriptGroupAssignments {
		assignment.ResourceDeviceManagementScriptGroupAssignments[i].OdataType = odataTypeCreateDeviceShellScriptGroupAssign
	}

	// Construct endpoint
	endpoint := fmt.Sprintf("%s/%s/assign", uriBetaDeviceShellScripts, scriptID)

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

// CreateDeviceShellScriptWithAssignment creates a new device management script and assigns it.
func (c *Client) CreateDeviceShellScriptWithAssignment(request *ResourceDeviceShellScript, assignment *AssignmentDeviceManagementScript) (*ResponseDeviceShellScript, error) {
	// Create the device management script
	createdScript, err := c.CreateDeviceShellScript(request)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script", err)
	}

	// Assign the script
	_, err = c.CreateDeviceShellScriptAssignment(createdScript.ID, assignment)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script assignment", err)
	}

	return createdScript, nil
}

// UpdateDeviceShellScriptByID updates a Device Shell Script by its ID using the PATCH method.
func (c *Client) UpdateDeviceShellScriptByID(scriptID string, request *ResourceDeviceShellScript) (*ResponseDeviceShellScript, error) {
	// Construct the endpoint URL
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceShellScripts, scriptID)

	// Set the request OData type
	request.ODataType = odataTypeDeviceShellScript

	var updatedScript ResponseDeviceShellScript
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, &updatedScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedUpdateByID, "device shell script", scriptID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedScript, nil
}

// UpdateDeviceShellScriptByDisplayName updates an existing Device Shell script by its display name.
// Since there is no dedicated endpoint for this, it first retrieves the script by name to get its ID,
// then updates it using the UpdateDeviceShellScriptByID function.
func (c *Client) UpdateDeviceShellScriptByDisplayName(displayName string, updateRequest *ResourceDeviceShellScript) (*ResponseDeviceShellScript, error) {
	// Retrieve the script by display name to get its ID
	scripts, err := c.GetDeviceShellScripts()
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
	updatedScript, err := c.UpdateDeviceShellScriptByID(scriptID, updateRequest)
	if err != nil {
		return nil, err
	}

	return updatedScript, nil
}

// DeleteDeviceShellScriptByID deletes an existing device shell script by its ID.
func (c *Client) DeleteDeviceShellScriptByID(scriptID string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceShellScripts, scriptID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedDeleteByID, "device shell script", scriptID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDeviceShellScriptByDisplayName deletes an existing device Shell script by its display name.
func (c *Client) DeleteDeviceShellScriptByDisplayName(displayName string) error {
	script, err := c.GetDeviceShellScriptByDisplayName(displayName)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGetByName, "device shell script", displayName, err)
	}

	return c.DeleteDeviceShellScriptByID(script.ID)
}
