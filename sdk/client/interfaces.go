package client

import (
	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/devicesandappmanagement/cloudpc/cloudpc"
)

// Client holds references to different service clients.
type Client struct {
	HTTP    *httpclient.Client
	CloudPC *cloudpc.Client
}

// NewClient initializes and returns a Client with all dependencies injected.
func NewClient(http *httpclient.Client) *Client {
	cloudPCClient := cloudpc.NewClient(http)
	return &Client{
		HTTP:    http,
		CloudPC: cloudPCClient,
	}
}

// Build functions retained as they are...
