// graphbeta_shared_device_management_script_group_assigments.go
// Graph Beta Api - Various
// Documentation:
// Intune location:
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-devicemanagementscriptgroupassignment?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"time"
)

// ResourceAssignmentsList represents the response structure for group assignment requests.
type ResourceAssignmentsList struct {
	Value []ResourceAssignment `json:"value"`
}

// ResourceAssignment represents a general group assignment structure.
type ResourceAssignment struct {
	OdataType string                   `json:"@odata.type"`
	ID        string                   `json:"id"`
	Target    ResourceAssignmentTarget `json:"target"`
}

// ResourceAssignmentTarget represents the target for a group assignment.
type ResourceAssignmentTarget struct {
	OdataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId,omitempty"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType,omitempty"`
	CollectionId                               string `json:"collectionId,omitempty"`
	GroupId                                    string `json:"groupId,omitempty"`
}

// RunSchedule represents the schedule details for running a script.
type RunSchedule struct {
	OdataType string    `json:"@odata.type"`
	Interval  int       `json:"interval"`
	UseUtc    bool      `json:"useUtc"`
	Time      time.Time `json:"time"`
}

// PostAssignmentRequest wraps the assignment requests for various script types.
type PostAssignmentRequest struct {
	Assignments []AssignmentRequest `json:"assignments"`
}

// AssignmentRequest represents the request structure for creating an assignment.
type AssignmentRequest struct {
	OdataType string       `json:"@odata.type"`
	ID        string       `json:"id"`
	Target    Target       `json:"target"`
	Schedule  *RunSchedule `json:"runSchedule,omitempty"`
}

// GetResourceGroupAssignmentsByResourceID retrieves all group assignments for a specified resource.
func (c *Client) GetResourceGroupAssignmentsByResourceID(resourceTypeURI, resourceID string) (*ResourceAssignmentsList, error) {
	endpoint := fmt.Sprintf("%s/%s/assignments", resourceTypeURI, resourceID)

	var assignments ResourceAssignmentsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &assignments)
	if err != nil {
		return nil, fmt.Errorf("Failed to get group assignments for resource ID %s: %v", resourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &assignments, nil
}
