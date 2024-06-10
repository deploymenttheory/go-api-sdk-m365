package extract

import (
	"strings"

	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/helpers"
)

// ProcessField removes special characters and capitalizes the first character of each section
func ProcessField(field string) string {
	// Remove special characters and split by dot
	sections := strings.Split(field, ".")
	for i, section := range sections {
		sections[i] = helpers.Capitalize(section)
	}
	// Join sections to form the final string
	return strings.Join(sections, "")
}
