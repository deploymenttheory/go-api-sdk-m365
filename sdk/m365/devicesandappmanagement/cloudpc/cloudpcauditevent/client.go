package cloudpcauditevent

import (
	"github.com/deploymenttheory/go-api-http-client/httpclient"
)

// Client struct defines a custom type for handling specific API interactions.
// It embeds an instance of *httpclient.Client which will be used to make HTTP requests
// to the msgraph service.
type Client struct {
	HTTP *httpclient.Client
}

// NewClient is a constructor function that initializes a new Client object for msgraph services.
// It takes a pointer to an httpclient.Client as an argument and returns a pointer to the newly created
// Client instance. This setup allows the use of a single httpclient.Client across multiple services,
// promoting reusability and configurability.
func NewClient(http *httpclient.Client) *Client {
	return &Client{
		HTTP: http,
	}
}
