// graphbeta_device_management_scripts.go
// Graph Beta Api - Intune: Proactive Remediation
// Documentation: https://learn.microsoft.com/en-us/mem/intune/fundamentals/remediations
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesMenu/~/remediations
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledevices
// Jamf Pro Classic API requires the structs to support an XML data structure.

package intune

import (
	"fmt"
	"time"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const urideviceManagementScripts = "/beta/deviceManagement/deviceManagementScripts"

// Structs for the device management scripts
type ResourceDeviceManagementScriptsList struct {
	OdataContext       string `json:"@odata.context"`
	MicrosoftGraphTips string `json:"@microsoft.graph.tips"`
	Value              []struct {
		EnforceSignatureCheck bool        `json:"enforceSignatureCheck"`
		RunAs32Bit            bool        `json:"runAs32Bit"`
		ID                    string      `json:"id"`
		DisplayName           string      `json:"displayName"`
		Description           string      `json:"description"`
		ScriptContent         interface{} `json:"scriptContent"`
		CreatedDateTime       time.Time   `json:"createdDateTime"`
		LastModifiedDateTime  time.Time   `json:"lastModifiedDateTime"`
		RunAsAccount          string      `json:"runAsAccount"`
		FileName              string      `json:"fileName"`
		RoleScopeTagIds       []string    `json:"roleScopeTagIds"`
	} `json:"value"`
}

type ResourceDeviceManagementScript struct {
	OdataContext          string    `json:"@odata.context"`
	MicrosoftGraphTips    string    `json:"@microsoft.graph.tips"`
	EnforceSignatureCheck bool      `json:"enforceSignatureCheck"`
	RunAs32Bit            bool      `json:"runAs32Bit"`
	ID                    string    `json:"id"`
	DisplayName           string    `json:"displayName"`
	Description           string    `json:"description"`
	ScriptContent         string    `json:"scriptContent"`
	CreatedDateTime       time.Time `json:"createdDateTime"`
	LastModifiedDateTime  time.Time `json:"lastModifiedDateTime"`
	RunAsAccount          string    `json:"runAsAccount"`
	FileName              string    `json:"fileName"`
	RoleScopeTagIds       []string  `json:"roleScopeTagIds"`
}

type ResourceDeviceManagementScriptRequest struct {
	OdataType             string   `json:"@odata.type"`
	DisplayName           string   `json:"displayName"`
	Description           string   `json:"description"`
	ScriptContent         string   `json:"scriptContent"`
	RunAsAccount          string   `json:"runAsAccount"`
	EnforceSignatureCheck bool     `json:"enforceSignatureCheck"`
	FileName              string   `json:"fileName"`
	RoleScopeTagIds       []string `json:"roleScopeTagIds"`
	RunAs32Bit            bool     `json:"runAs32Bit"`
}

// GetDeviceManagementScripts gets a list of all Intune Device Management Scripts
func (c *Client) GetDeviceManagementScripts() (*ResourceDeviceManagementScriptsList, error) {
	endpoint := urideviceManagementScripts

	var deviceManagementScripts ResourceDeviceManagementScriptsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceManagementScripts)

	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management scripts", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceManagementScripts, nil
}

// GetDeviceManagementScriptByID retrieves a Device Management Script by its ID.
func (c *Client) GetDeviceManagementScriptByID(id string) (*ResourceDeviceManagementScript, error) {
	endpoint := fmt.Sprintf("%s/%s", urideviceManagementScripts, id)

	var deviceManagementScript ResourceDeviceManagementScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceManagementScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device management script", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceManagementScript, nil
}

// GetDeviceManagementScriptByName retrieves a device management script by its display name.
func (c *Client) GetDeviceManagementScriptByDisplayName(displayName string) (*ResourceDeviceManagementScript, error) {
	scripts, err := c.GetDeviceManagementScripts()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device management scripts", err)
	}

	var scriptID string
	for _, script := range scripts.Value {
		if script.DisplayName == displayName {
			scriptID = script.ID
			break
		}
	}

	if scriptID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device management script", displayName, err)
	}

	return c.GetDeviceManagementScriptByID(scriptID)
}

// CreateDeviceManagementScript creates a new device management script.
func (c *Client) CreateDeviceManagementScript(request *ResourceDeviceManagementScriptRequest) (*ResourceDeviceManagementScript, error) {
	endpoint := urideviceManagementScripts

	var createdScript ResourceDeviceManagementScript
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &createdScript)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device management script", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdScript, nil
}
