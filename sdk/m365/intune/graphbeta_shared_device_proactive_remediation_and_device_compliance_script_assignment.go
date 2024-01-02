// graphbeta_shared_device_health_and_device_compliance_script_assignment.go
// Graph Beta Api - Assignment action for device_health_scripts and device_health_scripts.
// Documentation:
// Intune location:
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-devices-devicehealthscriptassignment?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	ODataTypeDeviceHealthScriptAssignment                   = "#microsoft.graph.deviceHealthScriptAssignment"
	ODataTypeConfigurationManagerCollectionAssignmentTarget = "microsoft.graph.configurationManagerCollectionAssignmentTarget"
	ODataTypeDeviceHealthScriptDailySchedule                = "microsoft.graph.deviceHealthScriptDailySchedule"
)

// ResponseDeviceHealthScriptAssignmentList represents a list of device health script assignments.
type ResponseDeviceHealthScriptAssignmentList struct {
	Value []DeviceHealthScriptAssignmentItem `json:"value"`
}

// DeviceHealthScriptAssignment represents an individual device health script assignment.
type DeviceHealthScriptAssignmentItem struct {
	ODataType            string                               `json:"@odata.type"`
	ID                   string                               `json:"id"`
	Target               ConfigurationManagerCollectionTarget `json:"target"`
	RunRemediationScript bool                                 `json:"runRemediationScript"`
	RunSchedule          DeviceHealthScriptAssignmentSchedule `json:"runSchedule"`
}

// ConfigurationManagerCollectionTarget represents the target for the device health script assignment.
type ConfigurationManagerCollectionTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	CollectionID                               string `json:"collectionId"`
}

// DeviceHealthScriptAssignmentSchedule represents the schedule for a device health script assignment.
type DeviceHealthScriptAssignmentSchedule struct {
	ODataType string `json:"@odata.type"`
	Interval  int    `json:"interval"`
	UseUTC    bool   `json:"useUtc"`
	Time      string `json:"time"`
}

// ResourceDeviceHealthScriptAssignment represents the request structure for creating a device health script assignment.
type ResourceDeviceHealthScriptAssignment struct {
	ODataType            string                               `json:"@odata.type"`
	Target               DeviceHealthScriptAssignmentTarget   `json:"target"`
	RunRemediationScript bool                                 `json:"runRemediationScript"`
	RunSchedule          DeviceHealthScriptAssignmentSchedule `json:"runSchedule"`
}

// DeviceHealthScriptAssignmentTarget represents the target for the device health script assignment.
type DeviceHealthScriptAssignmentTarget struct {
	ODataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterID   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	CollectionID                               string `json:"collectionId"`
}

// ResponseDeviceHealthScriptAssignment represents the response structure for a device health script assignment.
type ResponseDeviceHealthScriptAssignment struct {
	ODataType            string                               `json:"@odata.type"`
	ID                   string                               `json:"id"`
	Target               ConfigurationManagerCollectionTarget `json:"target"`
	RunRemediationScript bool                                 `json:"runRemediationScript"`
	RunSchedule          DeviceHealthScriptAssignmentSchedule `json:"runSchedule"`
}

// GetProactiveRemediationAssignments retrieves a list of Device Health Script Assignments from Microsoft Graph API.
func (c *Client) GetProactiveRemediationAssignments(scriptID string) (*ResponseDeviceHealthScriptAssignmentList, error) {
	endpoint := fmt.Sprintf("%s/%s/assignments", uriBetaProactiveRemediations, scriptID)

	var response ResponseDeviceHealthScriptAssignmentList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)

	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "proactive remediation script assignment", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetDeviceComplianceScriptAssignment retrieves a list of Device Compliance Script Assignments from Microsoft Graph API.
func (c *Client) GetDeviceComplianceScriptAssignments(scriptID string) (*ResponseDeviceHealthScriptAssignmentList, error) {
	endpoint := fmt.Sprintf("%s/%s/assignments", uriBetaDeviceComplianceScripts, scriptID)

	var response ResponseDeviceHealthScriptAssignmentList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)

	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device compliance script assignment", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
