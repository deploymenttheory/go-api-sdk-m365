// sdk/msgraphclient/msgraph_clients.go
package msgraphclient

import (
	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/devicesandappmanagement/cloudpc/cloudpc"
	"github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/devicesandappmanagement/cloudpc/cloudpcauditevent"
)

// Microsoft Graph API clients, each tailored for specific parts of the Cloud PC service.
// This structure is designed to act as a unified entry point for managing interactions
// with different Microsoft Graph endpoints, facilitating easy and organized access to
// multiple services through a single instance.
type Client struct {
	HTTP *httpclient.Client // HTTP is a generic client used to make HTTP requests. This client is configured
	// with base settings such as authentication headers, base URLs, and other
	// necessary configurations that are common across all the requests made to the
	// Microsoft Graph API.

	CloudPC *cloudpc.Client // CloudPC provides a specialized client for interacting with Cloud PC related
	// services within Microsoft Graph. This client handles operations such as
	// provisioning, managing, and monitoring Cloud PCs, leveraging the HTTP client
	// for actual communication.

	CloudPCAudit *cloudpcauditevent.Client // CloudPCAudit offers a dedicated client for accessing Cloud PC audit event
	// services. This client is used to fetch audit logs, monitor events, and
	// perform security compliance checks specific to Cloud PC operations, again
	// using the underlying HTTP client for network interactions.
}

// NewClient initializes and returns a Client with all dependencies injected. This function serves as a factory
// method that creates a new instance of the Client struct, fully configured with all necessary sub-clients for
// interacting with various aspects of the Microsoft Graph API related to Cloud PC services.
//
// Parameters:
//
//	http *httpclient.Client - A pre-configured instance of an HTTP client that handles the lower-level HTTP
//	                          communications. This client should already be set up with authentication configurations,
//	                          such as tokens or other necessary headers, and any other global settings that should
//	                          be applied to all outgoing HTTP requests.
//
// Returns:
//
//	*Client - A pointer to the newly created Client instance, which includes:
//
//	- HTTP: 					A shared HTTP client used for all network interactions.
//	- CloudPC: 			A client dedicated to handling Cloud PC service interactions, such as managing virtual desktops,
//	           			provisioning, and lifecycle operations.
//	- CloudPCAudit: 	A client focused on accessing and managing Cloud PC audit event logs, which are crucial for
//	           			monitoring and compliance purposes.
//
// Usage:
// This function is typically called at the start of an application or when a new set of services is needed. It centralizes
// the creation and configuration of service-specific clients, ensuring that all components use a consistent HTTP client
// setup. This architecture helps maintain clean separation of concerns and promotes reuse of common configurations and
// connections.
func NewClient(http *httpclient.Client) *Client {
	cloudPCClient := cloudpc.NewClient(http)
	cloudPCAuditClient := cloudpcauditevent.NewClient(http)
	return &Client{
		HTTP:         http,
		CloudPC:      cloudPCClient,
		CloudPCAudit: cloudPCAuditClient,
	}
}
