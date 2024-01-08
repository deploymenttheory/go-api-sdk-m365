// graphbeta_device_management_remote_assistance_settings.go
// Graph Beta Api - Intune Tenant Administration: Remote Help Settings
// Documentation: https://learn.microsoft.com/en-us/mem/intune/fundamentals/remote-help
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/TenantAdminMenu/~/remote
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-remoteassistance-remoteassistancesettings?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriGraphBetaDeviceManagementRemoteAssistanceSettings = "/beta/deviceManagement/remoteAssistanceSettings"
	odataTypeRemoteAssistanceSettings                    = "#microsoft.graph.remoteAssistanceSettings"
)

type ResponseRemoteAssistanceSettings struct {
	OdataContext                     string `json:"@odata.context"`
	ID                               string `json:"id"`
	RemoteAssistanceState            string `json:"remoteAssistanceState"`
	AllowSessionsToUnenrolledDevices bool   `json:"allowSessionsToUnenrolledDevices"`
	BlockChat                        bool   `json:"blockChat"`
}

// GetRemoteHelpSettings retrieves the settings for Intune Suite Remore Help
func (c *Client) GetRemoteHelpSettings() (*ResponseRemoteAssistanceSettings, error) {
	endpoint := uriGraphBetaDeviceManagementRemoteAssistanceSettings

	var responseRemoteHelpSettings ResponseRemoteAssistanceSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseRemoteHelpSettings)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "remote help settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseRemoteHelpSettings, nil
}

// UpdateRemoteHelpSettings retrieves the settings for Intune Suite Remore Help
func (c *Client) UpdateRemoteHelpSettings(updateRequest *ResponseRemoteAssistanceSettings) (*ResponseRemoteAssistanceSettings, error) {
	endpoint := uriGraphBetaDeviceManagementRemoteAssistanceSettings

	var responseRemoteHelpSettings ResponseRemoteAssistanceSettings
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, nil, &responseRemoteHelpSettings)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "remote help settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseRemoteHelpSettings, nil
}
