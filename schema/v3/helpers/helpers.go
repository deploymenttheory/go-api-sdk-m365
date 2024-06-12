package helpers

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
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

// PrepareNameSafeStructName performs
// Illegal Characters Removal: It uses a regular expression to replace all non-alphanumeric characters with spaces.
// Segment Handling: It splits the sanitized string into segments, capitalizes each segment, and concatenates them.
// Leading Numbers Handling: It prefixes the resulting string with N if it starts with a number.
func PrepareNameSafeStructName(name string) string {
	// Remove illegal characters using a regular expression
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}
	safeName := reg.ReplaceAllString(name, " ")

	// Split by space (previously illegal characters)
	segments := strings.Split(safeName, " ")
	for i, segment := range segments {
		segments[i] = Capitalize(segment)
	}

	// Join the segments
	result := strings.Join(segments, "")

	// Ensure the name does not start with a number
	if len(result) > 0 && result[0] >= '0' && result[0] <= '9' {
		result = "N" + result
	}

	return result
}

// PrepareNameSafeStructFieldName creates a Go-safe field name by capitalizing each segment and preserving underscores.
func PrepareNameSafeStructFieldName(name string) string {
	// Remove illegal characters using a regular expression, except underscores
	reg, err := regexp.Compile("[^a-zA-Z0-9_]+")
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}
	safeName := reg.ReplaceAllString(name, " ")

	// Split by space (previously illegal characters)
	segments := strings.Split(safeName, " ")
	for i, segment := range segments {
		segments[i] = Capitalize(segment)
	}

	// Join the segments
	result := strings.Join(segments, "")

	// Ensure the name does not start with a number
	if len(result) > 0 && unicode.IsDigit(rune(result[0])) {
		result = "N" + result
	}

	return result
}

// Capitalize capitalizes the first character of a string and keeps the rest as is
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// ConvertMSGraphOpenAPITypeToGoType converts OpenAPI types from the MSGraph
// api spec and translates them to go struct types.
// ConvertMSGraphOpenAPITypeToGoType converts OpenAPI types from the MSGraph
// API spec and translates them to Go struct types.
func ConvertMSGraphOpenAPITypeToGoType(openAPIType string) string {
	switch openAPIType {
	case "true", "false":
		return "bool"
	case "0":
		return "int"
	case "String":
		return "string"
	case "0001-01-01T00:00:00.0000000+00:00", "0001-01-01":
		return "time.Time"
	case "00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000 (identifier)":
		return "uuid.UUID" // UUIDs are typically represented as uuid.UUID in Go
	case "Stream":
		return "io.Reader" // this interface is more flexible and idiomatic for handling streams of data than "[]byte".
	case "Duration":
		return "time.Duration"
		// Intervals duration pattern: '^-?P([0-9]+D)?(T([0-9]+H)?([0-9]+M)?([0-9]+([.][0-9]+)?S)?)?$'
		// DateTime duration pattern: '^[0-9]{4,}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]([.][0-9]{1,12})?(Z|[+-][0-9][0-9]:[0-9][0-9])$'
	default:
		if strings.HasPrefix(openAPIType, "microsoft.graph.") {
			return PrepareNameSafeStructName(openAPIType)
		}
		return "interface{}"
	}
}
