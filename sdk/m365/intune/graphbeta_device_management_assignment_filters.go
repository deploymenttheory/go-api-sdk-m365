// graphbeta_device_management_assignment_filters.go
// Graph Beta Api - Intune: Assignment Filters
// Documentation: https://learn.microsoft.com/en-us/mem/intune/fundamentals/filters
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesMenu/~/assignmentFilters
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-policyset-deviceandappmanagementassignmentfilter?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriBetaDeviceManagementAssignmentFilters   = "/beta/deviceManagement/assignmentFilters"
	odataTypeDeviceManagementAssignmentFilters = "#microsoft.graph.deviceAndAppManagementAssignmentFilter"
)

// ResponseAssignmentFiltersList represents a list of Assignment Filters.
type ResponseAssignmentFiltersList struct {
	ODataContext string                     `json:"@odata.context"`
	Value        []ResponseAssignmentFilter `json:"value"`
}

// ResponseAssignmentFilter represents an Assignment Filter resource.
type ResponseAssignmentFilter struct {
	ID                             string                            `json:"id"`
	CreatedDateTime                string                            `json:"createdDateTime"`
	LastModifiedDateTime           string                            `json:"lastModifiedDateTime"`
	DisplayName                    string                            `json:"displayName"`
	Description                    string                            `json:"description"`
	Platform                       string                            `json:"platform"`
	Rule                           string                            `json:"rule"`
	RoleScopeTags                  []string                          `json:"roleScopeTags"`
	AssignmentFilterManagementType string                            `json:"assignmentFilterManagementType"`
	Payloads                       []ResponseAssignmentFilterPayload `json:"payloads"`
}

// ResponseAssignmentFilterPayload represents the payload part of an Assignment Filter.
type ResponseAssignmentFilterPayload struct {
	ODataType            string `json:"@odata.type"`
	PayloadId            string `json:"payloadId"`
	PayloadType          string `json:"payloadType"`
	GroupId              string `json:"groupId"`
	AssignmentFilterType string `json:"assignmentFilterType"`
}

// ResourceDeviceManagementAssignmentFilter represents the request payload for creating a new Assignment Filter.
type ResourceDeviceManagementAssignmentFilter struct {
	ODataType                      string                                            `json:"@odata.type"`
	DisplayName                    string                                            `json:"displayName"`
	Description                    string                                            `json:"description"`
	Platform                       string                                            `json:"platform"`
	Rule                           string                                            `json:"rule"`
	RoleScopeTags                  []string                                          `json:"roleScopeTags"`
	Payloads                       []ResourceDeviceManagementAssignmentFilterPayload `json:"payloads"`
	AssignmentFilterManagementType string                                            `json:"assignmentFilterManagementType"`
}

// ResourceDeviceManagementAssignmentFilterPayload represents the payload part of an Assignment Filter.
type ResourceDeviceManagementAssignmentFilterPayload struct {
	ODataType            string `json:"@odata.type"`
	PayloadId            string `json:"payloadId"`
	PayloadType          string `json:"payloadType"`
	GroupId              string `json:"groupId"`
	AssignmentFilterType string `json:"assignmentFilterType"`
}

// GetDeviceManagementAssignmentFilters gets a list of all Intune Assignment Filters.
func (c *Client) GetDeviceManagementAssignmentFilters() (*ResponseAssignmentFiltersList, error) {
	endpoint := uriBetaDeviceManagementAssignmentFilters

	var assignmentFilters ResponseAssignmentFiltersList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &assignmentFilters)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "assignment filters", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &assignmentFilters, nil
}

// GetDeviceManagementAssignmentFilterByID retrieves a specific Assignment Filter by its ID.
func (c *Client) GetDeviceManagementAssignmentFilterByID(filterID string) (*ResponseAssignmentFilter, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementAssignmentFilters, filterID)

	var assignmentFilter ResponseAssignmentFilter
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &assignmentFilter)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "assignment filter", filterID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &assignmentFilter, nil
}

// GetDeviceManagementAssignmentFilterByDisplayName retrieves a specific intune Assignment Filter by its display name.
func (c *Client) GetDeviceManagementAssignmentFilterByDisplayName(displayName string) (*ResponseAssignmentFilter, error) {
	// Retrieve all assignment filters
	filtersList, err := c.GetDeviceManagementAssignmentFilters()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "assignment filters", err)
	}

	// Search for the filter with the matching display name
	var filterID string
	for _, filter := range filtersList.Value {
		if filter.DisplayName == displayName {
			filterID = filter.ID
			break
		}
	}

	if filterID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "assignment filter", displayName, "Filter not found")
	}
	// Retrieve the full details of the filter using its ID
	return c.GetDeviceManagementAssignmentFilterByID(filterID)
}

// CreateDeviceManagementAssignmentFilter creates a new Assignment Filter.
func (c *Client) CreateDeviceManagementAssignmentFilter(request *ResourceDeviceManagementAssignmentFilter) (*ResponseAssignmentFilter, error) {
	// Set graph metadata values
	request.ODataType = odataTypeDeviceManagementAssignmentFilters

	endpoint := uriBetaDeviceManagementAssignmentFilters

	var createdFilter ResponseAssignmentFilter
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &createdFilter)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "assignment filter", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdFilter, nil
}

// UpdateDeviceManagementAssignmentFilterByID updates a specific Assignment Filter by its ID.
func (c *Client) UpdateDeviceManagementAssignmentFilterByID(filterID string, request *ResourceDeviceManagementAssignmentFilter) (*ResponseAssignmentFilter, error) {
	// Set graph metadata values
	request.ODataType = odataTypeDeviceManagementAssignmentFilters

	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementAssignmentFilters, filterID)

	var updatedFilter ResponseAssignmentFilter
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, &updatedFilter)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedUpdateByID, "assignment filter", filterID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedFilter, nil
}

// UpdateDeviceManagementAssignmentFilterByDisplayName updates a specific Assignment Filter by its display name.
func (c *Client) UpdateDeviceManagementAssignmentFilterByDisplayName(displayName string, request *ResourceDeviceManagementAssignmentFilter) (*ResponseAssignmentFilter, error) {
	// Retrieve all assignment filters
	filtersList, err := c.GetDeviceManagementAssignmentFilters()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "assignment filters", err)
	}

	// Search for the filter with the matching display name
	var filterID string
	for _, filter := range filtersList.Value {
		if filter.DisplayName == displayName {
			filterID = filter.ID
			break
		}
	}

	if filterID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "assignment filter", displayName, "Filter not found")
	}

	// Update the filter by its ID using the provided request
	return c.UpdateDeviceManagementAssignmentFilterByID(filterID, request)
}

// DeleteDeviceManagementAssignmentFilterByID deletes a specific intune Assignment Filter by its ID.
func (c *Client) DeleteDeviceManagementAssignmentFilterByID(filterID string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceManagementAssignmentFilters, filterID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedDeleteByID, "assignment filter", filterID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDeviceManagementAssignmentFilterByDisplayName deletes a specific Assignment Filter by its display name.
func (c *Client) DeleteDeviceManagementAssignmentFilterByDisplayName(displayName string) error {
	// Retrieve all assignment filters
	filtersList, err := c.GetDeviceManagementAssignmentFilters()
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGet, "assignment filters", err)
	}

	// Search for the filter with the matching display name
	var filterID string
	for _, filter := range filtersList.Value {
		if filter.DisplayName == displayName {
			filterID = filter.ID
			break
		}
	}

	if filterID == "" {
		return fmt.Errorf(shared.ErrorMsgFailedGetByName, "assignment filter", displayName, "Filter not found")
	}

	// Delete the filter by its ID
	return c.DeleteDeviceManagementAssignmentFilterByID(filterID)
}
