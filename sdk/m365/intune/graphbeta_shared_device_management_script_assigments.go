// graphbeta_shared_device_management_script_assigments.go
// Graph Beta Api - Assignment action for device_management_scripts, device_shell_scripts and device_health_scripts.
// Documentation:
// Intune location:
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-devicemanagementscriptassignment?view=graph-rest-beta
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-devicemanagementscriptgroupassignment?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
)

// AssignmentDeviceManagementScript represents the request of a script assignment
type AssignmentDeviceManagementScript struct {
	ResourceDeviceManagementScriptGroupAssignments []ResourceDeviceManagementScriptGroupAssignment `json:"deviceManagementScriptGroupAssignments,omitempty"`
	ResourceDeviceManagementScriptAssignments      []ResourceDeviceManagementScriptAssignment      `json:"deviceManagementScriptAssignments,omitempty"`
}

// ResourceDeviceManagementScriptGroupAssignment represents a group assignment of a device management script
type ResourceDeviceManagementScriptGroupAssignment struct {
	OdataType     string `json:"@odata.type,omitempty"`
	ID            string `json:"id,omitempty"`
	TargetGroupID string `json:"targetGroupId,omitempty"`
}

// ResourceDeviceManagementScriptAssignment represents an assignment of a device management script
type ResourceDeviceManagementScriptAssignment struct {
	OdataType string `json:"@odata.type,omitempty"`
	ID        string `json:"id,omitempty"`
	Target    struct {
		OdataType                                  string `json:"@odata.type,omitempty"`
		DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId,omitempty"`
		DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType,omitempty"`
		CollectionID                               string `json:"collectionId,omitempty"`
	} `json:"target,omitempty"`
}

// GetResourceGroupAssignmentsByResourceID retrieves all group assignments for a specified resource.
func (c *Client) GetResourceGroupAssignmentsByResourceID(resourceTypeURI, resourceID string) (*AssignmentDeviceManagementScript, error) {
	endpoint := fmt.Sprintf("%s/%s/assignments", resourceTypeURI, resourceID)

	var assignments AssignmentDeviceManagementScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &assignments)
	if err != nil {
		return nil, fmt.Errorf("failed to get group assignments for resource ID %s: %v", resourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &assignments, nil
}
