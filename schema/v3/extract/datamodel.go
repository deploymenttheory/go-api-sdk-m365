package extract

import (
	"strings"

	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/helpers"
)

// StructTypeField removes special characters and capitalizes the first character of each section
func StructTypeField(field string) string {
	// Split the field by dot
	sections := strings.Split(field, ".")

	// Capitalize each section and concatenate them
	var result string
	for _, section := range sections {
		result += helpers.Capitalize(section)
	}

	return result
}
