// cloudpc/msgraph_cloudpc.go
// Graph Api - Cloud PC
// Documentation: https://learn.microsoft.com/en-us/graph/api/virtualendpoint-list-cloudpcs?view=graph-rest-1.0&tabs=http
// Intune location:
// API reference: https://learn.microsoft.com/en-us/graph/api/virtualendpoint-list-cloudpcs?view=graph-rest-1.0&tabs=http
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package cloudpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

// ResponseCloudPCList represents a list of Cloud PC entries
type ResponseCloudPCList struct {
	Value []ResourceCloudPC `json:"value"`
}

const (
	uriCloudPC                = "/v1.0/deviceManagement/virtualEndpoint/cloudPCs"
	ODataTypeCloudPC          = "#microsoft.graph.cloudPC"
	ODataTypeCloudPCParameter = ""
)

// ResourceCloudPC represents a single Cloud PC entry
type ResourceCloudPC struct {
	ODataType                string    `json:"@odata.type"`
	AADDeviceID              string    `json:"aadDeviceId"`
	ID                       string    `json:"id"`
	DisplayName              string    `json:"displayName"`
	ImageDisplayName         string    `json:"imageDisplayName"`
	ManagedDeviceID          string    `json:"managedDeviceId"`
	ManagedDeviceName        string    `json:"managedDeviceName"`
	ProvisioningPolicyID     string    `json:"provisioningPolicyId"`
	ProvisioningPolicyName   string    `json:"provisioningPolicyName"`
	OnPremisesConnectionName string    `json:"onPremisesConnectionName"`
	ServicePlanID            string    `json:"servicePlanId"`
	ServicePlanName          string    `json:"servicePlanName"`
	UserPrincipalName        string    `json:"userPrincipalName"`
	LastModifiedDateTime     time.Time `json:"lastModifiedDateTime"`
	GracePeriodEndDateTime   time.Time `json:"gracePeriodEndDateTime"`
	ProvisioningType         string    `json:"provisioningType"`
}

// GetDeviceCategories retrieves a list of Intune Device Categories from Microsoft Graph API.
func (c *Client) ListCloudPCs() (*ResponseCloudPCList, error) {
	endpoint := uriCloudPC

	var response ResponseCloudPCList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetCloudPCByID retrieves a specific Cloud PC by ID
func (c *Client) GetCloudPCByID(cloudPCID string) (*ResourceCloudPC, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCloudPC, cloudPCID)

	var response ResourceCloudPC
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// EndGracePeriodForCloudPC ends the grace period for a specified Cloud PC by ID
func (c *Client) EndGracePeriodForCloudPC(cloudPCID string) error {
	endpoint := fmt.Sprintf("%s/%s/endGracePeriod", uriCloudPC, cloudPCID)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc end grace period", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// RebootCloudPC sends a command to reboot a specified Cloud PC by ID
func (c *Client) RebootCloudPC(cloudPCID string) error {
	endpoint := fmt.Sprintf("%s%s/reboot", uriCloudPC, cloudPCID)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc reboot", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// RenameCloudPC sends a command to rename a specified Cloud PC by ID
func (c *Client) RenameCloudPC(cloudPCID string, newName string) error {
	endpoint := fmt.Sprintf("%s/%s/rename", uriCloudPC, cloudPCID)

	jsonBody, err := json.Marshal(map[string]string{"displayName": newName})
	if err != nil {
		return fmt.Errorf("error marshaling request body: %v", err)
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, bytes.NewBuffer(jsonBody), nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc rename", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// RestoreCloudPC sends a command to restore a specified Cloud PC by ID from a snapshot
func (c *Client) RestoreCloudPC(cloudPCID string, cloudPcSnapshotID string) error {
	endpoint := fmt.Sprintf("%s/%s/restore", uriCloudPC, cloudPCID)

	// Create the JSON body for the POST request
	requestBody := map[string]string{"cloudPcSnapshotId": cloudPcSnapshotID}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error marshaling request body: %v", err)
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, bytes.NewBuffer(jsonBody), nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc restore", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// TroubleshootCloudPC sends a command to troubleshoot a specified Cloud PC by ID
func (c *Client) TroubleshootCloudPC(cloudPCID string) error {
	endpoint := fmt.Sprintf("%s/%s/troubleshoot", uriCloudPC, cloudPCID)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc troubleshoot", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}