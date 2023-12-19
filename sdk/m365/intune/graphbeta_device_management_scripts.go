package intune

import (
	"fmt"
	"time"
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

// GetDeviceManagementScripts gets a list of all Intune Device Management Scripts
func (c *Client) GetDeviceManagementScripts() (*ResourceDeviceManagementScriptsList, error) {
	endpoint := urideviceManagementScripts

	var deviceManagementScripts ResourceDeviceManagementScriptsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceManagementScripts)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Sites: %v", err)
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
		return nil, fmt.Errorf("failed to fetch Site by ID: %v", err)
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
		return nil, fmt.Errorf("failed to fetch all scripts: %v", err)
	}

	var scriptID string
	for _, script := range scripts.Value {
		if script.DisplayName == displayName {
			scriptID = script.ID
			break
		}
	}

	if scriptID == "" {
		return nil, fmt.Errorf("no script found with the name: %s", displayName)
	}

	return c.GetDeviceManagementScriptByID(scriptID)
}
