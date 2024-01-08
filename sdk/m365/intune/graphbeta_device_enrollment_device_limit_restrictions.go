// graphbeta_device_enrollment_restrictions.go
// Graph Beta Api - Intune: Device Enrollment Restrictions (Windows, Android, macOS and iOS)
// Documentation: https://learn.microsoft.com/en-us/mem/intune/enrollment/enrollment-restrictions-set
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_Enrollment/DeviceTypeRestrictionsBlade
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-shared-deviceenrollmentconfiguration?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	odataTypeDeviceEnrollmentLimitRestrictionsConfiguration = "#microsoft.graph.deviceEnrollmentLimitConfiguration"
)

// GetDeviceLimitRestrictions retrieves a list of all device enrollment restriction configurations
// with a specific OData type.
func (c *Client) GetDeviceLimitRestrictions() (*ResponseDeviceEnrollmentConfigurationsList, error) {
	endpoint := uriBetaGraphDeviceEnrollmentConfigurations

	var enrollmentConfigurations ResponseDeviceEnrollmentConfigurationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &enrollmentConfigurations)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device enrollment configurations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	var filteredConfigurations []ResourceDeviceEnrollmentConfiguration
	for _, config := range enrollmentConfigurations.Value {
		if config.OdataType == odataTypeDeviceEnrollmentLimitRestrictionsConfiguration {
			filteredConfigurations = append(filteredConfigurations, config)
		}
	}

	enrollmentConfigurations.Value = filteredConfigurations
	return &enrollmentConfigurations, nil
}

// GetDeviceLimitRestrictionByID retrieves a specific device enrollment configuration by its ID.
func (c *Client) GetDeviceLimitRestrictionByID(id string) (*ResourceDeviceEnrollmentConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBetaGraphDeviceEnrollmentConfigurations, id)

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

// GetDeviceLimitRestrictionByDisplayName retrieves a device management script by its display name.
func (c *Client) GetDeviceLimitRestrictionByDisplayName(displayName string) (*ResourceDeviceEnrollmentConfiguration, error) {
	deviceEnrollmentConfigurations, err := c.GetDeviceLimitRestrictions()
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

	return c.GetDeviceLimitRestrictionByID(deviceEnrollmentConfigurationID)
}
