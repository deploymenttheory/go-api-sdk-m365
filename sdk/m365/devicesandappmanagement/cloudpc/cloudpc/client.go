package cloudpc

import (
	"github.com/deploymenttheory/go-api-http-client/httpclient"
)

// Client wraps the HTTP client for Cloud PC specific API interactions.
type Client struct {
	HTTP *httpclient.Client
}

// NewClient creates a new client for Cloud PC service.
func NewClient(http *httpclient.Client) *Client {
	return &Client{
		HTTP: http,
	}
}
