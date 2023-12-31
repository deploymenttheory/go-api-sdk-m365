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

const uriBetaDeviceManagementScripts = "/beta/deviceManagement/deviceManagementScripts"

/* Struct hierarchy using embedded anonymous structs for reference
type ResourceDeviceManagementScriptsList struct {
    ODataContext string `json:"@odata.context"`
    Value        []struct {
        OdataContext          string    `json:"@odata.context"`
        MicrosoftGraphTips    string    `json:"@microsoft.graph.tips"`
        EnforceSignatureCheck bool      `json:"enforceSignatureCheck"`
        RunAs32Bit            bool      `json:"runAs32Bit"`
        ID                    string    `json:"id"`
        DisplayName           string    `json:"displayName"`
        Description           string    `json:"description"`
        ScriptContent         string    `json:"scriptContent"`
        CreatedDateTime       time.Time `json:"createdDateTime"`
        LastModifiedDateTime  time.Time `json:"lastModifiedDateTime"`
        RunAsAccount          string    `json:"runAsAccount"`
        FileName              string    `json:"fileName"`
        RoleScopeTagIds       []string  `json:"roleScopeTagIds"`
        Assignments           []struct {
            ID     string `json:"id"`
            Target struct {
                OdataType                                  string `json:"@odata.type"`
                DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
                DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
                GroupId                                    string `json:"groupId"`
            } `json:"target"`
        } `json:"assignments"`
        RunSummary *struct {
            // Define fields for RunSummary
        } `json:"runSummary"`
    } `json:"value"`
}
*/

// ResourceDeviceManagementScriptsList represents a list of Device Management Scripts.
type ResourceDeviceManagementScriptsList struct {
	ODataContext string                           `json:"@odata.context"`
	Value        []ResourceDeviceManagementScript `json:"value"`
}

// ResourceDeviceManagementScript represents a Device Management Script resource
type ResourceDeviceManagementScript struct {
	OdataContext          string                                     `json:"@odata.context"`
	MicrosoftGraphTips    string                                     `json:"@microsoft.graph.tips"`
	EnforceSignatureCheck bool                                       `json:"enforceSignatureCheck"`
	RunAs32Bit            bool                                       `json:"runAs32Bit"`
	ID                    string                                     `json:"id"`
	DisplayName           string                                     `json:"displayName"`
	Description           string                                     `json:"description"`
	ScriptContent         string                                     `json:"scriptContent"`
	CreatedDateTime       time.Time                                  `json:"createdDateTime"`
	LastModifiedDateTime  time.Time                                  `json:"lastModifiedDateTime"`
	RunAsAccount          string                                     `json:"runAsAccount"`
	FileName              string                                     `json:"fileName"`
	RoleScopeTagIds       []string                                   `json:"roleScopeTagIds"`
	Assignments           []ResourceDeviceManagementScriptAssignment `json:"assignments"`
	RunSummary            *RunSummary                                `json:"runSummary"`
}

// Struct representing an assignment of a Device Management Script
type ResourceDeviceManagementScriptAssignment struct {
	ID     string `json:"id"`
	Target Target `json:"target"`
}

// Struct representing the target of a script assignment
type Target struct {
	OdataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	GroupId                                    string `json:"groupId"`
}

// Struct representing the run summary of a Device Management Script
type RunSummary struct {
	// Define the fields of RunSummary based on the API response
}

// Struct for a Device Management Script resource creation and update request
type ResourceDeviceManagementScriptRequest struct {
	OdataType             string   `json:"@odata.type"`
	DisplayName           string   `json:"displayName"`
	Description           string   `json:"description"`
	ScriptContent         string   `json:"scriptContent"`
	RunAsAccount          string   `json:"runAsAccount"`
	EnforceSignatureCheck bool     `json:"enforceSignatureCheck"`
	FileName              string   `json:"fileName"`
	RoleScopeTagIds       []string `json:"roleScopeTagIds"`
	RunAs32Bit            bool     `json:"runAs32Bit"`
}

// GetDeviceManagementScripts gets a list of all Intune Device Management Scripts
// with expanded information on assignments and run summary.
func (c *Client) GetDeviceManagementScripts() (*ResourceDeviceManagementScriptsList, error) {
	// Append query parameters to the endpoint URL
	endpoint := uriBetaDeviceManagementScripts + "?$expand=assignments,runSummary"

	var deviceManagementScripts ResourceDeviceManagementScriptsList
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
func (c *Client) GetDeviceManagementScriptByID(id string) (*ResourceDeviceManagementScript, error) {
	endpoint := fmt.Sprintf("%s/%s?$expand=assignments", uriBetaDeviceManagementScripts, id)

	var deviceManagementScript ResourceDeviceManagementScript
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
func (c *Client) GetDeviceManagementScriptByDisplayName(displayName string) (*ResourceDeviceManagementScript, error) {
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
func (c *Client) CreateDeviceManagementScript(request *ResourceDeviceManagementScriptRequest) (*ResourceDeviceManagementScript, error) {
	endpoint := uriBetaDeviceManagementScripts

	var createdScript ResourceDeviceManagementScript
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &createdScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdScript, nil
}

// UpdateDeviceManagementScriptByID updates an existing device management script by its ID.
// This endpoint supports the HTTP method PATCH for resource attribute updates.
func (c *Client) UpdateDeviceManagementScriptByID(id string, updateRequest *ResourceDeviceManagementScriptRequest) (*ResourceDeviceManagementScript, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementScripts, id)

	var updatedScript ResourceDeviceManagementScript
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, updateRequest, &updatedScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedUpdateByID, "device management script", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedScript, nil
}

// UpdateDeviceManagementScriptByDisplayName updates an existing device management script by its display name.
// Since there is no dedicated endpoint for this, it first retrieves the script by name to get its ID,
// then updates it using the UpdateDeviceManagementScriptByID function.
func (c *Client) UpdateDeviceManagementScriptByDisplayName(displayName string, updateRequest *ResourceDeviceManagementScriptRequest) (*ResourceDeviceManagementScript, error) {
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
