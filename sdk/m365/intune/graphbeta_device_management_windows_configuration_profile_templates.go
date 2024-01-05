// graphbeta_device_management_windows_configuration_profile_templates.go
// Graph Beta Api - Intune: Windows configuration profiles (Templates and Custom)
// Documentation: https://learn.microsoft.com/en-us/mem/intune/configuration/custom-settings-windows-10
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesWindowsMenu/~/configProfiles
// API reference: https://learn.microsoft.com/en-us/graph/api/intune-shared-deviceconfiguration-list?view=graph-rest-beta
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"
	"strings"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriGraphBetaDeviceManagementWindowsDeviceConfiguration = "/beta/deviceManagement/deviceConfigurations"
	odataTypeWindowsCustomConfigurationProfile             = "#microsoft.graph.windows10CustomConfiguration"
)

// ResourceDeviceConfigurationProfilesList represents a response containing a list of Device Configuration Profiles.
type ResourceDeviceConfigurationProfilesList struct {
	ODataContext       string                               `json:"@odata.context"`
	MicrosoftGraphTips string                               `json:"@microsoft.graph.tips"`
	Value              []ResourceDeviceConfigurationProfile `json:"value"`
}

// ResourceDeviceConfigurationProfile represents a single windows device configuration profile template.
type ResourceDeviceConfigurationProfile struct {
	ODataType                                   string                                       `json:"@odata.type"`
	ID                                          string                                       `json:"id"`
	CreatedDateTime                             string                                       `json:"createdDateTime"`
	LastModifiedDateTime                        string                                       `json:"lastModifiedDateTime"`
	Description                                 string                                       `json:"description"`
	DisplayName                                 string                                       `json:"displayName"`
	Version                                     int                                          `json:"version"`
	RoleScopeTagIds                             []string                                     `json:"roleScopeTagIds"`
	SupportsScopeTags                           bool                                         `json:"supportsScopeTags"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	// Fields for Template - Custom OMA Uri
	OmaSettings []DeviceConfigurationProfileOmaSetting `json:"omaSettings"`
	// Fields for Template - Delivery Optimization
	RestrictPeerSelectionBy                                   string   `json:"restrictPeerSelectionBy,omitempty"`
	GroupIdSource                                             string   `json:"groupIdSource,omitempty"`
	BackgroundDownloadFromHttpDelayInSeconds                  int      `json:"backgroundDownloadFromHttpDelayInSeconds,omitempty"`
	ForegroundDownloadFromHttpDelayInSeconds                  int      `json:"foregroundDownloadFromHttpDelayInSeconds,omitempty"`
	MinimumRamAllowedToPeerInGigabytes                        int      `json:"minimumRamAllowedToPeerInGigabytes,omitempty"`
	MinimumDiskSizeAllowedToPeerInGigabytes                   int      `json:"minimumDiskSizeAllowedToPeerInGigabytes,omitempty"`
	MinimumFileSizeToCacheInMegabytes                         int      `json:"minimumFileSizeToCacheInMegabytes,omitempty"`
	MinimumBatteryPercentageAllowedToUpload                   int      `json:"minimumBatteryPercentageAllowedToUpload,omitempty"`
	ModifyCacheLocation                                       string   `json:"modifyCacheLocation,omitempty"`
	MaximumCacheAgeInDays                                     int      `json:"maximumCacheAgeInDays,omitempty"`
	VpnPeerCaching                                            string   `json:"vpnPeerCaching,omitempty"`
	CacheServerHostNames                                      []string `json:"cacheServerHostNames,omitempty"`
	CacheServerForegroundDownloadFallbackToHttpDelayInSeconds int      `json:"cacheServerForegroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
	CacheServerBackgroundDownloadFallbackToHttpDelayInSeconds int      `json:"cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
	BandwidthMode                                             struct {
		MaximumDownloadBandwidthInKilobytesPerSecond int `json:"maximumDownloadBandwidthInKilobytesPerSecond,omitempty"`
		MaximumUploadBandwidthInKilobytesPerSecond   int `json:"maximumUploadBandwidthInKilobytesPerSecond,omitempty"`
	} `json:"bandwidthMode,omitempty"`
	MaximumCacheSize struct {
		MaximumCacheSizeInGigabytes int `json:"maximumCacheSizeInGigabytes,omitempty"`
	} `json:"maximumCacheSize,omitempty"`
	// Fields for Template - Identity protection
	UseSecurityKeyForSignin                      bool   `json:"useSecurityKeyForSignin,omitempty"`
	EnhancedAntiSpoofingForFacialFeaturesEnabled bool   `json:"enhancedAntiSpoofingForFacialFeaturesEnabled,omitempty"`
	PinMinimumLength                             *int   `json:"pinMinimumLength,omitempty"`
	PinMaximumLength                             *int   `json:"pinMaximumLength,omitempty"`
	PinUppercaseCharactersUsage                  string `json:"pinUppercaseCharactersUsage,omitempty"`
	PinLowercaseCharactersUsage                  string `json:"pinLowercaseCharactersUsage,omitempty"`
	PinSpecialCharactersUsage                    string `json:"pinSpecialCharactersUsage,omitempty"`
	PinExpirationInDays                          *int   `json:"pinExpirationInDays,omitempty"`
	PinPreviousBlockCount                        *int   `json:"pinPreviousBlockCount,omitempty"`
	PinRecoveryEnabled                           bool   `json:"pinRecoveryEnabled,omitempty"`
	SecurityDeviceRequired                       bool   `json:"securityDeviceRequired,omitempty"`
	UnlockWithBiometricsEnabled                  bool   `json:"unlockWithBiometricsEnabled,omitempty"`
	UseCertificatesForOnPremisesAuthEnabled      bool   `json:"useCertificatesForOnPremisesAuthEnabled,omitempty"`
	WindowsHelloForBusinessBlocked               bool   `json:"windowsHelloForBusinessBlocked,omitempty"`
	AssignmentsODataContext                      string `json:"assignments@odata.context,omitempty"`
	// Fields for Template - Device Firmware Configuration Interface
	ChangeUefiSettingsPermission   string `json:"changeUefiSettingsPermission,omitempty"`
	VirtualizationOfCpuAndIO       string `json:"virtualizationOfCpuAndIO,omitempty"`
	Cameras                        string `json:"cameras,omitempty"`
	MicrophonesAndSpeakers         string `json:"microphonesAndSpeakers,omitempty"`
	Radios                         string `json:"radios,omitempty"`
	BootFromExternalMedia          string `json:"bootFromExternalMedia,omitempty"`
	BootFromBuiltInNetworkAdapters string `json:"bootFromBuiltInNetworkAdapters,omitempty"`
	WindowsPlatformBinaryTable     string `json:"windowsPlatformBinaryTable,omitempty"`
	SimultaneousMultiThreading     string `json:"simultaneousMultiThreading,omitempty"`
	FrontCamera                    string `json:"frontCamera,omitempty"`
	RearCamera                     string `json:"rearCamera,omitempty"`
	InfraredCamera                 string `json:"infraredCamera,omitempty"`
	Microphone                     string `json:"microphone,omitempty"`
	Bluetooth                      string `json:"bluetooth,omitempty"`
	WirelessWideAreaNetwork        string `json:"wirelessWideAreaNetwork,omitempty"`
	NearFieldCommunication         string `json:"nearFieldCommunication,omitempty"`
	WiFi                           string `json:"wiFi,omitempty"`
	UsbTypeAPort                   string `json:"usbTypeAPort,omitempty"`
	SdCard                         string `json:"sdCard,omitempty"`
	WakeOnLAN                      string `json:"wakeOnLAN,omitempty"`
	WakeOnPower                    string `json:"wakeOnPower,omitempty"`
	// Fields for Template - Device Restrictions

	// configuration profile assignments
	Assignments []DeviceConfigurationProfileAssignment `json:"assignments"`
}

// DeviceManagementApplicabilityRuleOsEdition represents the OS edition applicability rule.
type DeviceManagementApplicabilityRuleOsEdition struct {
	ODataType      string   `json:"@odata.type"`
	OsEditionTypes []string `json:"osEditionTypes"`
	Name           string   `json:"name"`
	RuleType       string   `json:"ruleType"`
}

// DeviceManagementApplicabilityRuleOsVersion represents the OS version applicability rule.
type DeviceManagementApplicabilityRuleOsVersion struct {
	ODataType    string `json:"@odata.type"`
	MinOSVersion string `json:"minOSVersion"`
	MaxOSVersion string `json:"maxOSVersion"`
	Name         string `json:"name"`
	RuleType     string `json:"ruleType"`
}

// DeviceManagementApplicabilityRuleDeviceMode represents the device mode applicability rule.
type DeviceManagementApplicabilityRuleDeviceMode struct {
	ODataType  string `json:"@odata.type"`
	DeviceMode string `json:"deviceMode"`
	Name       string `json:"name"`
	RuleType   string `json:"ruleType"`
}

// Modify the DeviceConfigurationProfileOmaSetting struct to handle additional fields.
type DeviceConfigurationProfileOmaSetting struct {
	ODataType              string      `json:"@odata.type"`
	DisplayName            string      `json:"displayName"`
	Description            string      `json:"description"`
	OmaUri                 string      `json:"omaUri"`
	SecretReferenceValueId string      `json:"secretReferenceValueId"`
	IsEncrypted            bool        `json:"isEncrypted"`
	Value                  interface{} `json:"value"`
	IsReadOnly             bool        `json:"isReadOnly"`
	FileName               string      `json:"fileName"`
}

// DeviceConfigurationProfileAssignment represents an assignment for a Device Configuration Profile.
type DeviceConfigurationProfileAssignment struct {
	ID       string                                     `json:"id"`
	Source   string                                     `json:"source"`
	SourceId string                                     `json:"sourceId"`
	Intent   string                                     `json:"intent"`
	Target   DeviceConfigurationProfileAssignmentTarget `json:"target"`
}

// DeviceConfigurationProfileAssignmentTarget represents the target of a configuration profile assignment.
type DeviceConfigurationProfileAssignmentTarget struct {
	ODataType                                  string `json:"@odata.type"`
	GroupId                                    string `json:"groupId"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType"`
}

// GetWindowsDeviceConfigurationProfiles retrieves a list of Windows device configuration profiles from Microsoft Graph API.
// Because this is a shared endpoint, an OdataType match is used to filter the response so that only windows configuration
// profiles are returned.
func (c *Client) GetWindowsDeviceConfigurationProfiles() (*ResourceDeviceConfigurationProfilesList, error) {
	endpoint := uriGraphBetaDeviceManagementWindowsDeviceConfiguration + "?$expand=assignments"

	var responseDeviceConfigurationProfiles ResourceDeviceConfigurationProfilesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseDeviceConfigurationProfiles)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device configuration profiles", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Filter to include only Windows device profiles
	var windowsProfiles []ResourceDeviceConfigurationProfile
	for _, profile := range responseDeviceConfigurationProfiles.Value {
		if strings.HasPrefix(profile.ODataType, "#microsoft.graph.windows") {
			windowsProfiles = append(windowsProfiles, profile)
		}
	}

	responseDeviceConfigurationProfiles.Value = windowsProfiles
	return &responseDeviceConfigurationProfiles, nil
}

// GetWindowsDeviceConfigurationProfileByID retrieves a Windows device configuration profile by ID from Microsoft Graph API.
// This function verifies that the called profile ID corresponds to a Windows configuration profile.
// It also decrypts any encrypted OMA settings within the profile if present.
func (c *Client) GetWindowsDeviceConfigurationProfileByID(id string) (*ResourceDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/%s?$expand=assignments", uriGraphBetaDeviceManagementWindowsDeviceConfiguration, id)

	var responseDeviceConfigurationProfile ResourceDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseDeviceConfigurationProfile)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device configuration profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Check if the profile is a Windows device configuration profile
	if !strings.HasPrefix(responseDeviceConfigurationProfile.ODataType, "#microsoft.graph.windows") {
		return nil, fmt.Errorf("profile with ID %s is not a Windows device configuration profile", id)
	}

	// Check and decrypt any encrypted OMA settings
	for i, setting := range responseDeviceConfigurationProfile.OmaSettings {
		if setting.IsEncrypted {
			decryptedValue, err := c.GetDecryptedOmaSetting(uriGraphBetaDeviceManagementWindowsDeviceConfiguration, id, setting.SecretReferenceValueId)
			if err != nil {
				return nil, fmt.Errorf("failed to decrypt OMA setting: %v", err)
			}
			responseDeviceConfigurationProfile.OmaSettings[i].Value = decryptedValue
		}
	}

	return &responseDeviceConfigurationProfile, nil
}
