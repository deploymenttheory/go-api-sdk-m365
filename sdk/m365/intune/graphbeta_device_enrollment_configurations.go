// graphbeta_device_enrollment_configurations.go
// Graph Beta Api - Intune: Device Enrollment
// Documentation: https://learn.microsoft.com/en-us/mem/intune/fundamentals/remediations
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/windowsEnrollment
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-shared-deviceenrollmentconfiguration?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const uriBetaDeviceEnrollmentConfigurations = "/beta/deviceManagement/deviceEnrollmentConfigurations"

// ResourceDeviceEnrollmentConfigurationsList represents the response structure for device enrollment configuration requests.
type ResourceDeviceEnrollmentConfigurationsList struct {
	Value []ResourceDeviceEnrollmentConfiguration `json:"value"`
}

// DeviceEnrollmentConfiguration represents a device enrollment configuration.
type ResourceDeviceEnrollmentConfiguration struct {
	OdataType            string    `json:"@odata.type"`
	ID                   string    `json:"id"`
	DisplayName          string    `json:"displayName"`
	Description          string    `json:"description"`
	Priority             int       `json:"priority"`
	CreatedDateTime      time.Time `json:"createdDateTime"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	Version              int       `json:"version"`
}

// GetDeviceEnrollmentConfigurations retrieves a list of all device enrollment configurations.
func (c *Client) GetDeviceEnrollmentConfigurations() (*ResourceDeviceEnrollmentConfigurationsList, error) {
	endpoint := uriBetaDeviceEnrollmentConfigurations

	var enrollmentConfigurations ResourceDeviceEnrollmentConfigurationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &enrollmentConfigurations)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device enrollment configurations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &enrollmentConfigurations, nil
}

// GetDeviceEnrollmentConfigurationByID retrieves a specific device enrollment configuration by its ID.
func (c *Client) GetDeviceEnrollmentConfigurationByID(id string) (*ResourceDeviceEnrollmentConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceEnrollmentConfigurations, id)

	var enrollmentConfiguration ResourceDeviceEnrollmentConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &enrollmentConfiguration)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device enrollment configuration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &enrollmentConfiguration, nil
}

// GetDeviceEnrollmentConfigurationByDisplayName retrieves a device management script by its display name.
func (c *Client) GetDeviceEnrollmentConfigurationByDisplayName(displayName string) (*ResourceDeviceEnrollmentConfiguration, error) {
	deviceEnrollmentConfigurations, err := c.GetDeviceEnrollmentConfigurations()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device enrollment configuration", err)
	}

	var deviceEnrollmentConfigurationID string
	for _, deviceEnrollmentConfiguration := range deviceEnrollmentConfigurations.Value {
		if deviceEnrollmentConfiguration.DisplayName == displayName {
			deviceEnrollmentConfigurationID = deviceEnrollmentConfiguration.ID
			break
		}
	}

	if deviceEnrollmentConfigurationID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device enrollment configuration", displayName, err)
	}

	return c.GetDeviceEnrollmentConfigurationByID(deviceEnrollmentConfigurationID)
}
