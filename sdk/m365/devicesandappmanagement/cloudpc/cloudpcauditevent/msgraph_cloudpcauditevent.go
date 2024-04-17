// cloudpc/msgraph_cloudpcauditevent.go
// Graph Api - Cloud PC Audit Event
// Documentation: https://learn.microsoft.com/en-us/graph/api/resources/cloudpcauditevent?view=graph-rest-1.0
// Intune location:
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/cloudpcauditevent?view=graph-rest-1.0
// ODATA query options reference: https://learn.microsoft.com/en-us/graph/query-parameters?tabs=http
// Microsoft Graph requires the structs to support a JSON data structure.

package cloudpcauditevent

import (
	"fmt"
	"time"

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const (
	uriCloudPCAuditEvent                = "/v1.0/deviceManagement/virtualEndpoint/auditEvents"
	ODataTypeCloudPCAuditEvent          = ""
	ODataTypeCloudPCParameterAuditEvent = ""
)

// Response

// ResponseCloudPCAuditEvents represents a list of Cloud PC audit event entries
type ResponseCloudPCAuditEvents struct {
	Value []ResourceAuditEventItem `json:"value"`
}

// ResponseAuditActivityTypes represents a list of audit activity types
type ResponseAuditActivityTypes struct {
	Value []string `json:"value"`
}

// Resource

// ResourceAuditEventItem represents a single audit event for a Cloud PC
type ResourceAuditEventItem struct {
	ODataType             string                          `json:"@odata.type"`
	ID                    string                          `json:"id"`
	DisplayName           string                          `json:"displayName"`
	ComponentName         string                          `json:"componentName"`
	Activity              string                          `json:"activity"`
	ActivityDateTime      time.Time                       `json:"activityDateTime"`
	ActivityType          string                          `json:"activityType"`
	ActivityOperationType string                          `json:"activityOperationType"`
	ActivityResult        string                          `json:"activityResult"`
	CorrelationId         string                          `json:"correlationId"`
	Category              string                          `json:"category"`
	Actor                 AuditEventItemSubsetActor       `json:"actor"`
	Resources             []AuditEventItemSubsetResources `json:"resources"`
}

// AuditEventItemSubsetActor represents the actor of an audit event
type AuditEventItemSubsetActor struct {
	ODataType              string                       `json:"@odata.type"`
	Type                   string                       `json:"type"`
	UserPermissions        []string                     `json:"userPermissions"`
	ApplicationId          string                       `json:"applicationId"`
	ApplicationDisplayName string                       `json:"applicationDisplayName"`
	UserPrincipalName      string                       `json:"userPrincipalName"`
	ServicePrincipalName   string                       `json:"servicePrincipalName"`
	IPAddress              string                       `json:"ipAddress"`
	UserId                 string                       `json:"userId"`
	UserRoleScopeTags      []AuditEventUserRoleScopeTag `json:"userRoleScopeTags"`
	RemoteTenantId         string                       `json:"remoteTenantId"`
	RemoteUserId           string                       `json:"remoteUserId"`
}

// AuditEventUserRoleScopeTag represents information about user role scope tags
type AuditEventUserRoleScopeTag struct {
	ODataType      string `json:"@odata.type"`
	DisplayName    string `json:"displayName"`
	RoleScopeTagId string `json:"roleScopeTagId"`
}

// AuditEventItemSubsetResources represents a resource involved in an audit event
type AuditEventItemSubsetResources struct {
	ODataType          string                       `json:"@odata.type"`
	DisplayName        string                       `json:"displayName"`
	ModifiedProperties []AuditEventModifiedProperty `json:"modifiedProperties"`
	ResourceId         string                       `json:"resourceId"`
}

// AuditEventModifiedProperty represents a property that has been modified in an audit event
type AuditEventModifiedProperty struct {
	ODataType   string `json:"@odata.type"`
	DisplayName string `json:"displayName"`
	OldValue    string `json:"oldValue"`
	NewValue    string `json:"newValue"`
}

// ListCloudPCAuditEvents retrieves a list of Cloud PC audit events from Microsoft Graph API.
func (c *Client) ListCloudPCAuditEvents() (*ResponseCloudPCAuditEvents, error) {
	endpoint := uriCloudPCAuditEvent

	var response ResponseCloudPCAuditEvents
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc audit events", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetCloudPCAuditEventByID retrieves a specific Cloud PC audit event by ID from Microsoft Graph API.
func (c *Client) GetCloudPCAuditEventByID(auditEventID string) (*ResourceAuditEventItem, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCloudPCAuditEvent, auditEventID)

	var response ResourceAuditEventItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "cloud pc audit event", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetAuditActivityTypes retrieves a list of Cloud PC audit activity types from Microsoft Graph API.
func (c *Client) GetAuditActivityTypes() ([]string, error) {
	endpoint := fmt.Sprintf("%s%s", uriCloudPCAuditEvent, "/getAuditActivityTypes")

	var response ResponseAuditActivityTypes
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "audit activity types", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return response.Value, nil
}
