// http_client_oauth.go
/* The http_client_auth package focuses on authentication mechanisms for an HTTP client.
It provides structures and methods for handling OAuth-based authentication
*/
package http_client

import "fmt"

// SetGraphAuthenticationMethod interprets and sets the credentials for the HTTP Client.
func (c *Client) SetGraphAuthenticationMethod(creds map[string]string) {
	// Check for OAuth App credentials
	if clientID, ok := creds["clientID"]; ok {
		c.OAuthCredentials.ClientID = clientID

		if clientSecret, ok := creds["clientSecret"]; ok {
			// Client Secret is present, use OAuth App authentication
			c.OAuthCredentials.ClientSecret = clientSecret
			c.AuthMethod = "oauthApp"
		} else if certPath, ok := creds["certificatePath"]; ok {
			// Certificate path is present, use OAuth Certificate authentication
			c.OAuthCredentials.CertificatePath = certPath
			c.AuthMethod = "oauthCertificate"

			// Optionally, load additional certificate details if provided
			if certKeyPath, ok := creds["certificateKeyPath"]; ok {
				c.OAuthCredentials.CertificateKeyPath = certKeyPath
			}
			if thumbprint, ok := creds["certThumbprint"]; ok {
				c.OAuthCredentials.CertThumbprint = thumbprint
			}
		} else {
			// Neither Client Secret nor Certificate Path is provided
			fmt.Errorf("OAuth credentials are incomplete: either client secret or certificate path must be provided")
		}
	} else {
		fmt.Errorf("client ID is required for OAuth authentication")
	}
}
