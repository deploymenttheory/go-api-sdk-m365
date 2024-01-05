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
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriGraphBetaDeviceManagementWindowsDeviceConfiguration = "/beta/deviceManagement/deviceConfigurations"
	odataTypeWindowsCustomConfigurationProfile             = "#microsoft.graph.windows10CustomConfiguration"
)

// ResourceWindowsConfigurationProfileTemplatesList represents a response containing a list of Device Configuration Profiles.
type ResourceWindowsConfigurationProfileTemplatesList struct {
	ODataContext       string                                        `json:"@odata.context"`
	MicrosoftGraphTips string                                        `json:"@microsoft.graph.tips"`
	Value              []ResourceWindowsConfigurationProfileTemplate `json:"value"`
}

// ResourceWindowsConfigurationProfileTemplate represents a single windows device configuration profile template.
type ResourceWindowsConfigurationProfileTemplate struct {
	ODataType                                   string                                       `json:"@odata.type"`
	ID                                          string                                       `json:"id"`
	CreatedDateTime                             string                                       `json:"createdDateTime"`
	LastModifiedDateTime                        string                                       `json:"lastModifiedDateTime"`
	Description                                 string                                       `json:"description"`
	DisplayName                                 string                                       `json:"displayName"`
	Version                                     int                                          `json:"version"`
	RoleScopeTagIds                             []string                                     `json:"roleScopeTagIds"`
	SupportsScopeTags                           bool                                         `json:"supportsScopeTags"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode"`
	// Fields for Template - Custom OMA Uri
	OmaSettings []DeviceConfigurationProfileOmaSetting `json:"omaSettings,omitempty"`
	// Fields for Template - Delivery Optimization
	RestrictPeerSelectionBy                                   string                                     `json:"restrictPeerSelectionBy,omitempty"`
	GroupIdSource                                             string                                     `json:"groupIdSource,omitempty"`
	BackgroundDownloadFromHttpDelayInSeconds                  int                                        `json:"backgroundDownloadFromHttpDelayInSeconds,omitempty"`
	ForegroundDownloadFromHttpDelayInSeconds                  int                                        `json:"foregroundDownloadFromHttpDelayInSeconds,omitempty"`
	MinimumRamAllowedToPeerInGigabytes                        int                                        `json:"minimumRamAllowedToPeerInGigabytes,omitempty"`
	MinimumDiskSizeAllowedToPeerInGigabytes                   int                                        `json:"minimumDiskSizeAllowedToPeerInGigabytes,omitempty"`
	MinimumFileSizeToCacheInMegabytes                         int                                        `json:"minimumFileSizeToCacheInMegabytes,omitempty"`
	MinimumBatteryPercentageAllowedToUpload                   int                                        `json:"minimumBatteryPercentageAllowedToUpload,omitempty"`
	ModifyCacheLocation                                       string                                     `json:"modifyCacheLocation,omitempty"`
	MaximumCacheAgeInDays                                     int                                        `json:"maximumCacheAgeInDays,omitempty"`
	VpnPeerCaching                                            string                                     `json:"vpnPeerCaching,omitempty"`
	CacheServerHostNames                                      []string                                   `json:"cacheServerHostNames,omitempty"`
	CacheServerForegroundDownloadFallbackToHttpDelayInSeconds int                                        `json:"cacheServerForegroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
	CacheServerBackgroundDownloadFallbackToHttpDelayInSeconds int                                        `json:"cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
	BandwidthMode                                             DeliveryOptimizationSubsetBandwidthMode    `json:"bandwidthMode,omitempty"`
	MaximumCacheSize                                          DeliveryOptimizationSubsetMaximumCacheSize `json:"maximumCacheSize,omitempty"`
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
	TaskManagerBlockEndTask                               bool                                                   `json:"taskManagerBlockEndTask,omitempty"`
	EnergySaverOnBatteryThresholdPercentage               int                                                    `json:"energySaverOnBatteryThresholdPercentage,omitempty"`
	EnergySaverPluggedInThresholdPercentage               int                                                    `json:"energySaverPluggedInThresholdPercentage,omitempty"`
	PowerLidCloseActionOnBattery                          string                                                 `json:"powerLidCloseActionOnBattery,omitempty"`
	PowerLidCloseActionPluggedIn                          string                                                 `json:"powerLidCloseActionPluggedIn,omitempty"`
	PowerButtonActionOnBattery                            string                                                 `json:"powerButtonActionOnBattery,omitempty"`
	PowerButtonActionPluggedIn                            string                                                 `json:"powerButtonActionPluggedIn,omitempty"`
	PowerSleepButtonActionOnBattery                       string                                                 `json:"powerSleepButtonActionOnBattery,omitempty"`
	PowerSleepButtonActionPluggedIn                       string                                                 `json:"powerSleepButtonActionPluggedIn,omitempty"`
	PowerHybridSleepOnBattery                             string                                                 `json:"powerHybridSleepOnBattery,omitempty"`
	PowerHybridSleepPluggedIn                             string                                                 `json:"powerHybridSleepPluggedIn,omitempty"`
	Windows10AppsForceUpdateSchedule                      string                                                 `json:"windows10AppsForceUpdateSchedule,omitempty"`
	EnableAutomaticRedeployment                           bool                                                   `json:"enableAutomaticRedeployment,omitempty"`
	MicrosoftAccountSignInAssistantSettings               string                                                 `json:"microsoftAccountSignInAssistantSettings,omitempty"`
	AuthenticationAllowSecondaryDevice                    bool                                                   `json:"authenticationAllowSecondaryDevice,omitempty"`
	AuthenticationWebSignIn                               string                                                 `json:"authenticationWebSignIn,omitempty"`
	AuthenticationPreferredAzureADTenantDomainName        string                                                 `json:"authenticationPreferredAzureADTenantDomainName,omitempty"`
	CryptographyAllowFipsAlgorithmPolicy                  bool                                                   `json:"cryptographyAllowFipsAlgorithmPolicy,omitempty"`
	DisplayAppListWithGdiDPIScalingTurnedOn               []string                                               `json:"displayAppListWithGdiDPIScalingTurnedOn,omitempty"`
	DisplayAppListWithGdiDPIScalingTurnedOff              []string                                               `json:"displayAppListWithGdiDPIScalingTurnedOff,omitempty"`
	EnterpriseCloudPrintDiscoveryEndPoint                 string                                                 `json:"enterpriseCloudPrintDiscoveryEndPoint,omitempty"`
	EnterpriseCloudPrintOAuthAuthority                    string                                                 `json:"enterpriseCloudPrintOAuthAuthority,omitempty"`
	EnterpriseCloudPrintOAuthClientIdentifier             string                                                 `json:"enterpriseCloudPrintOAuthClientIdentifier,omitempty"`
	EnterpriseCloudPrintResourceIdentifier                string                                                 `json:"enterpriseCloudPrintResourceIdentifier,omitempty"`
	EnterpriseCloudPrintDiscoveryMaxLimit                 int                                                    `json:"enterpriseCloudPrintDiscoveryMaxLimit,omitempty"`
	EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier string                                                 `json:"enterpriseCloudPrintMopriaDiscoveryResourceIdentifier,omitempty"`
	ExperienceDoNotSyncBrowserSettings                    string                                                 `json:"experienceDoNotSyncBrowserSettings,omitempty"`
	MessagingBlockSync                                    bool                                                   `json:"messagingBlockSync,omitempty"`
	MessagingBlockMMS                                     bool                                                   `json:"messagingBlockMMS,omitempty"`
	MessagingBlockRichCommunicationServices               bool                                                   `json:"messagingBlockRichCommunicationServices,omitempty"`
	PrinterNames                                          []string                                               `json:"printerNames,omitempty"`
	PrinterDefaultName                                    string                                                 `json:"printerDefaultName,omitempty"`
	PrinterBlockAddition                                  bool                                                   `json:"printerBlockAddition,omitempty"`
	SearchBlockDiacritics                                 bool                                                   `json:"searchBlockDiacritics,omitempty"`
	SearchDisableAutoLanguageDetection                    bool                                                   `json:"searchDisableAutoLanguageDetection,omitempty"`
	SearchDisableIndexingEncryptedItems                   bool                                                   `json:"searchDisableIndexingEncryptedItems,omitempty"`
	SearchEnableRemoteQueries                             bool                                                   `json:"searchEnableRemoteQueries,omitempty"`
	SearchDisableUseLocation                              bool                                                   `json:"searchDisableUseLocation,omitempty"`
	SearchDisableLocation                                 bool                                                   `json:"searchDisableLocation,omitempty"`
	SearchDisableIndexerBackoff                           bool                                                   `json:"searchDisableIndexerBackoff,omitempty"`
	SearchDisableIndexingRemovableDrive                   bool                                                   `json:"searchDisableIndexingRemovableDrive,omitempty"`
	SearchEnableAutomaticIndexSizeManangement             bool                                                   `json:"searchEnableAutomaticIndexSizeManangement,omitempty"`
	SearchBlockWebResults                                 bool                                                   `json:"searchBlockWebResults,omitempty"`
	FindMyFiles                                           string                                                 `json:"findMyFiles,omitempty"`
	SecurityBlockAzureADJoinedDevicesAutoEncryption       bool                                                   `json:"securityBlockAzureADJoinedDevicesAutoEncryption,omitempty"`
	DiagnosticsDataSubmissionMode                         string                                                 `json:"diagnosticsDataSubmissionMode,omitempty"`
	OneDriveDisableFileSync                               bool                                                   `json:"oneDriveDisableFileSync,omitempty"`
	SystemTelemetryProxyServer                            string                                                 `json:"systemTelemetryProxyServer,omitempty"`
	EdgeTelemetryForMicrosoft365Analytics                 string                                                 `json:"edgeTelemetryForMicrosoft365Analytics,omitempty"`
	InkWorkspaceAccess                                    string                                                 `json:"inkWorkspaceAccess,omitempty"`
	InkWorkspaceAccessState                               string                                                 `json:"inkWorkspaceAccessState,omitempty"`
	InkWorkspaceBlockSuggestedApps                        bool                                                   `json:"inkWorkspaceBlockSuggestedApps,omitempty"`
	SmartScreenEnableAppInstallControl                    bool                                                   `json:"smartScreenEnableAppInstallControl,omitempty"`
	SmartScreenAppInstallControl                          string                                                 `json:"smartScreenAppInstallControl,omitempty"`
	PersonalizationDesktopImageURL                        string                                                 `json:"personalizationDesktopImageUrl,omitempty"`
	PersonalizationLockScreenImageURL                     string                                                 `json:"personalizationLockScreenImageUrl,omitempty"`
	BluetoothAllowedServices                              []string                                               `json:"bluetoothAllowedServices,omitempty"`
	BluetoothBlockAdvertising                             bool                                                   `json:"bluetoothBlockAdvertising,omitempty"`
	BluetoothBlockPromptedProximalConnections             bool                                                   `json:"bluetoothBlockPromptedProximalConnections,omitempty"`
	BluetoothBlockDiscoverableMode                        bool                                                   `json:"bluetoothBlockDiscoverableMode,omitempty"`
	BluetoothBlockPrePairing                              bool                                                   `json:"bluetoothBlockPrePairing,omitempty"`
	EdgeBlockAutofill                                     bool                                                   `json:"edgeBlockAutofill,omitempty"`
	EdgeBlocked                                           bool                                                   `json:"edgeBlocked,omitempty"`
	EdgeCookiePolicy                                      string                                                 `json:"edgeCookiePolicy,omitempty"`
	EdgeBlockDeveloperTools                               bool                                                   `json:"edgeBlockDeveloperTools,omitempty"`
	EdgeBlockSendingDoNotTrackHeader                      bool                                                   `json:"edgeBlockSendingDoNotTrackHeader,omitempty"`
	EdgeBlockExtensions                                   bool                                                   `json:"edgeBlockExtensions,omitempty"`
	EdgeBlockInPrivateBrowsing                            bool                                                   `json:"edgeBlockInPrivateBrowsing,omitempty"`
	EdgeBlockJavaScript                                   bool                                                   `json:"edgeBlockJavaScript,omitempty"`
	EdgeBlockPasswordManager                              bool                                                   `json:"edgeBlockPasswordManager,omitempty"`
	EdgeBlockAddressBarDropdown                           bool                                                   `json:"edgeBlockAddressBarDropdown,omitempty"`
	EdgeBlockCompatibilityList                            bool                                                   `json:"edgeBlockCompatibilityList,omitempty"`
	EdgeClearBrowsingDataOnExit                           bool                                                   `json:"edgeClearBrowsingDataOnExit,omitempty"`
	EdgeAllowStartPagesModification                       bool                                                   `json:"edgeAllowStartPagesModification,omitempty"`
	EdgeDisableFirstRunPage                               bool                                                   `json:"edgeDisableFirstRunPage,omitempty"`
	EdgeBlockLiveTileDataCollection                       bool                                                   `json:"edgeBlockLiveTileDataCollection,omitempty"`
	EdgeSyncFavoritesWithInternetExplorer                 bool                                                   `json:"edgeSyncFavoritesWithInternetExplorer,omitempty"`
	EdgeFavoritesListLocation                             string                                                 `json:"edgeFavoritesListLocation,omitempty"`
	EdgeBlockEditFavorites                                bool                                                   `json:"edgeBlockEditFavorites,omitempty"`
	EdgeNewTabPageURL                                     string                                                 `json:"edgeNewTabPageURL,omitempty"`
	EdgeHomeButtonConfiguration                           interface{}                                            `json:"edgeHomeButtonConfiguration,omitempty"`
	EdgeHomeButtonConfigurationEnabled                    bool                                                   `json:"edgeHomeButtonConfigurationEnabled,omitempty"`
	EdgeOpensWith                                         string                                                 `json:"edgeOpensWith,omitempty"`
	EdgeBlockSideloadingExtensions                        bool                                                   `json:"edgeBlockSideloadingExtensions,omitempty"`
	EdgeRequiredExtensionPackageFamilyNames               []string                                               `json:"edgeRequiredExtensionPackageFamilyNames,omitempty"`
	EdgeBlockPrinting                                     bool                                                   `json:"edgeBlockPrinting,omitempty"`
	EdgeFavoritesBarVisibility                            string                                                 `json:"edgeFavoritesBarVisibility,omitempty"`
	EdgeBlockSavingHistory                                bool                                                   `json:"edgeBlockSavingHistory,omitempty"`
	EdgeBlockFullScreenMode                               bool                                                   `json:"edgeBlockFullScreenMode,omitempty"`
	EdgeBlockWebContentOnNewTabPage                       bool                                                   `json:"edgeBlockWebContentOnNewTabPage,omitempty"`
	EdgeBlockTabPreloading                                bool                                                   `json:"edgeBlockTabPreloading,omitempty"`
	EdgeBlockPrelaunch                                    bool                                                   `json:"edgeBlockPrelaunch,omitempty"`
	EdgeShowMessageWhenOpeningInternetExplorerSites       string                                                 `json:"edgeShowMessageWhenOpeningInternetExplorerSites,omitempty"`
	EdgePreventCertificateErrorOverride                   bool                                                   `json:"edgePreventCertificateErrorOverride,omitempty"`
	EdgeKioskModeRestriction                              string                                                 `json:"edgeKioskModeRestriction,omitempty"`
	EdgeKioskResetAfterIdleTimeInMinutes                  int                                                    `json:"edgeKioskResetAfterIdleTimeInMinutes,omitempty"`
	CellularBlockDataWhenRoaming                          bool                                                   `json:"cellularBlockDataWhenRoaming,omitempty"`
	CellularBlockVpn                                      bool                                                   `json:"cellularBlockVpn,omitempty"`
	CellularBlockVpnWhenRoaming                           bool                                                   `json:"cellularBlockVpnWhenRoaming,omitempty"`
	CellularData                                          string                                                 `json:"cellularData,omitempty"`
	DefenderRequireRealTimeMonitoring                     bool                                                   `json:"defenderRequireRealTimeMonitoring,omitempty"`
	DefenderRequireBehaviorMonitoring                     bool                                                   `json:"defenderRequireBehaviorMonitoring,omitempty"`
	DefenderRequireNetworkInspectionSystem                bool                                                   `json:"defenderRequireNetworkInspectionSystem,omitempty"`
	DefenderScanDownloads                                 bool                                                   `json:"defenderScanDownloads,omitempty"`
	DefenderScheduleScanEnableLowCPUPriority              bool                                                   `json:"defenderScheduleScanEnableLowCpuPriority,omitempty"`
	DefenderDisableCatchupQuickScan                       bool                                                   `json:"defenderDisableCatchupQuickScan,omitempty"`
	DefenderDisableCatchupFullScan                        bool                                                   `json:"defenderDisableCatchupFullScan,omitempty"`
	DefenderScanScriptsLoadedInInternetExplorer           bool                                                   `json:"defenderScanScriptsLoadedInInternetExplorer,omitempty"`
	DefenderBlockEndUserAccess                            bool                                                   `json:"defenderBlockEndUserAccess,omitempty"`
	DefenderSignatureUpdateIntervalInHours                int                                                    `json:"defenderSignatureUpdateIntervalInHours,omitempty"`
	DefenderMonitorFileActivity                           string                                                 `json:"defenderMonitorFileActivity,omitempty"`
	DefenderDaysBeforeDeletingQuarantinedMalware          int                                                    `json:"defenderDaysBeforeDeletingQuarantinedMalware,omitempty"`
	DefenderScanMaxCPU                                    int                                                    `json:"defenderScanMaxCpu,omitempty"`
	DefenderScanArchiveFiles                              bool                                                   `json:"defenderScanArchiveFiles,omitempty"`
	DefenderScanIncomingMail                              bool                                                   `json:"defenderScanIncomingMail,omitempty"`
	DefenderScanRemovableDrivesDuringFullScan             bool                                                   `json:"defenderScanRemovableDrivesDuringFullScan,omitempty"`
	DefenderScanMappedNetworkDrivesDuringFullScan         bool                                                   `json:"defenderScanMappedNetworkDrivesDuringFullScan,omitempty"`
	DefenderScanNetworkFiles                              bool                                                   `json:"defenderScanNetworkFiles,omitempty"`
	DefenderRequireCloudProtection                        bool                                                   `json:"defenderRequireCloudProtection,omitempty"`
	DefenderCloudBlockLevel                               string                                                 `json:"defenderCloudBlockLevel,omitempty"`
	DefenderCloudExtendedTimeout                          int                                                    `json:"defenderCloudExtendedTimeout,omitempty"`
	DefenderCloudExtendedTimeoutInSeconds                 int                                                    `json:"defenderCloudExtendedTimeoutInSeconds,omitempty"`
	DefenderPromptForSampleSubmission                     string                                                 `json:"defenderPromptForSampleSubmission,omitempty"`
	DefenderScheduledQuickScanTime                        string                                                 `json:"defenderScheduledQuickScanTime,omitempty"`
	DefenderScanType                                      string                                                 `json:"defenderScanType,omitempty"`
	DefenderSystemScanSchedule                            string                                                 `json:"defenderSystemScanSchedule,omitempty"`
	DefenderScheduledScanTime                             string                                                 `json:"defenderScheduledScanTime,omitempty"`
	DefenderPotentiallyUnwantedAppAction                  string                                                 `json:"defenderPotentiallyUnwantedAppAction,omitempty"`
	DefenderPotentiallyUnwantedAppActionSetting           string                                                 `json:"defenderPotentiallyUnwantedAppActionSetting,omitempty"`
	DefenderSubmitSamplesConsentType                      interface{}                                            `json:"defenderSubmitSamplesConsentType,omitempty"`
	DefenderBlockOnAccessProtection                       bool                                                   `json:"defenderBlockOnAccessProtection,omitempty"`
	DefenderFileExtensionsToExclude                       []string                                               `json:"defenderFileExtensionsToExclude,omitempty"`
	DefenderFilesAndFoldersToExclude                      []string                                               `json:"defenderFilesAndFoldersToExclude,omitempty"`
	DefenderProcessesToExclude                            []string                                               `json:"defenderProcessesToExclude,omitempty"`
	LockScreenAllowTimeoutConfiguration                   bool                                                   `json:"lockScreenAllowTimeoutConfiguration,omitempty"`
	LockScreenBlockActionCenterNotifications              bool                                                   `json:"lockScreenBlockActionCenterNotifications,omitempty"`
	LockScreenBlockCortana                                bool                                                   `json:"lockScreenBlockCortana,omitempty"`
	LockScreenBlockToastNotifications                     bool                                                   `json:"lockScreenBlockToastNotifications,omitempty"`
	LockScreenTimeoutInSeconds                            int                                                    `json:"lockScreenTimeoutInSeconds,omitempty"`
	LockScreenActivateAppsWithVoice                       string                                                 `json:"lockScreenActivateAppsWithVoice,omitempty"`
	PasswordBlockSimple                                   bool                                                   `json:"passwordBlockSimple,omitempty"`
	PasswordExpirationDays                                int                                                    `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumLength                                 int                                                    `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeScreenTimeout        int                                                    `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasswordMinimumCharacterSetCount                      int                                                    `json:"passwordMinimumCharacterSetCount,omitempty"`
	PasswordPreviousPasswordBlockCount                    int                                                    `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequired                                      bool                                                   `json:"passwordRequired,omitempty"`
	PasswordRequireWhenResumeFromIdleState                bool                                                   `json:"passwordRequireWhenResumeFromIdleState,omitempty"`
	PasswordRequiredType                                  string                                                 `json:"passwordRequiredType,omitempty"`
	PasswordSignInFailureCountBeforeFactoryReset          int                                                    `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	PasswordMinimumAgeInDays                              interface{}                                            `json:"passwordMinimumAgeInDays,omitempty"`
	PrivacyAdvertisingID                                  string                                                 `json:"privacyAdvertisingId,omitempty"`
	PrivacyAutoAcceptPairingAndConsentPrompts             bool                                                   `json:"privacyAutoAcceptPairingAndConsentPrompts,omitempty"`
	PrivacyDisableLaunchExperience                        bool                                                   `json:"privacyDisableLaunchExperience,omitempty"`
	PrivacyBlockInputPersonalization                      bool                                                   `json:"privacyBlockInputPersonalization,omitempty"`
	PrivacyBlockPublishUserActivities                     bool                                                   `json:"privacyBlockPublishUserActivities,omitempty"`
	PrivacyBlockActivityFeed                              bool                                                   `json:"privacyBlockActivityFeed,omitempty"`
	ActivateAppsWithVoice                                 string                                                 `json:"activateAppsWithVoice,omitempty"`
	StartBlockUnpinningAppsFromTaskbar                    bool                                                   `json:"startBlockUnpinningAppsFromTaskbar,omitempty"`
	StartMenuAppListVisibility                            string                                                 `json:"startMenuAppListVisibility,omitempty"`
	StartMenuHideChangeAccountSettings                    bool                                                   `json:"startMenuHideChangeAccountSettings,omitempty"`
	StartMenuHideFrequentlyUsedApps                       bool                                                   `json:"startMenuHideFrequentlyUsedApps,omitempty"`
	StartMenuHideHibernate                                bool                                                   `json:"startMenuHideHibernate,omitempty"`
	StartMenuHideLock                                     bool                                                   `json:"startMenuHideLock,omitempty"`
	StartMenuHidePowerButton                              bool                                                   `json:"startMenuHidePowerButton,omitempty"`
	StartMenuHideRecentJumpLists                          bool                                                   `json:"startMenuHideRecentJumpLists,omitempty"`
	StartMenuHideRecentlyAddedApps                        bool                                                   `json:"startMenuHideRecentlyAddedApps,omitempty"`
	StartMenuHideRestartOptions                           bool                                                   `json:"startMenuHideRestartOptions,omitempty"`
	StartMenuHideShutDown                                 bool                                                   `json:"startMenuHideShutDown,omitempty"`
	StartMenuHideSignOut                                  bool                                                   `json:"startMenuHideSignOut,omitempty"`
	StartMenuHideSleep                                    bool                                                   `json:"startMenuHideSleep,omitempty"`
	StartMenuHideSwitchAccount                            bool                                                   `json:"startMenuHideSwitchAccount,omitempty"`
	StartMenuHideUserTile                                 bool                                                   `json:"startMenuHideUserTile,omitempty"`
	StartMenuLayoutEdgeAssetsXML                          interface{}                                            `json:"startMenuLayoutEdgeAssetsXml,omitempty"`
	StartMenuLayoutXML                                    interface{}                                            `json:"startMenuLayoutXml,omitempty"`
	StartMenuMode                                         string                                                 `json:"startMenuMode,omitempty"`
	StartMenuPinnedFolderDocuments                        string                                                 `json:"startMenuPinnedFolderDocuments,omitempty"`
	StartMenuPinnedFolderDownloads                        string                                                 `json:"startMenuPinnedFolderDownloads,omitempty"`
	StartMenuPinnedFolderFileExplorer                     string                                                 `json:"startMenuPinnedFolderFileExplorer,omitempty"`
	StartMenuPinnedFolderHomeGroup                        string                                                 `json:"startMenuPinnedFolderHomeGroup,omitempty"`
	StartMenuPinnedFolderMusic                            string                                                 `json:"startMenuPinnedFolderMusic,omitempty"`
	StartMenuPinnedFolderNetwork                          string                                                 `json:"startMenuPinnedFolderNetwork,omitempty"`
	StartMenuPinnedFolderPersonalFolder                   string                                                 `json:"startMenuPinnedFolderPersonalFolder,omitempty"`
	StartMenuPinnedFolderPictures                         string                                                 `json:"startMenuPinnedFolderPictures,omitempty"`
	StartMenuPinnedFolderSettings                         string                                                 `json:"startMenuPinnedFolderSettings,omitempty"`
	StartMenuPinnedFolderVideos                           string                                                 `json:"startMenuPinnedFolderVideos,omitempty"`
	SettingsBlockSettingsApp                              bool                                                   `json:"settingsBlockSettingsApp,omitempty"`
	SettingsBlockSystemPage                               bool                                                   `json:"settingsBlockSystemPage,omitempty"`
	SettingsBlockDevicesPage                              bool                                                   `json:"settingsBlockDevicesPage,omitempty"`
	SettingsBlockNetworkInternetPage                      bool                                                   `json:"settingsBlockNetworkInternetPage,omitempty"`
	SettingsBlockPersonalizationPage                      bool                                                   `json:"settingsBlockPersonalizationPage,omitempty"`
	SettingsBlockAccountsPage                             bool                                                   `json:"settingsBlockAccountsPage,omitempty"`
	SettingsBlockTimeLanguagePage                         bool                                                   `json:"settingsBlockTimeLanguagePage,omitempty"`
	SettingsBlockEaseOfAccessPage                         bool                                                   `json:"settingsBlockEaseOfAccessPage,omitempty"`
	SettingsBlockPrivacyPage                              bool                                                   `json:"settingsBlockPrivacyPage,omitempty"`
	SettingsBlockUpdateSecurityPage                       bool                                                   `json:"settingsBlockUpdateSecurityPage,omitempty"`
	SettingsBlockAppsPage                                 bool                                                   `json:"settingsBlockAppsPage,omitempty"`
	SettingsBlockGamingPage                               bool                                                   `json:"settingsBlockGamingPage,omitempty"`
	WindowsSpotlightBlockConsumerSpecificFeatures         bool                                                   `json:"windowsSpotlightBlockConsumerSpecificFeatures,omitempty"`
	WindowsSpotlightBlocked                               bool                                                   `json:"windowsSpotlightBlocked,omitempty"`
	WindowsSpotlightBlockOnActionCenter                   bool                                                   `json:"windowsSpotlightBlockOnActionCenter,omitempty"`
	WindowsSpotlightBlockTailoredExperiences              bool                                                   `json:"windowsSpotlightBlockTailoredExperiences,omitempty"`
	WindowsSpotlightBlockThirdPartyNotifications          bool                                                   `json:"windowsSpotlightBlockThirdPartyNotifications,omitempty"`
	WindowsSpotlightBlockWelcomeExperience                bool                                                   `json:"windowsSpotlightBlockWelcomeExperience,omitempty"`
	WindowsSpotlightBlockWindowsTips                      bool                                                   `json:"windowsSpotlightBlockWindowsTips,omitempty"`
	WindowsSpotlightConfigureOnLockScreen                 string                                                 `json:"windowsSpotlightConfigureOnLockScreen,omitempty"`
	NetworkProxyApplySettingsDeviceWide                   bool                                                   `json:"networkProxyApplySettingsDeviceWide,omitempty"`
	NetworkProxyDisableAutoDetect                         bool                                                   `json:"networkProxyDisableAutoDetect,omitempty"`
	NetworkProxyAutomaticConfigurationURL                 string                                                 `json:"networkProxyAutomaticConfigurationUrl,omitempty"`
	AccountsBlockAddingNonMicrosoftAccountEmail           bool                                                   `json:"accountsBlockAddingNonMicrosoftAccountEmail,omitempty"`
	AntiTheftModeBlocked                                  bool                                                   `json:"antiTheftModeBlocked,omitempty"`
	BluetoothBlocked                                      bool                                                   `json:"bluetoothBlocked,omitempty"`
	CameraBlocked                                         bool                                                   `json:"cameraBlocked,omitempty"`
	ConnectedDevicesServiceBlocked                        bool                                                   `json:"connectedDevicesServiceBlocked,omitempty"`
	CertificatesBlockManualRootCertificateInstallation    bool                                                   `json:"certificatesBlockManualRootCertificateInstallation,omitempty"`
	CopyPasteBlocked                                      bool                                                   `json:"copyPasteBlocked,omitempty"`
	CortanaBlocked                                        bool                                                   `json:"cortanaBlocked,omitempty"`
	DeviceManagementBlockFactoryResetOnMobile             bool                                                   `json:"deviceManagementBlockFactoryResetOnMobile,omitempty"`
	DeviceManagementBlockManualUnenroll                   bool                                                   `json:"deviceManagementBlockManualUnenroll,omitempty"`
	SafeSearchFilter                                      string                                                 `json:"safeSearchFilter,omitempty"`
	EdgeBlockPopups                                       bool                                                   `json:"edgeBlockPopups,omitempty"`
	EdgeBlockSearchSuggestions                            bool                                                   `json:"edgeBlockSearchSuggestions,omitempty"`
	EdgeBlockSearchEngineCustomization                    bool                                                   `json:"edgeBlockSearchEngineCustomization,omitempty"`
	EdgeBlockSendingIntranetTrafficToInternetExplorer     bool                                                   `json:"edgeBlockSendingIntranetTrafficToInternetExplorer,omitempty"`
	EdgeSendIntranetTrafficToInternetExplorer             bool                                                   `json:"edgeSendIntranetTrafficToInternetExplorer,omitempty"`
	EdgeRequireSmartScreen                                bool                                                   `json:"edgeRequireSmartScreen,omitempty"`
	EdgeEnterpriseModeSiteListLocation                    string                                                 `json:"edgeEnterpriseModeSiteListLocation,omitempty"`
	EdgeFirstRunURL                                       string                                                 `json:"edgeFirstRunUrl,omitempty"`
	EdgeHomepageUrls                                      []string                                               `json:"edgeHomepageUrls,omitempty"`
	EdgeBlockAccessToAboutFlags                           bool                                                   `json:"edgeBlockAccessToAboutFlags,omitempty"`
	SmartScreenBlockPromptOverride                        bool                                                   `json:"smartScreenBlockPromptOverride,omitempty"`
	SmartScreenBlockPromptOverrideForFiles                bool                                                   `json:"smartScreenBlockPromptOverrideForFiles,omitempty"`
	WebRtcBlockLocalhostIPAddress                         bool                                                   `json:"webRtcBlockLocalhostIpAddress,omitempty"`
	InternetSharingBlocked                                bool                                                   `json:"internetSharingBlocked,omitempty"`
	SettingsBlockAddProvisioningPackage                   bool                                                   `json:"settingsBlockAddProvisioningPackage,omitempty"`
	SettingsBlockRemoveProvisioningPackage                bool                                                   `json:"settingsBlockRemoveProvisioningPackage,omitempty"`
	SettingsBlockChangeSystemTime                         bool                                                   `json:"settingsBlockChangeSystemTime,omitempty"`
	SettingsBlockEditDeviceName                           bool                                                   `json:"settingsBlockEditDeviceName,omitempty"`
	SettingsBlockChangeRegion                             bool                                                   `json:"settingsBlockChangeRegion,omitempty"`
	SettingsBlockChangeLanguage                           bool                                                   `json:"settingsBlockChangeLanguage,omitempty"`
	SettingsBlockChangePowerSleep                         bool                                                   `json:"settingsBlockChangePowerSleep,omitempty"`
	LocationServicesBlocked                               bool                                                   `json:"locationServicesBlocked,omitempty"`
	MicrosoftAccountBlocked                               bool                                                   `json:"microsoftAccountBlocked,omitempty"`
	MicrosoftAccountBlockSettingsSync                     bool                                                   `json:"microsoftAccountBlockSettingsSync,omitempty"`
	NfcBlocked                                            bool                                                   `json:"nfcBlocked,omitempty"`
	ResetProtectionModeBlocked                            bool                                                   `json:"resetProtectionModeBlocked,omitempty"`
	ScreenCaptureBlocked                                  bool                                                   `json:"screenCaptureBlocked,omitempty"`
	StorageBlockRemovableStorage                          bool                                                   `json:"storageBlockRemovableStorage,omitempty"`
	StorageRequireMobileDeviceEncryption                  bool                                                   `json:"storageRequireMobileDeviceEncryption,omitempty"`
	UsbBlocked                                            bool                                                   `json:"usbBlocked,omitempty"`
	VoiceRecordingBlocked                                 bool                                                   `json:"voiceRecordingBlocked,omitempty"`
	WiFiBlockAutomaticConnectHotspots                     bool                                                   `json:"wiFiBlockAutomaticConnectHotspots,omitempty"`
	WiFiBlocked                                           bool                                                   `json:"wiFiBlocked,omitempty"`
	WiFiBlockManualConfiguration                          bool                                                   `json:"wiFiBlockManualConfiguration,omitempty"`
	WiFiScanInterval                                      int                                                    `json:"wiFiScanInterval,omitempty"`
	WirelessDisplayBlockProjectionToThisDevice            bool                                                   `json:"wirelessDisplayBlockProjectionToThisDevice,omitempty"`
	WirelessDisplayBlockUserInputFromReceiver             bool                                                   `json:"wirelessDisplayBlockUserInputFromReceiver,omitempty"`
	WirelessDisplayRequirePinForPairing                   bool                                                   `json:"wirelessDisplayRequirePinForPairing,omitempty"`
	WindowsStoreBlocked                                   bool                                                   `json:"windowsStoreBlocked,omitempty"`
	AppsAllowTrustedAppsSideloading                       string                                                 `json:"appsAllowTrustedAppsSideloading,omitempty"`
	WindowsStoreBlockAutoUpdate                           bool                                                   `json:"windowsStoreBlockAutoUpdate,omitempty"`
	DeveloperUnlockSetting                                string                                                 `json:"developerUnlockSetting,omitempty"`
	SharedUserAppDataAllowed                              bool                                                   `json:"sharedUserAppDataAllowed,omitempty"`
	AppsBlockWindowsStoreOriginatedApps                   bool                                                   `json:"appsBlockWindowsStoreOriginatedApps,omitempty"`
	WindowsStoreEnablePrivateStoreOnly                    bool                                                   `json:"windowsStoreEnablePrivateStoreOnly,omitempty"`
	StorageRestrictAppDataToSystemVolume                  bool                                                   `json:"storageRestrictAppDataToSystemVolume,omitempty"`
	StorageRestrictAppInstallToSystemVolume               bool                                                   `json:"storageRestrictAppInstallToSystemVolume,omitempty"`
	GameDvrBlocked                                        bool                                                   `json:"gameDvrBlocked,omitempty"`
	ExperienceBlockDeviceDiscovery                        bool                                                   `json:"experienceBlockDeviceDiscovery,omitempty"`
	ExperienceBlockErrorDialogWhenNoSIM                   bool                                                   `json:"experienceBlockErrorDialogWhenNoSIM,omitempty"`
	ExperienceBlockTaskSwitcher                           bool                                                   `json:"experienceBlockTaskSwitcher,omitempty"`
	LogonBlockFastUserSwitching                           bool                                                   `json:"logonBlockFastUserSwitching,omitempty"`
	TenantLockdownRequireNetworkDuringOutOfBoxExperience  bool                                                   `json:"tenantLockdownRequireNetworkDuringOutOfBoxExperience,omitempty"`
	AppManagementMSIAllowUserControlOverInstall           bool                                                   `json:"appManagementMSIAllowUserControlOverInstall,omitempty"`
	AppManagementMSIAlwaysInstallWithElevatedPrivileges   bool                                                   `json:"appManagementMSIAlwaysInstallWithElevatedPrivileges,omitempty"`
	DataProtectionBlockDirectMemoryAccess                 bool                                                   `json:"dataProtectionBlockDirectMemoryAccess,omitempty"`
	AppManagementPackageFamilyNamesToLaunchAfterLogOn     []interface{}                                          `json:"appManagementPackageFamilyNamesToLaunchAfterLogOn,omitempty"`
	UninstallBuiltInApps                                  bool                                                   `json:"uninstallBuiltInApps,omitempty"`
	ConfigureTimeZone                                     *DeviceRestrictionsSubsetConfigureTimeZone             `json:"configureTimeZone,omitempty"`
	DefenderDetectedMalwareActions                        DeviceRestrictionsSubsetDefenderDetectedMalwareActions `json:"defenderDetectedMalwareActions,omitempty"`
	NetworkProxyServer                                    DeviceRestrictionsSubsetNetworkProxyServer             `json:"networkProxyServer,omitempty"`
	EdgeSearchEngine                                      DeviceRestrictionsSubsetEdgeSearchEngine               `json:"edgeSearchEngine,omitempty"`
	// Fields for Template - Device restrictions (Windows 10 Team)
	AzureOperationalInsightsBlockTelemetry bool   `json:"azureOperationalInsightsBlockTelemetry,omitempty"`
	AzureOperationalInsightsWorkspaceId    string `json:"azureOperationalInsightsWorkspaceId,omitempty"`
	AzureOperationalInsightsWorkspaceKey   string `json:"azureOperationalInsightsWorkspaceKey,omitempty"`
	ConnectAppBlockAutoLaunch              bool   `json:"connectAppBlockAutoLaunch,omitempty"`
	MaintenanceWindowBlocked               bool   `json:"maintenanceWindowBlocked,omitempty"`
	MaintenanceWindowDurationInHours       int    `json:"maintenanceWindowDurationInHours,omitempty"`
	MaintenanceWindowStartTime             string `json:"maintenanceWindowStartTime,omitempty"`
	MiracastChannel                        string `json:"miracastChannel,omitempty"`
	MiracastBlocked                        bool   `json:"miracastBlocked,omitempty"`
	MiracastRequirePin                     bool   `json:"miracastRequirePin,omitempty"`
	SettingsBlockMyMeetingsAndFiles        bool   `json:"settingsBlockMyMeetingsAndFiles,omitempty"`
	SettingsBlockSessionResume             bool   `json:"settingsBlockSessionResume,omitempty"`
	SettingsBlockSigninSuggestions         bool   `json:"settingsBlockSigninSuggestions,omitempty"`
	SettingsDefaultVolume                  int    `json:"settingsDefaultVolume,omitempty"`
	SettingsScreenTimeoutInMinutes         int    `json:"settingsScreenTimeoutInMinutes,omitempty"`
	SettingsSessionTimeoutInMinutes        int    `json:"settingsSessionTimeoutInMinutes,omitempty"`
	SettingsSleepTimeoutInMinutes          int    `json:"settingsSleepTimeoutInMinutes,omitempty"`
	WelcomeScreenBlockAutomaticWakeUp      bool   `json:"welcomeScreenBlockAutomaticWakeUp,omitempty"`
	WelcomeScreenBackgroundImageUrl        string `json:"welcomeScreenBackgroundImageUrl,omitempty"`
	WelcomeScreenMeetingInformation        string `json:"welcomeScreenMeetingInformation,omitempty"`
	// Fields for Template - Domain Join
	ComputerNameStaticPrefix          string `json:"computerNameStaticPrefix,omitempty"`
	ComputerNameSuffixRandomCharCount int    `json:"computerNameSuffixRandomCharCount,omitempty"`
	ActiveDirectoryDomainName         string `json:"activeDirectoryDomainName,omitempty"`
	OrganizationalUnit                string `json:"organizationalUnit,omitempty"`
	// Fields for Template - Edition upgrade and mode switch
	LicenseType   string `json:"licenseType,omitempty"`
	TargetEdition string `json:"targetEdition,omitempty"`
	License       string `json:"license,omitempty"`
	ProductKey    string `json:"productKey,omitempty"`
	WindowsSMode  string `json:"windowsSMode,omitempty"`
	// Fields for Template - Email
	UsernameSource        string `json:"usernameSource,omitempty"`
	UsernameAADSource     string `json:"usernameAADSource,omitempty"`
	UserDomainNameSource  string `json:"userDomainNameSource,omitempty"`
	CustomDomainName      string `json:"customDomainName,omitempty"`
	AccountName           string `json:"accountName,omitempty"`
	SyncCalendar          bool   `json:"syncCalendar,omitempty"`
	SyncContacts          bool   `json:"syncContacts,omitempty"`
	SyncTasks             bool   `json:"syncTasks,omitempty"`
	DurationOfEmailToSync string `json:"durationOfEmailToSync,omitempty"`
	EmailAddressSource    string `json:"emailAddressSource,omitempty"`
	EmailSyncSchedule     string `json:"emailSyncSchedule,omitempty"`
	HostName              string `json:"hostName,omitempty"`
	RequireSsl            bool   `json:"requireSsl,omitempty"`
	// Fields for Template Endpoint Protection
	DmaGuardDeviceEnumerationPolicy                                              string                                  `json:"dmaGuardDeviceEnumerationPolicy,omitempty"`
	UserRightsAccessCredentialManagerAsTrustedCaller                             interface{}                             `json:"userRightsAccessCredentialManagerAsTrustedCaller,omitempty"`
	UserRightsBlockAccessFromNetwork                                             interface{}                             `json:"userRightsBlockAccessFromNetwork,omitempty"`
	UserRightsDenyLocalLogOn                                                     interface{}                             `json:"userRightsDenyLocalLogOn,omitempty"`
	UserRightsDebugPrograms                                                      interface{}                             `json:"userRightsDebugPrograms,omitempty"`
	XboxServicesEnableXboxGameSaveTask                                           bool                                    `json:"xboxServicesEnableXboxGameSaveTask,omitempty"`
	XboxServicesAccessoryManagementServiceStartupMode                            string                                  `json:"xboxServicesAccessoryManagementServiceStartupMode,omitempty"`
	XboxServicesLiveAuthManagerServiceStartupMode                                string                                  `json:"xboxServicesLiveAuthManagerServiceStartupMode,omitempty"`
	XboxServicesLiveGameSaveServiceStartupMode                                   string                                  `json:"xboxServicesLiveGameSaveServiceStartupMode,omitempty"`
	XboxServicesLiveNetworkingServiceStartupMode                                 string                                  `json:"xboxServicesLiveNetworkingServiceStartupMode,omitempty"`
	LocalSecurityOptionsBlockMicrosoftAccounts                                   bool                                    `json:"localSecurityOptionsBlockMicrosoftAccounts,omitempty"`
	LocalSecurityOptionsBlockRemoteLogonWithBlankPassword                        bool                                    `json:"localSecurityOptionsBlockRemoteLogonWithBlankPassword,omitempty"`
	LocalSecurityOptionsDisableAdministratorAccount                              bool                                    `json:"localSecurityOptionsDisableAdministratorAccount,omitempty"`
	LocalSecurityOptionsAdministratorAccountName                                 string                                  `json:"localSecurityOptionsAdministratorAccountName,omitempty"`
	LocalSecurityOptionsDisableGuestAccount                                      bool                                    `json:"localSecurityOptionsDisableGuestAccount,omitempty"`
	LocalSecurityOptionsGuestAccountName                                         string                                  `json:"localSecurityOptionsGuestAccountName,omitempty"`
	LocalSecurityOptionsAllowUndockWithoutHavingToLogon                          bool                                    `json:"localSecurityOptionsAllowUndockWithoutHavingToLogon,omitempty"`
	LocalSecurityOptionsBlockUsersInstallingPrinterDrivers                       bool                                    `json:"localSecurityOptionsBlockUsersInstallingPrinterDrivers,omitempty"`
	LocalSecurityOptionsBlockRemoteOpticalDriveAccess                            bool                                    `json:"localSecurityOptionsBlockRemoteOpticalDriveAccess,omitempty"`
	LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser                string                                  `json:"localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser,omitempty"`
	LocalSecurityOptionsMachineInactivityLimit                                   int                                     `json:"localSecurityOptionsMachineInactivityLimit,omitempty"`
	LocalSecurityOptionsMachineInactivityLimitInMinutes                          int                                     `json:"localSecurityOptionsMachineInactivityLimitInMinutes,omitempty"`
	LocalSecurityOptionsDoNotRequireCtrlAltDel                                   bool                                    `json:"localSecurityOptionsDoNotRequireCtrlAltDel,omitempty"`
	LocalSecurityOptionsHideLastSignedInUser                                     bool                                    `json:"localSecurityOptionsHideLastSignedInUser,omitempty"`
	LocalSecurityOptionsHideUsernameAtSignIn                                     bool                                    `json:"localSecurityOptionsHideUsernameAtSignIn,omitempty"`
	LocalSecurityOptionsLogOnMessageTitle                                        interface{}                             `json:"localSecurityOptionsLogOnMessageTitle,omitempty"`
	LocalSecurityOptionsLogOnMessageText                                         interface{}                             `json:"localSecurityOptionsLogOnMessageText,omitempty"`
	LocalSecurityOptionsAllowPKU2UAuthenticationRequests                         bool                                    `json:"localSecurityOptionsAllowPKU2UAuthenticationRequests,omitempty"`
	LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool      bool                                    `json:"localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool,omitempty"`
	LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager                interface{}                             `json:"localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager,omitempty"`
	LocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients             string                                  `json:"localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients,omitempty"`
	LocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers             string                                  `json:"localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers,omitempty"`
	LanManagerAuthenticationLevel                                                string                                  `json:"lanManagerAuthenticationLevel,omitempty"`
	LanManagerWorkstationDisableInsecureGuestLogons                              bool                                    `json:"lanManagerWorkstationDisableInsecureGuestLogons,omitempty"`
	LocalSecurityOptionsClearVirtualMemoryPageFile                               bool                                    `json:"localSecurityOptionsClearVirtualMemoryPageFile,omitempty"`
	LocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn              bool                                    `json:"localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn,omitempty"`
	LocalSecurityOptionsAllowUIAccessApplicationElevation                        bool                                    `json:"localSecurityOptionsAllowUIAccessApplicationElevation,omitempty"`
	LocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations bool                                    `json:"localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations,omitempty"`
	LocalSecurityOptionsOnlyElevateSignedExecutables                             bool                                    `json:"localSecurityOptionsOnlyElevateSignedExecutables,omitempty"`
	LocalSecurityOptionsAdministratorElevationPromptBehavior                     string                                  `json:"localSecurityOptionsAdministratorElevationPromptBehavior,omitempty"`
	LocalSecurityOptionsStandardUserElevationPromptBehavior                      string                                  `json:"localSecurityOptionsStandardUserElevationPromptBehavior,omitempty"`
	LocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation           bool                                    `json:"localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation,omitempty"`
	LocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation      bool                                    `json:"localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation,omitempty"`
	LocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations              bool                                    `json:"localSecurityOptionsAllowUIAccessApplicationsForSecureLocations,omitempty"`
	LocalSecurityOptionsUseAdminApprovalMode                                     bool                                    `json:"localSecurityOptionsUseAdminApprovalMode,omitempty"`
	LocalSecurityOptionsUseAdminApprovalModeForAdministrators                    bool                                    `json:"localSecurityOptionsUseAdminApprovalModeForAdministrators,omitempty"`
	LocalSecurityOptionsInformationShownOnLockScreen                             string                                  `json:"localSecurityOptionsInformationShownOnLockScreen,omitempty"`
	LocalSecurityOptionsInformationDisplayedOnLockScreen                         string                                  `json:"localSecurityOptionsInformationDisplayedOnLockScreen,omitempty"`
	LocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees   bool                                    `json:"localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees,omitempty"`
	LocalSecurityOptionsClientDigitallySignCommunicationsAlways                  bool                                    `json:"localSecurityOptionsClientDigitallySignCommunicationsAlways,omitempty"`
	LocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers      bool                                    `json:"localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers,omitempty"`
	LocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways           bool                                    `json:"localSecurityOptionsDisableServerDigitallySignCommunicationsAlways,omitempty"`
	LocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees   bool                                    `json:"localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees,omitempty"`
	LocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares             bool                                    `json:"localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares,omitempty"`
	LocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts              bool                                    `json:"localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts,omitempty"`
	LocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares          bool                                    `json:"localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares,omitempty"`
	LocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange        bool                                    `json:"localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange,omitempty"`
	LocalSecurityOptionsSmartCardRemovalBehavior                                 string                                  `json:"localSecurityOptionsSmartCardRemovalBehavior,omitempty"`
	DefenderSecurityCenterDisableAppBrowserUI                                    bool                                    `json:"defenderSecurityCenterDisableAppBrowserUI,omitempty"`
	DefenderSecurityCenterDisableFamilyUI                                        bool                                    `json:"defenderSecurityCenterDisableFamilyUI,omitempty"`
	DefenderSecurityCenterDisableHealthUI                                        bool                                    `json:"defenderSecurityCenterDisableHealthUI,omitempty"`
	DefenderSecurityCenterDisableNetworkUI                                       bool                                    `json:"defenderSecurityCenterDisableNetworkUI,omitempty"`
	DefenderSecurityCenterDisableVirusUI                                         bool                                    `json:"defenderSecurityCenterDisableVirusUI,omitempty"`
	DefenderSecurityCenterDisableAccountUI                                       bool                                    `json:"defenderSecurityCenterDisableAccountUI,omitempty"`
	DefenderSecurityCenterDisableClearTpmUI                                      bool                                    `json:"defenderSecurityCenterDisableClearTpmUI,omitempty"`
	DefenderSecurityCenterDisableHardwareUI                                      bool                                    `json:"defenderSecurityCenterDisableHardwareUI,omitempty"`
	DefenderSecurityCenterDisableNotificationAreaUI                              bool                                    `json:"defenderSecurityCenterDisableNotificationAreaUI,omitempty"`
	DefenderSecurityCenterDisableRansomwareUI                                    bool                                    `json:"defenderSecurityCenterDisableRansomwareUI,omitempty"`
	DefenderSecurityCenterDisableSecureBootUI                                    interface{}                             `json:"defenderSecurityCenterDisableSecureBootUI,omitempty"`
	DefenderSecurityCenterDisableTroubleshootingUI                               interface{}                             `json:"defenderSecurityCenterDisableTroubleshootingUI,omitempty"`
	DefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI                   bool                                    `json:"defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI,omitempty"`
	DefenderSecurityCenterOrganizationDisplayName                                string                                  `json:"defenderSecurityCenterOrganizationDisplayName,omitempty"`
	DefenderSecurityCenterHelpEmail                                              string                                  `json:"defenderSecurityCenterHelpEmail,omitempty"`
	DefenderSecurityCenterHelpPhone                                              string                                  `json:"defenderSecurityCenterHelpPhone,omitempty"`
	DefenderSecurityCenterHelpURL                                                string                                  `json:"defenderSecurityCenterHelpURL,omitempty"`
	DefenderSecurityCenterNotificationsFromApp                                   string                                  `json:"defenderSecurityCenterNotificationsFromApp,omitempty"`
	DefenderSecurityCenterITContactDisplay                                       string                                  `json:"defenderSecurityCenterITContactDisplay,omitempty"`
	WindowsDefenderTamperProtection                                              string                                  `json:"windowsDefenderTamperProtection,omitempty"`
	FirewallBlockStatefulFTP                                                     bool                                    `json:"firewallBlockStatefulFTP,omitempty"`
	FirewallIdleTimeoutForSecurityAssociationInSeconds                           int                                     `json:"firewallIdleTimeoutForSecurityAssociationInSeconds,omitempty"`
	FirewallPreSharedKeyEncodingMethod                                           string                                  `json:"firewallPreSharedKeyEncodingMethod,omitempty"`
	FirewallIPSecExemptionsNone                                                  bool                                    `json:"firewallIPSecExemptionsNone,omitempty"`
	FirewallIPSecExemptionsAllowNeighborDiscovery                                bool                                    `json:"firewallIPSecExemptionsAllowNeighborDiscovery,omitempty"`
	FirewallIPSecExemptionsAllowICMP                                             bool                                    `json:"firewallIPSecExemptionsAllowICMP,omitempty"`
	FirewallIPSecExemptionsAllowRouterDiscovery                                  bool                                    `json:"firewallIPSecExemptionsAllowRouterDiscovery,omitempty"`
	FirewallIPSecExemptionsAllowDHCP                                             bool                                    `json:"firewallIPSecExemptionsAllowDHCP,omitempty"`
	FirewallCertificateRevocationListCheckMethod                                 string                                  `json:"firewallCertificateRevocationListCheckMethod,omitempty"`
	FirewallMergeKeyingModuleSettings                                            bool                                    `json:"firewallMergeKeyingModuleSettings,omitempty"`
	FirewallPacketQueueingMethod                                                 string                                  `json:"firewallPacketQueueingMethod,omitempty"`
	FirewallProfileDomain                                                        EndpointProtectionSubsetFirewallProfile `json:"firewallProfileDomain,omitempty"`
	FirewallProfilePublic                                                        EndpointProtectionSubsetFirewallProfile `json:"firewallProfilePublic,omitempty"`
	FirewallProfilePrivate                                                       EndpointProtectionSubsetFirewallProfile `json:"firewallProfilePrivate,omitempty"`
	FirewallRules                                                                []EndpointProtectionSubsetFirewallRule  `json:"firewallRules,omitempty"`
	DefenderAdobeReaderLaunchChildProcess                                        string                                  `json:"defenderAdobeReaderLaunchChildProcess,omitempty"`
	DefenderAttackSurfaceReductionExcludedPaths                                  []string                                `json:"defenderAttackSurfaceReductionExcludedPaths,omitempty"`
	DefenderOfficeAppsOtherProcessInjectionType                                  string                                  `json:"defenderOfficeAppsOtherProcessInjectionType,omitempty"`
	DefenderOfficeAppsOtherProcessInjection                                      string                                  `json:"defenderOfficeAppsOtherProcessInjection,omitempty"`
	DefenderOfficeCommunicationAppsLaunchChildProcess                            string                                  `json:"defenderOfficeCommunicationAppsLaunchChildProcess,omitempty"`
	DefenderOfficeAppsExecutableContentCreationOrLaunchType                      string                                  `json:"defenderOfficeAppsExecutableContentCreationOrLaunchType,omitempty"`
	DefenderOfficeAppsExecutableContentCreationOrLaunch                          string                                  `json:"defenderOfficeAppsExecutableContentCreationOrLaunch,omitempty"`
	DefenderOfficeAppsLaunchChildProcessType                                     string                                  `json:"defenderOfficeAppsLaunchChildProcessType,omitempty"`
	DefenderOfficeAppsLaunchChildProcess                                         string                                  `json:"defenderOfficeAppsLaunchChildProcess,omitempty"`
	DefenderOfficeMacroCodeAllowWin32ImportsType                                 string                                  `json:"defenderOfficeMacroCodeAllowWin32ImportsType,omitempty"`
	DefenderOfficeMacroCodeAllowWin32Imports                                     string                                  `json:"defenderOfficeMacroCodeAllowWin32Imports,omitempty"`
	DefenderScriptObfuscatedMacroCodeType                                        string                                  `json:"defenderScriptObfuscatedMacroCodeType,omitempty"`
	DefenderScriptObfuscatedMacroCode                                            string                                  `json:"defenderScriptObfuscatedMacroCode,omitempty"`
	DefenderScriptDownloadedPayloadExecutionType                                 string                                  `json:"defenderScriptDownloadedPayloadExecutionType,omitempty"`
	DefenderScriptDownloadedPayloadExecution                                     string                                  `json:"defenderScriptDownloadedPayloadExecution,omitempty"`
	DefenderPreventCredentialStealingType                                        string                                  `json:"defenderPreventCredentialStealingType,omitempty"`
	DefenderProcessCreationType                                                  string                                  `json:"defenderProcessCreationType,omitempty"`
	DefenderProcessCreation                                                      string                                  `json:"defenderProcessCreation,omitempty"`
	DefenderUntrustedUSBProcessType                                              string                                  `json:"defenderUntrustedUSBProcessType,omitempty"`
	DefenderUntrustedUSBProcess                                                  string                                  `json:"defenderUntrustedUSBProcess,omitempty"`
	DefenderUntrustedExecutableType                                              string                                  `json:"defenderUntrustedExecutableType,omitempty"`
	DefenderUntrustedExecutable                                                  string                                  `json:"defenderUntrustedExecutable,omitempty"`
	DefenderEmailContentExecutionType                                            string                                  `json:"defenderEmailContentExecutionType,omitempty"`
	DefenderEmailContentExecution                                                string                                  `json:"defenderEmailContentExecution,omitempty"`
	DefenderAdvancedRansomewareProtectionType                                    string                                  `json:"defenderAdvancedRansomewareProtectionType,omitempty"`
	DefenderGuardMyFoldersType                                                   string                                  `json:"defenderGuardMyFoldersType,omitempty"`
	DefenderGuardedFoldersAllowedAppPaths                                        []string                                `json:"defenderGuardedFoldersAllowedAppPaths,omitempty"`
	DefenderAdditionalGuardedFolders                                             []string                                `json:"defenderAdditionalGuardedFolders,omitempty"`
	DefenderNetworkProtectionType                                                string                                  `json:"defenderNetworkProtectionType,omitempty"`
	DefenderExploitProtectionXML                                                 interface{}                             `json:"defenderExploitProtectionXml,omitempty"`
	DefenderExploitProtectionXMLFileName                                         interface{}                             `json:"defenderExploitProtectionXmlFileName,omitempty"`
	DefenderSecurityCenterBlockExploitProtectionOverride                         bool                                    `json:"defenderSecurityCenterBlockExploitProtectionOverride,omitempty"`
	DefenderBlockPersistenceThroughWmiType                                       string                                  `json:"defenderBlockPersistenceThroughWmiType,omitempty"`
	AppLockerApplicationControl                                                  string                                  `json:"appLockerApplicationControl,omitempty"`
	DeviceGuardLocalSystemAuthorityCredentialGuardSettings                       string                                  `json:"deviceGuardLocalSystemAuthorityCredentialGuardSettings,omitempty"`
	DeviceGuardEnableVirtualizationBasedSecurity                                 bool                                    `json:"deviceGuardEnableVirtualizationBasedSecurity,omitempty"`
	DeviceGuardEnableSecureBootWithDMA                                           bool                                    `json:"deviceGuardEnableSecureBootWithDMA,omitempty"`
	DeviceGuardSecureBootWithDMA                                                 string                                  `json:"deviceGuardSecureBootWithDMA,omitempty"`
	DeviceGuardLaunchSystemGuard                                                 string                                  `json:"deviceGuardLaunchSystemGuard,omitempty"`
	SmartScreenEnableInShell                                                     bool                                    `json:"smartScreenEnableInShell,omitempty"`
	SmartScreenBlockOverrideForFiles                                             bool                                    `json:"smartScreenBlockOverrideForFiles,omitempty"`
	ApplicationGuardEnabled                                                      bool                                    `json:"applicationGuardEnabled,omitempty"`
	ApplicationGuardEnabledOptions                                               string                                  `json:"applicationGuardEnabledOptions,omitempty"`
	ApplicationGuardBlockFileTransfer                                            string                                  `json:"applicationGuardBlockFileTransfer,omitempty"`
	ApplicationGuardBlockNonEnterpriseContent                                    bool                                    `json:"applicationGuardBlockNonEnterpriseContent,omitempty"`
	ApplicationGuardAllowPersistence                                             bool                                    `json:"applicationGuardAllowPersistence,omitempty"`
	ApplicationGuardForceAuditing                                                bool                                    `json:"applicationGuardForceAuditing,omitempty"`
	ApplicationGuardBlockClipboardSharing                                        string                                  `json:"applicationGuardBlockClipboardSharing,omitempty"`
	ApplicationGuardAllowPrintToPDF                                              bool                                    `json:"applicationGuardAllowPrintToPDF,omitempty"`
	ApplicationGuardAllowPrintToXPS                                              bool                                    `json:"applicationGuardAllowPrintToXPS,omitempty"`
	ApplicationGuardAllowPrintToLocalPrinters                                    bool                                    `json:"applicationGuardAllowPrintToLocalPrinters,omitempty"`
	ApplicationGuardAllowPrintToNetworkPrinters                                  bool                                    `json:"applicationGuardAllowPrintToNetworkPrinters,omitempty"`
	ApplicationGuardAllowVirtualGPU                                              bool                                    `json:"applicationGuardAllowVirtualGPU,omitempty"`
	ApplicationGuardAllowFileSaveOnHost                                          bool                                    `json:"applicationGuardAllowFileSaveOnHost,omitempty"`
	ApplicationGuardAllowCameraMicrophoneRedirection                             interface{}                             `json:"applicationGuardAllowCameraMicrophoneRedirection,omitempty"`
	ApplicationGuardCertificateThumbprints                                       []interface{}                           `json:"applicationGuardCertificateThumbprints,omitempty"`
	BitLockerAllowStandardUserEncryption                                         bool                                    `json:"bitLockerAllowStandardUserEncryption,omitempty"`
	BitLockerDisableWarningForOtherDiskEncryption                                bool                                    `json:"bitLockerDisableWarningForOtherDiskEncryption,omitempty"`
	BitLockerEnableStorageCardEncryptionOnMobile                                 bool                                    `json:"bitLockerEnableStorageCardEncryptionOnMobile,omitempty"`
	BitLockerEncryptDevice                                                       bool                                    `json:"bitLockerEncryptDevice,omitempty"`
	BitLockerRecoveryPasswordRotation                                            string                                  `json:"bitLockerRecoveryPasswordRotation,omitempty"`
	DefenderDisableScanArchiveFiles                                              interface{}                             `json:"defenderDisableScanArchiveFiles,omitempty"`
	DefenderAllowScanArchiveFiles                                                interface{}                             `json:"defenderAllowScanArchiveFiles,omitempty"`
	DefenderDisableBehaviorMonitoring                                            interface{}                             `json:"defenderDisableBehaviorMonitoring,omitempty"`
	DefenderAllowBehaviorMonitoring                                              interface{}                             `json:"defenderAllowBehaviorMonitoring,omitempty"`
	DefenderDisableCloudProtection                                               interface{}                             `json:"defenderDisableCloudProtection,omitempty"`
	DefenderAllowCloudProtection                                                 interface{}                             `json:"defenderAllowCloudProtection,omitempty"`
	DefenderEnableScanIncomingMail                                               interface{}                             `json:"defenderEnableScanIncomingMail,omitempty"`
	DefenderEnableScanMappedNetworkDrivesDuringFullScan                          interface{}                             `json:"defenderEnableScanMappedNetworkDrivesDuringFullScan,omitempty"`
	DefenderDisableScanRemovableDrivesDuringFullScan                             interface{}                             `json:"defenderDisableScanRemovableDrivesDuringFullScan,omitempty"`
	DefenderAllowScanRemovableDrivesDuringFullScan                               interface{}                             `json:"defenderAllowScanRemovableDrivesDuringFullScan,omitempty"`
	DefenderDisableScanDownloads                                                 interface{}                             `json:"defenderDisableScanDownloads,omitempty"`
	DefenderAllowScanDownloads                                                   interface{}                             `json:"defenderAllowScanDownloads,omitempty"`
	DefenderDisableIntrusionPreventionSystem                                     interface{}                             `json:"defenderDisableIntrusionPreventionSystem,omitempty"`
	DefenderAllowIntrusionPreventionSystem                                       interface{}                             `json:"defenderAllowIntrusionPreventionSystem,omitempty"`
	DefenderDisableOnAccessProtection                                            interface{}                             `json:"defenderDisableOnAccessProtection,omitempty"`
	DefenderAllowOnAccessProtection                                              interface{}                             `json:"defenderAllowOnAccessProtection,omitempty"`
	DefenderDisableRealTimeMonitoring                                            interface{}                             `json:"defenderDisableRealTimeMonitoring,omitempty"`
	DefenderAllowRealTimeMonitoring                                              interface{}                             `json:"defenderAllowRealTimeMonitoring,omitempty"`
	DefenderDisableScanNetworkFiles                                              interface{}                             `json:"defenderDisableScanNetworkFiles,omitempty"`
	DefenderAllowScanNetworkFiles                                                interface{}                             `json:"defenderAllowScanNetworkFiles,omitempty"`
	DefenderDisableScanScriptsLoadedInInternetExplorer                           interface{}                             `json:"defenderDisableScanScriptsLoadedInInternetExplorer,omitempty"`
	DefenderAllowScanScriptsLoadedInInternetExplorer                             interface{}                             `json:"defenderAllowScanScriptsLoadedInInternetExplorer,omitempty"`
	//DefenderBlockEndUserAccess                                                   interface{}   `json:"defenderBlockEndUserAccess,omitempty"`
	DefenderAllowEndUserAccess                  interface{} `json:"defenderAllowEndUserAccess,omitempty"`
	DefenderScanMaxCPUPercentage                interface{} `json:"defenderScanMaxCpuPercentage,omitempty"`
	DefenderCheckForSignaturesBeforeRunningScan interface{} `json:"defenderCheckForSignaturesBeforeRunningScan,omitempty"`
	//DefenderCloudBlockLevel                                                      interface{}   `json:"defenderCloudBlockLevel,omitempty"`
	//DefenderCloudExtendedTimeoutInSeconds                                        interface{}   `json:"defenderCloudExtendedTimeoutInSeconds,omitempty"`
	//DefenderDaysBeforeDeletingQuarantinedMalware                                 interface{}   `json:"defenderDaysBeforeDeletingQuarantinedMalware,omitempty"`
	//DefenderDisableCatchupFullScan                                               interface{}   `json:"defenderDisableCatchupFullScan,omitempty"`
	//DefenderDisableCatchupQuickScan                                              interface{}   `json:"defenderDisableCatchupQuickScan,omitempty"`
	DefenderEnableLowCPUPriority interface{} `json:"defenderEnableLowCpuPriority,omitempty"`
	//DefenderFileExtensionsToExclude                                              []interface{} `json:"defenderFileExtensionsToExclude,omitempty"`
	//DefenderFilesAndFoldersToExclude                                             []interface{} `json:"defenderFilesAndFoldersToExclude,omitempty"`
	//DefenderProcessesToExclude                                                   []interface{} `json:"defenderProcessesToExclude,omitempty"`
	//DefenderPotentiallyUnwantedAppAction                                         interface{}   `json:"defenderPotentiallyUnwantedAppAction,omitempty"`
	DefenderScanDirection interface{} `json:"defenderScanDirection,omitempty"`
	//DefenderScanType                                                             interface{}   `json:"defenderScanType,omitempty"`
	//DefenderScheduledQuickScanTime                                               interface{}   `json:"defenderScheduledQuickScanTime,omitempty"`
	DefenderScheduledScanDay interface{} `json:"defenderScheduledScanDay,omitempty"`
	//DefenderScheduledScanTime                                                    interface{}   `json:"defenderScheduledScanTime,omitempty"`
	//DefenderSignatureUpdateIntervalInHours                                       interface{}   `json:"defenderSignatureUpdateIntervalInHours,omitempty"`
	//DefenderSubmitSamplesConsentType                                             interface{}   `json:"defenderSubmitSamplesConsentType,omitempty"`
	//DefenderDetectedMalwareActions                                               interface{}   `json:"defenderDetectedMalwareActions,omitempty"`
	UserRightsAllowAccessFromNetwork        EndpointProtectionSubsetUserRights                    `json:"userRightsAllowAccessFromNetwork,omitempty"`
	UserRightsActAsPartOfTheOperatingSystem EndpointProtectionSubsetUserRights                    `json:"userRightsActAsPartOfTheOperatingSystem,omitempty"`
	UserRightsLocalLogOn                    EndpointProtectionSubsetUserRights                    `json:"userRightsLocalLogOn,omitempty"`
	UserRightsBackupData                    EndpointProtectionSubsetUserRights                    `json:"userRightsBackupData,omitempty"`
	UserRightsChangeSystemTime              EndpointProtectionSubsetUserRights                    `json:"userRightsChangeSystemTime,omitempty"`
	UserRightsCreateGlobalObjects           EndpointProtectionSubsetUserRights                    `json:"userRightsCreateGlobalObjects,omitempty"`
	UserRightsCreatePageFile                EndpointProtectionSubsetUserRights                    `json:"userRightsCreatePageFile,omitempty"`
	UserRightsCreatePermanentSharedObjects  EndpointProtectionSubsetUserRights                    `json:"userRightsCreatePermanentSharedObjects,omitempty"`
	UserRightsCreateSymbolicLinks           EndpointProtectionSubsetUserRights                    `json:"userRightsCreateSymbolicLinks,omitempty"`
	UserRightsCreateToken                   EndpointProtectionSubsetUserRights                    `json:"userRightsCreateToken,omitempty"`
	UserRightsRemoteDesktopServicesLogOn    EndpointProtectionSubsetUserRights                    `json:"userRightsRemoteDesktopServicesLogOn,omitempty"`
	UserRightsDelegation                    EndpointProtectionSubsetUserRights                    `json:"userRightsDelegation,omitempty"`
	UserRightsGenerateSecurityAudits        EndpointProtectionSubsetUserRights                    `json:"userRightsGenerateSecurityAudits,omitempty"`
	UserRightsImpersonateClient             EndpointProtectionSubsetUserRights                    `json:"userRightsImpersonateClient,omitempty"`
	UserRightsIncreaseSchedulingPriority    EndpointProtectionSubsetUserRights                    `json:"userRightsIncreaseSchedulingPriority,omitempty"`
	UserRightsLoadUnloadDrivers             EndpointProtectionSubsetUserRights                    `json:"userRightsLoadUnloadDrivers,omitempty"`
	UserRightsLockMemory                    EndpointProtectionSubsetUserRights                    `json:"userRightsLockMemory,omitempty"`
	UserRightsManageAuditingAndSecurityLogs EndpointProtectionSubsetUserRights                    `json:"userRightsManageAuditingAndSecurityLogs,omitempty"`
	UserRightsManageVolumes                 EndpointProtectionSubsetUserRights                    `json:"userRightsManageVolumes,omitempty"`
	UserRightsModifyFirmwareEnvironment     EndpointProtectionSubsetUserRights                    `json:"userRightsModifyFirmwareEnvironment,omitempty"`
	UserRightsModifyObjectLabels            EndpointProtectionSubsetUserRights                    `json:"userRightsModifyObjectLabels,omitempty"`
	UserRightsProfileSingleProcess          EndpointProtectionSubsetUserRights                    `json:"userRightsProfileSingleProcess,omitempty"`
	UserRightsRemoteShutdown                EndpointProtectionSubsetUserRights                    `json:"userRightsRemoteShutdown,omitempty"`
	UserRightsRestoreData                   EndpointProtectionSubsetUserRights                    `json:"userRightsRestoreData,omitempty"`
	UserRightsTakeOwnership                 EndpointProtectionSubsetUserRights                    `json:"userRightsTakeOwnership,omitempty"`
	BitLockerSystemDrivePolicy              EndpointProtectionSubsetBitLockerSystemDrivePolicy    `json:"bitLockerSystemDrivePolicy,omitempty"`
	BitLockerFixedDrivePolicy               EndpointProtectionSubsetBitLockerFixedDrivePolicy     `json:"bitLockerFixedDrivePolicy,omitempty"`
	BitLockerRemovableDrivePolicy           EndpointProtectionSubsetBitLockerRemovableDrivePolicy `json:"bitLockerRemovableDrivePolicy,omitempty"`
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
	// Fields for Template - Imported Administrative Templates
	// TODO
	// Fields for Template - Kiosk
	KioskBrowserDefaultUrl                 string                                     `json:"kioskBrowserDefaultUrl"`
	KioskBrowserEnableHomeButton           bool                                       `json:"kioskBrowserEnableHomeButton"`
	KioskBrowserEnableNavigationButtons    bool                                       `json:"kioskBrowserEnableNavigationButtons"`
	KioskBrowserEnableEndSessionButton     bool                                       `json:"kioskBrowserEnableEndSessionButton"`
	KioskBrowserRestartOnIdleTimeInMinutes *interface{}                               `json:"kioskBrowserRestartOnIdleTimeInMinutes"`
	KioskBrowserBlockedURLs                []string                                   `json:"kioskBrowserBlockedURLs"`
	KioskBrowserBlockedUrlExceptions       []string                                   `json:"kioskBrowserBlockedUrlExceptions"`
	EdgeKioskEnablePublicBrowsing          bool                                       `json:"edgeKioskEnablePublicBrowsing"`
	KioskProfiles                          []KioskSubsetKioskProfile                  `json:"kioskProfiles"`
	WindowsKioskForceUpdateSchedule        KioskSubsetWindowsKioskForceUpdateSchedule `json:"windowsKioskForceUpdateSchedule"`
	// configuration profile assignments
	AssignmentsODataContext string                                 `json:"assignments@odata.context,omitempty"`
	Assignments             []DeviceConfigurationProfileAssignment `json:"assignments"`
}

// Subsets

// DeliveryOptimizationSubsetBandwidthMode represents the bandwidth mode configuration in Delivery Optimization.
type DeliveryOptimizationSubsetBandwidthMode struct {
	MaximumDownloadBandwidthInKilobytesPerSecond int `json:"maximumDownloadBandwidthInKilobytesPerSecond,omitempty"`
	MaximumUploadBandwidthInKilobytesPerSecond   int `json:"maximumUploadBandwidthInKilobytesPerSecond,omitempty"`
}

// DeliveryOptimizationSubsetMaximumCacheSize represents the maximum cache size configuration in Delivery Optimization.
type DeliveryOptimizationSubsetMaximumCacheSize struct {
	MaximumCacheSizeInGigabytes int `json:"maximumCacheSizeInGigabytes,omitempty"`
}

type DeviceRestrictionsSubsetConfigureTimeZone struct {
	// Define the fields based on the expected properties for ConfigureTimeZone
	// Example (adjust according to actual data model):
	TimeZoneName string `json:"timeZoneName,omitempty"`
}

type DeviceRestrictionsSubsetDefenderDetectedMalwareActions struct {
	LowSeverity      string `json:"lowSeverity,omitempty"`
	ModerateSeverity string `json:"moderateSeverity,omitempty"`
	HighSeverity     string `json:"highSeverity,omitempty"`
	SevereSeverity   string `json:"severeSeverity,omitempty"`
}

type DeviceRestrictionsSubsetNetworkProxyServer struct {
	Address              string   `json:"address,omitempty"`
	Exceptions           []string `json:"exceptions,omitempty"`
	UseForLocalAddresses bool     `json:"useForLocalAddresses,omitempty"`
}

type DeviceRestrictionsSubsetEdgeSearchEngine struct {
	OdataType                        string `json:"@odata.type,omitempty"`
	EdgeSearchEngineOpenSearchXMLURL string `json:"edgeSearchEngineOpenSearchXmlUrl,omitempty"`
}

// DeviceManagementApplicabilityRuleOsEdition represents the OS edition applicability rule.
type DeviceManagementApplicabilityRuleOsEdition struct {
	ODataType      string   `json:"@odata.type,omitempty"`
	OsEditionTypes []string `json:"osEditionTypes,omitempty"`
	Name           string   `json:"name,omitempty"`
	RuleType       string   `json:"ruleType,omitempty"`
}

// DeviceManagementApplicabilityRuleOsVersion represents the OS version applicability rule.
type DeviceManagementApplicabilityRuleOsVersion struct {
	ODataType    string `json:"@odata.type,omitempty"`
	MinOSVersion string `json:"minOSVersion,omitempty"`
	MaxOSVersion string `json:"maxOSVersion,omitempty"`
	Name         string `json:"name,omitempty"`
	RuleType     string `json:"ruleType,omitempty"`
}

// DeviceManagementApplicabilityRuleDeviceMode represents the device mode applicability rule.
type DeviceManagementApplicabilityRuleDeviceMode struct {
	ODataType  string `json:"@odata.type,omitempty"`
	DeviceMode string `json:"deviceMode,omitempty"`
	Name       string `json:"name,omitempty"`
	RuleType   string `json:"ruleType,omitempty"`
}

// Modify the DeviceConfigurationProfileOmaSetting struct to handle additional fields.
type DeviceConfigurationProfileOmaSetting struct {
	ODataType              string      `json:"@odata.type,omitempty"`
	DisplayName            string      `json:"displayName,omitempty"`
	Description            string      `json:"description,omitempty"`
	OmaUri                 string      `json:"omaUri,omitempty"`
	SecretReferenceValueId string      `json:"secretReferenceValueId,omitempty"`
	IsEncrypted            bool        `json:"isEncrypted,omitempty"`
	Value                  interface{} `json:"value,omitempty"`
	IsReadOnly             bool        `json:"isReadOnly,omitempty"`
	FileName               string      `json:"fileName,omitempty"`
}

// EndpointProtectionSubsetFirewallProfile
type EndpointProtectionSubsetFirewallProfile struct {
	FirewallEnabled                                    string `json:"firewallEnabled,omitempty"`
	StealthModeRequired                                bool   `json:"stealthModeRequired,omitempty"`
	StealthModeBlocked                                 bool   `json:"stealthModeBlocked,omitempty"`
	IncomingTrafficRequired                            bool   `json:"incomingTrafficRequired,omitempty"`
	IncomingTrafficBlocked                             bool   `json:"incomingTrafficBlocked,omitempty"`
	UnicastResponsesToMulticastBroadcastsRequired      bool   `json:"unicastResponsesToMulticastBroadcastsRequired,omitempty"`
	UnicastResponsesToMulticastBroadcastsBlocked       bool   `json:"unicastResponsesToMulticastBroadcastsBlocked,omitempty"`
	InboundNotificationsRequired                       bool   `json:"inboundNotificationsRequired,omitempty"`
	InboundNotificationsBlocked                        bool   `json:"inboundNotificationsBlocked,omitempty"`
	AuthorizedApplicationRulesFromGroupPolicyMerged    bool   `json:"authorizedApplicationRulesFromGroupPolicyMerged,omitempty"`
	AuthorizedApplicationRulesFromGroupPolicyNotMerged bool   `json:"authorizedApplicationRulesFromGroupPolicyNotMerged,omitempty"`
	GlobalPortRulesFromGroupPolicyMerged               bool   `json:"globalPortRulesFromGroupPolicyMerged,omitempty"`
	GlobalPortRulesFromGroupPolicyNotMerged            bool   `json:"globalPortRulesFromGroupPolicyNotMerged,omitempty"`
	ConnectionSecurityRulesFromGroupPolicyMerged       bool   `json:"connectionSecurityRulesFromGroupPolicyMerged,omitempty"`
	ConnectionSecurityRulesFromGroupPolicyNotMerged    bool   `json:"connectionSecurityRulesFromGroupPolicyNotMerged,omitempty"`
	OutboundConnectionsRequired                        bool   `json:"outboundConnectionsRequired,omitempty"`
	OutboundConnectionsBlocked                         bool   `json:"outboundConnectionsBlocked,omitempty"`
	InboundConnectionsRequired                         bool   `json:"inboundConnectionsRequired,omitempty"`
	InboundConnectionsBlocked                          bool   `json:"inboundConnectionsBlocked,omitempty"`
	SecuredPacketExemptionAllowed                      bool   `json:"securedPacketExemptionAllowed,omitempty"`
	SecuredPacketExemptionBlocked                      bool   `json:"securedPacketExemptionBlocked,omitempty"`
	PolicyRulesFromGroupPolicyMerged                   bool   `json:"policyRulesFromGroupPolicyMerged,omitempty"`
	PolicyRulesFromGroupPolicyNotMerged                bool   `json:"policyRulesFromGroupPolicyNotMerged,omitempty"`
}

// EndpointProtectionSubsetFirewallRule
type EndpointProtectionSubsetFirewallRule struct {
	DisplayName             string        `json:"displayName,omitempty"`
	Description             string        `json:"description,omitempty"`
	PackageFamilyName       interface{}   `json:"packageFamilyName,omitempty"`
	FilePath                interface{}   `json:"filePath,omitempty"`
	ServiceName             interface{}   `json:"serviceName,omitempty"`
	Protocol                interface{}   `json:"protocol,omitempty"`
	LocalPortRanges         []interface{} `json:"localPortRanges,omitempty"`
	RemotePortRanges        []interface{} `json:"remotePortRanges,omitempty"`
	LocalAddressRanges      []string      `json:"localAddressRanges,omitempty"`
	RemoteAddressRanges     []string      `json:"remoteAddressRanges,omitempty"`
	ProfileTypes            string        `json:"profileTypes,omitempty"`
	Action                  string        `json:"action,omitempty"`
	TrafficDirection        string        `json:"trafficDirection,omitempty"`
	InterfaceTypes          string        `json:"interfaceTypes,omitempty"`
	EdgeTraversal           string        `json:"edgeTraversal,omitempty"`
	LocalUserAuthorizations interface{}   `json:"localUserAuthorizations,omitempty"`
}

// EndpointProtectionSubsetUserRights

type EndpointProtectionSubsetUserRights struct {
	State              string               `json:"state,omitempty"`
	LocalUsersOrGroups []LocalUsersOrGroups `json:"localUsersOrGroups,omitempty"`
}

type LocalUsersOrGroups struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	SecurityIdentifier string `json:"securityIdentifier,omitempty"`
}

// EndpointProtectionSubsetitLockerSystemDrivePolicy represents the policy for BitLocker on the system drive.
type EndpointProtectionSubsetBitLockerSystemDrivePolicy struct {
	EncryptionMethod                         string                   `json:"encryptionMethod,omitempty"`
	StartupAuthenticationRequired            bool                     `json:"startupAuthenticationRequired,omitempty"`
	StartupAuthenticationBlockWithoutTpmChip bool                     `json:"startupAuthenticationBlockWithoutTpmChip,omitempty"`
	StartupAuthenticationTpmUsage            string                   `json:"startupAuthenticationTpmUsage,omitempty"`
	StartupAuthenticationTpmPinUsage         string                   `json:"startupAuthenticationTpmPinUsage,omitempty"`
	StartupAuthenticationTpmKeyUsage         string                   `json:"startupAuthenticationTpmKeyUsage,omitempty"`
	StartupAuthenticationTpmPinAndKeyUsage   string                   `json:"startupAuthenticationTpmPinAndKeyUsage,omitempty"`
	MinimumPinLength                         int                      `json:"minimumPinLength,omitempty"`
	PrebootRecoveryEnableMessageAndURL       bool                     `json:"prebootRecoveryEnableMessageAndUrl,omitempty"`
	PrebootRecoveryMessage                   interface{}              `json:"prebootRecoveryMessage,omitempty"`
	PrebootRecoveryURL                       interface{}              `json:"prebootRecoveryUrl,omitempty"`
	RecoveryOptions                          BitLockerRecoveryOptions `json:"recoveryOptions,omitempty"`
}

// EndpointProtectionSubsetBitLockerFixedDrivePolicy represents the policy for BitLocker on fixed drives.
type EndpointProtectionSubsetBitLockerFixedDrivePolicy struct {
	EncryptionMethod                string                   `json:"encryptionMethod,omitempty"`
	RequireEncryptionForWriteAccess bool                     `json:"requireEncryptionForWriteAccess,omitempty"`
	RecoveryOptions                 BitLockerRecoveryOptions `json:"recoveryOptions,omitempty"`
}

// EndpointProtectionSubsetBitLockerRemovableDrivePolicy represents the policy for BitLocker on removable drives.
type EndpointProtectionSubsetBitLockerRemovableDrivePolicy struct {
	EncryptionMethod                  string `json:"encryptionMethod,omitempty"`
	RequireEncryptionForWriteAccess   bool   `json:"requireEncryptionForWriteAccess,omitempty"`
	BlockCrossOrganizationWriteAccess bool   `json:"blockCrossOrganizationWriteAccess,omitempty"`
}

// BitLockerRecoveryOptions represents the recovery options for BitLocker.
type BitLockerRecoveryOptions struct {
	BlockDataRecoveryAgent                         bool   `json:"blockDataRecoveryAgent,omitempty"`
	RecoveryPasswordUsage                          string `json:"recoveryPasswordUsage,omitempty"`
	RecoveryKeyUsage                               string `json:"recoveryKeyUsage,omitempty"`
	HideRecoveryOptions                            bool   `json:"hideRecoveryOptions,omitempty"`
	EnableRecoveryInformationSaveToStore           bool   `json:"enableRecoveryInformationSaveToStore,omitempty"`
	RecoveryInformationToStore                     string `json:"recoveryInformationToStore,omitempty"`
	EnableBitLockerAfterRecoveryInformationToStore bool   `json:"enableBitLockerAfterRecoveryInformationToStore,omitempty"`
}

// KioskProfile represents the 'kioskProfiles' JSON object.
type KioskSubsetKioskProfile struct {
	ProfileID                 string                     `json:"profileId"`
	ProfileName               string                     `json:"profileName"`
	AppConfiguration          WindowsKioskSingleWin32App `json:"appConfiguration"`
	UserAccountsConfiguration []WindowsKioskAutologon    `json:"userAccountsConfiguration"`
}

// WindowsKioskSingleWin32App represents the 'appConfiguration' JSON object.
type WindowsKioskSingleWin32App struct {
	ODataType string   `json:"@odata.type"`
	Win32App  Win32App `json:"win32App"`
}

// Win32App represents the 'win32App' JSON object.
type Win32App struct {
	StartLayoutTileSize         string       `json:"startLayoutTileSize"`
	Name                        *interface{} `json:"name"`
	AppType                     string       `json:"appType"`
	AutoLaunch                  bool         `json:"autoLaunch"`
	ClassicAppPath              string       `json:"classicAppPath"`
	EdgeNoFirstRun              bool         `json:"edgeNoFirstRun"`
	EdgeKioskIdleTimeoutMinutes *interface{} `json:"edgeKioskIdleTimeoutMinutes"`
	EdgeKioskType               string       `json:"edgeKioskType"`
	EdgeKiosk                   string       `json:"edgeKiosk"`
}

// WindowsKioskAutologon represents the 'userAccountsConfiguration' JSON object.
type WindowsKioskAutologon struct {
	ODataType string `json:"@odata.type"`
}

// KioskSubsetWindowsKioskForceUpdateSchedule represents the 'windowsKioskForceUpdateSchedule' JSON object.
type KioskSubsetWindowsKioskForceUpdateSchedule struct {
	StartDateTime                      time.Time `json:"startDateTime"`
	Recurrence                         string    `json:"recurrence"`
	DayOfWeek                          string    `json:"dayofWeek"`
	DayOfMonth                         int       `json:"dayofMonth"`
	RunImmediatelyIfAfterStartDateTime bool      `json:"runImmediatelyIfAfterStartDateTime"`
}

// DeviceConfigurationProfileAssignment represents an assignment for a Device Configuration Profile.
type DeviceConfigurationProfileAssignment struct {
	ID       string                                     `json:"id,omitempty"`
	Source   string                                     `json:"source,omitempty"`
	SourceId string                                     `json:"sourceId,omitempty"`
	Intent   string                                     `json:"intent,omitempty"`
	Target   DeviceConfigurationProfileAssignmentTarget `json:"target,omitempty"`
}

// DeviceConfigurationProfileAssignmentTarget represents the target of a configuration profile assignment.
type DeviceConfigurationProfileAssignmentTarget struct {
	ODataType                                  string `json:"@odata.type,omitempty"`
	GroupId                                    string `json:"groupId,omitempty"`
	DeviceAndAppManagementAssignmentFilterId   string `json:"deviceAndAppManagementAssignmentFilterId,omitempty"`
	DeviceAndAppManagementAssignmentFilterType string `json:"deviceAndAppManagementAssignmentFilterType,omitempty"`
}

// GetWindowsDeviceConfigurationProfiles retrieves a list of Windows device configuration profiles from Microsoft Graph API.
// Because this is a shared endpoint, an OdataType match is used to filter the response so that only windows configuration
// profiles are returned.
func (c *Client) GetWindowsDeviceConfigurationProfiles() (*ResourceWindowsConfigurationProfileTemplatesList, error) {
	endpoint := uriGraphBetaDeviceManagementWindowsDeviceConfiguration + "?$expand=assignments"

	var responseDeviceConfigurationProfiles ResourceWindowsConfigurationProfileTemplatesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseDeviceConfigurationProfiles)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device configuration profiles", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Filter to include only Windows device profiles
	var windowsProfiles []ResourceWindowsConfigurationProfileTemplate
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
func (c *Client) GetWindowsDeviceConfigurationProfileByID(id string) (*ResourceWindowsConfigurationProfileTemplate, error) {
	endpoint := fmt.Sprintf("%s/%s?$expand=assignments", uriGraphBetaDeviceManagementWindowsDeviceConfiguration, id)

	var responseDeviceConfigurationProfile ResourceWindowsConfigurationProfileTemplate
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
