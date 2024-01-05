// graphbeta_device_management_iOS_configuration_profiles.go
// Graph Beta Api - Intune: iOS configuration profiles (Templates and Custom)
// Documentation: https://learn.microsoft.com/en-us/mem/intune/configuration/custom-settings-ios
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesIosMenu/~/configProfiles
// API reference: https://learn.microsoft.com/en-us/graph/api/intune-shared-deviceconfiguration-list?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

const (
	uriGraphBetaDeviceManagementIOSDeviceConfiguration = "/beta/deviceManagement/deviceConfigurations"
	odataTypeMacOSCustomConfigurationProfile           = "#microsoft.graph.macOSCustomConfiguration"
	odataTypeIOSCustomConfigurationProfile             = "#microsoft.graph.iosCustomConfiguration"
	odataTypeIOSTemplateConfigurationProfile           = "#microsoft.graph.iosGeneralDeviceConfiguration"
)
