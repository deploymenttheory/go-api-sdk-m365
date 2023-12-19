// http_helpers.go
package http_client

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// ParseISO8601Date attempts to parse a string date in ISO 8601 format.
func ParseISO8601Date(dateStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateStr)
}

// EnsureHTTPScheme prefixes a URL with "http://" it defaults to "https://" doesn't already have an "https://".
func EnsureHTTPScheme(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return fmt.Sprintf("https://%s", url)
	}
	return url
}

// CheckDeprecationHeader checks the response headers for the Deprecation header and logs a warning if present.
func CheckDeprecationHeader(resp *http.Response, logger Logger) {
	deprecationHeader := resp.Header.Get("Deprecation")
	if deprecationHeader != "" {
		logger.Warn("API endpoint is deprecated as of", "Date", deprecationHeader)
	}
}
