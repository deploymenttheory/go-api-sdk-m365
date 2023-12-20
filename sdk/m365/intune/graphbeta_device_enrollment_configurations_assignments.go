// graphbeta_device_enrollment_configurations_assignments.go
// Graph Beta Api - Intune: Device Enrollment Assignments
// Documentation:
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/windowsEnrollment
// API reference: https://learn.microsoft.com/en-us/graph/api/intune-onboarding-enrollmentconfigurationassignment-list?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const uriBetaDeviceEnrollmentConfigurationAssignments = "/beta/deviceManagement/deviceEnrollmentConfigurations"

// ResourceDeviceEnrollmentConfigurationAssignmentsList represents the response structure for device enrollment configuration assignments.
type ResourceDeviceEnrollmentConfigurationAssignmentsList struct {
	Value []EnrollmentConfigurationAssignment `json:"value"`
}

// EnrollmentConfigurationAssignment represents an enrollment configuration assignment.
type EnrollmentConfigurationAssignment struct {
	OdataType string                     `json:"@odata.type"`
	ID        string                     `json:"id"`
	Target    EnrollmentAssignmentTarget `json:"target"`
	Source    string                     `json:"source"`
	SourceId  string                     `json:"sourceId"`
}

// EnrollmentAssignmentTarget represents the target for an enrollment configuration assignment.
type EnrollmentAssignmentTarget struct {
	OdataType                                  string `json:"@odata.type"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
	TargetType                                 string `json:"targetType"`
	EntraObjectId                              string `json:"entraObjectId"`
}

// GetDeviceEnrollmentConfigurationAssignmentsByDeviceEnrollmentConfigurationID retrieves all assignments for a device enrollment configuration by its ID.
func (c *Client) GetDeviceEnrollmentConfigurationAssignmentsByDeviceEnrollmentConfigurationID(configId string) (*ResourceDeviceEnrollmentConfigurationAssignmentsList, error) {
	endpoint := fmt.Sprintf("%s/%s/assignments", uriBetaDeviceEnrollmentConfigurationAssignments, configId)

	var assignments ResourceDeviceEnrollmentConfigurationAssignmentsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &assignments)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device enrollment configuration assignments", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &assignments, nil
}
