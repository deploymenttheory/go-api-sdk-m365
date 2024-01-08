// graphbeta_device_enrollment_shared.go
// Graph Beta Api - Intune: Device Enrollment Restrictions (SDK Shared resources)
// Documentation: https://learn.microsoft.com/en-us/mem/intune/enrollment/enrollment-restrictions-set
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_Enrollment/DeviceTypeRestrictionsBlade
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-shared-deviceenrollmentconfiguration?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import "time"

const (
	uriBetaGraphDeviceEnrollmentConfigurations = "/beta/deviceManagement/deviceEnrollmentConfigurations"
)

// ResponseDeviceEnrollmentConfigurationsList represents the response structure for device enrollment configuration requests.
type ResponseDeviceEnrollmentConfigurationsList struct {
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
