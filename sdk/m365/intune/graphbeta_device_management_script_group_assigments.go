// graphbeta_device_management_script_group_assigments.go
// Graph Beta Api - Intune: Windows Script group assignments
// Documentation: https://learn.microsoft.com/en-us/mem/intune/fundamentals/remediations
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesMenu/~/remediations
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-devicemanagementscriptgroupassignment?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const uriBetaDeviceManagementScriptsGroupAssignment = "/beta/deviceManagement/deviceManagementScripts"
const uriDeviceManagementScriptsGroupAssignment = "/deviceManagement/deviceManagementScripts"

// ResourceDeviceManagementScriptGroupAssignmentsList represents the response structure for group assignment requests.
type ResourceDeviceManagementScriptGroupAssignmentsList struct {
	Value []DeviceManagementScriptGroupAssignment `json:"value"`
}

// DeviceManagementScriptGroupAssignment represents a group assignment for a device management script.
type DeviceManagementScriptGroupAssignment struct {
	OdataType string                `json:"@odata.type"`
	ID        string                `json:"id"`
	Target    GroupAssignmentTarget `json:"target"`
}

// GroupAssignmentTarget represents the target for a group assignment.
type GroupAssignmentTarget struct {
	OdataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId,omitempty"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType,omitempty"`
	CollectionId                               string `json:"collectionId,omitempty"`
	GroupId                                    string `json:"groupId,omitempty"`
}

// RequestDeviceManagementScriptGroupAssignmentCreate represents the request structure for creating a group assignment.
type RequestDeviceManagementScriptGroupAssignmentCreate struct {
	OdataType     string `json:"@odata.type"`
	TargetGroupId string `json:"targetGroupId"`
}

// ResponseDeviceManagementScriptGroupAssignmentCreate represents the response structure for a group assignment creation.
type ResponseDeviceManagementScriptGroupAssignmentCreate struct {
	OdataType     string `json:"@odata.type"`
	ID            string `json:"id"`
	TargetGroupId string `json:"targetGroupId"`
}

// GetDeviceManagementScriptGroupAssignmentsByScriptID retrieves all group assignments for a device management script (Windows).
func (c *Client) GetDeviceManagementScriptGroupAssignmentsByScriptID(scriptId string) (*ResourceDeviceManagementScriptGroupAssignmentsList, error) {
	endpoint := fmt.Sprintf("%s/%s/assignments", uriBetaDeviceManagementScriptsGroupAssignment, scriptId)

	var groupAssignments ResourceDeviceManagementScriptGroupAssignmentsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &groupAssignments)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management script group assignments", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &groupAssignments, nil
}

// GetDeviceManagementScriptGroupAssignmentsByScriptName retrieves all group assignments for a device management script by its display name.
func (c *Client) GetDeviceManagementScriptGroupAssignmentsByScriptDisplayName(scriptName string) (*ResourceDeviceManagementScriptGroupAssignmentsList, error) {
	// Retrieve all scripts
	scripts, err := c.GetDeviceManagementScripts()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management scripts", err)
	}

	// Find the script with the matching name
	var scriptID string
	for _, script := range scripts.Value {
		if script.DisplayName == scriptName {
			scriptID = script.ID
			break
		}
	}

	// Check if the script with the specified name was found
	if scriptID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device management script", scriptName, "script not found")
	}

	// Retrieve group assignments for the found script ID
	return c.GetDeviceManagementScriptGroupAssignmentsByScriptID(scriptID)
}
