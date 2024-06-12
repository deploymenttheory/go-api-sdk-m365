package helpers

import (
	"fmt"
	"os"
	"strings"
)

// createFolderIfNotExist creates a folder if it doesn't exist
func CreateFolderIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}

// PrepareNameSafeStructName splits the input name by ".", capitalizes each segment, and concatenates them.
// This is useful for creating struct names from field names in a schema.
func PrepareNameSafeStructName(name string) string {
	segments := strings.Split(name, ".")
	for i, segment := range segments {
		segments[i] = Capitalize(segment)
	}
	return strings.Join(segments, "")
}

// Capitalize capitalizes the first character of a string and keeps the rest as is
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// ConvertOpenAPITypeToGoType converts OpenAPI types to Go struct field types.
func ConvertOpenAPITypeToGoType(openAPIType string) string {
	switch openAPIType {
	case "true", "false":
		return "bool"
	case "0":
		return "int"
	case "String":
		return "string"
	case "0001-01-01T00:00:00.0000000+00:00", "0001-01-01":
		return "time.Time"
	case "00000000-0000-0000-0000-000000000000":
		return "string" // UUIDs are typically represented as strings in Go
	default:
		if strings.HasPrefix(openAPIType, "microsoft.graph.") {
			return PrepareNameSafeStructName(openAPIType)
		}
		return "interface{}"
	}
}
