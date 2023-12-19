// graphbeta_device_management_scripts.go
// Graph Beta Api - Intune: Proactive Remediation
// Documentation: https://learn.microsoft.com/en-us/mem/intune/fundamentals/remediations
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesMenu/~/remediations
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-shared-devicemanagementscript?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const urideviceManagementScripts = "/beta/deviceManagement/deviceManagementScripts"

// Struct for a list of Device Management Scripts
type ResourceDeviceManagementScriptsList struct {
	OdataContext       string `json:"@odata.context"`
	MicrosoftGraphTips string `json:"@microsoft.graph.tips"`
	Value              []struct {
		EnforceSignatureCheck bool        `json:"enforceSignatureCheck"`
		RunAs32Bit            bool        `json:"runAs32Bit"`
		ID                    string      `json:"id"`
		DisplayName           string      `json:"displayName"`
		Description           string      `json:"description"`
		ScriptContent         interface{} `json:"scriptContent"`
		CreatedDateTime       time.Time   `json:"createdDateTime"`
		LastModifiedDateTime  time.Time   `json:"lastModifiedDateTime"`
		RunAsAccount          string      `json:"runAsAccount"`
		FileName              string      `json:"fileName"`
		RoleScopeTagIds       []string    `json:"roleScopeTagIds"`
	} `json:"value"`
}

// Struct for a Device Management Script resource
type ResourceDeviceManagementScript struct {
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

// ResourceDeviceManagementScriptAssign represents the structure for assigning a device management script.
type ResourceDeviceManagementScriptAssign struct {
	DeviceManagementScriptGroupAssignments []DeviceManagementScriptGroupAssignment `json:"deviceManagementScriptGroupAssignments"`
	DeviceManagementScriptAssignments      []DeviceManagementScriptAssignment      `json:"deviceManagementScriptAssignments"`
}

// DeviceManagementScriptGroupAssignment represents a group assignment for a device management script.
type DeviceManagementScriptGroupAssignment struct {
	OdataType     string `json:"@odata.type"`
	ID            string `json:"id"`
	TargetGroupID string `json:"targetGroupId"`
}

// DeviceManagementScriptAssignment represents a script assignment.
type DeviceManagementScriptAssignment struct {
	OdataType string           `json:"@odata.type"`
	ID        string           `json:"id"`
	Target    AssignmentTarget `json:"target"`
}

// AssignmentTarget represents the target for an assignment.
type AssignmentTarget struct {
	OdataType string `json:"@odata.type"`
}

// GetDeviceManagementScripts gets a list of all Intune Device Management Scripts
func (c *Client) GetDeviceManagementScripts() (*ResourceDeviceManagementScriptsList, error) {
	endpoint := urideviceManagementScripts

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
	endpoint := fmt.Sprintf("%s/%s", urideviceManagementScripts, id)

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

// GetDeviceManagementScriptByName retrieves a device management script by its display name.
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
	endpoint := urideviceManagementScripts

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
	endpoint := fmt.Sprintf("%s/%s", urideviceManagementScripts, id)

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
	endpoint := fmt.Sprintf("%s/%s", urideviceManagementScripts, id)

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

// AssignDeviceManagementScriptByGroupID assigns a device management script to a group.
func (c *Client) AssignDeviceManagementScriptByGroupID(scriptID, groupID string) error {
	endpoint := fmt.Sprintf("%s/%s/assign", urideviceManagementScripts, scriptID)

	// Construct the request body
	assignment := struct {
		DeviceManagementScriptGroupAssignments []struct {
			OdataType     string `json:"@odata.type"`
			ID            string `json:"id"`
			TargetGroupID string `json:"targetGroupId"`
		} `json:"deviceManagementScriptGroupAssignments"`
	}{
		DeviceManagementScriptGroupAssignments: []struct {
			OdataType     string `json:"@odata.type"`
			ID            string `json:"id"`
			TargetGroupID string `json:"targetGroupId"`
		}{
			{
				OdataType:     "#microsoft.graph.deviceManagementScriptGroupAssignment",
				ID:            "assignment-" + groupID,
				TargetGroupID: groupID,
			},
		},
	}

	requestData, err := json.Marshal(assignment)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedJsonMarshal, "device management script assignment", err)
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, bytes.NewBuffer(requestData), nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script assignment", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
