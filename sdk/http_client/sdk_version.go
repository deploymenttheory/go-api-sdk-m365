// sdk_version.go
package http_client

import "fmt"

const (
	SDKVersion    = "0.0.1"
	UserAgentBase = "go-jamfpro-api-intune"
)

func GetUserAgentHeader() string {
	return fmt.Sprintf("%s/%s", UserAgentBase, SDKVersion)
}
