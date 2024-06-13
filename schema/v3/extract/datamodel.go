package extract

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/helpers"
)

// StructField represents a field in a Go struct
type StructField struct {
	Name     string
	Type     string
	JSONName string
	Comment  string
}

// GoStruct represents a Go struct
type GoStruct struct {
	Name   string
	Fields []StructField
}

func ExtractAndSaveStructs(data []byte, filePath string) error {
	startTime := time.Now()

	// Extraction parameters for properties
	fieldPath := "components.schemas"
	fieldDepth := 1
	extractKey := true
	extractValue := true
	extractUniqueFieldsOnly := false
	sortFields := false
	delimiter := ""

	// Extract fields using ExtractField function
	extractedData, err := ExtractField(data, fieldPath, fieldDepth, extractKey, extractValue, extractUniqueFieldsOnly, sortFields, delimiter)
	if err != nil {
		return fmt.Errorf("failed to extract %s: %w", fieldPath, err)
	}

	// Create and open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write the package declaration and constants at the top of the file
	header := `package extractedmodels

import (
	"time"
	"io"
	"github.com/google/uuid"
)

`
	fmt.Fprintln(file, header)

	// Define the template for the Go structs
	const structTemplate = `
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ .Type }} ` + "`json:\"{{ .JSONName }},omitempty\"`" + ` // {{ .Comment }}
{{- end }}
}
`
	// Parse the template
	tmpl, err := template.New("struct").Parse(structTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Process the extracted fields and write each struct as it is extracted
	for _, kv := range extractedData {
		fields, err := extractAndGetStructFields(data, kv.Key)
		if err != nil {
			log.Fatalf("Failed to extract nested properties: %v", err)
		}

		structName := helpers.PrepareNameSafeStructName(kv.Key)
		goStruct := GoStruct{Name: structName, Fields: fields}

		// Execute the template with the current struct and write to the file
		err = tmpl.Execute(file, goStruct)
		if err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	}

	elapsedTime := time.Since(startTime)
	log.Printf("Struct generation completed in %s", elapsedTime)

	return nil
}

// extractAndGetStructFields extracts and returns key-value pairs as struct fields
func extractAndGetStructFields(data []byte, field string) ([]StructField, error) {
	fieldPath := fmt.Sprintf("components.schemas.%s.properties", field)
	fieldDepth := 1
	extractKey := true
	extractValue := true
	extractUniqueFieldsOnly := false
	sortFields := false
	delimiter := ""

	extractedNestedData, err := ExtractField(data, fieldPath, fieldDepth, extractKey, extractValue, extractUniqueFieldsOnly, sortFields, delimiter)
	if err != nil {
		// Log the error and continue instead of returning an error
		log.Printf("No properties field found for %s: %v", field, err)
		return nil, nil
	}

	if len(extractedNestedData) == 0 {
		log.Printf("No properties field found for %s", field)
		return nil, nil
	}

	var fields []StructField
	for _, kv := range extractedNestedData {
		var fieldType string
		fieldName := helpers.PrepareNameSafeStructFieldName(kv.Key)
		fieldComment := ""

		propertiesMap, ok := kv.Value.(map[string]interface{})
		if !ok {
			log.Printf("Expected a map for properties, got: %v", kv.Value)
			continue
		}

		fieldType = helpers.ConvertMSGraphOpenAPITypeToGoType(fmt.Sprintf("%v", propertiesMap["type"]))
		if desc, found := propertiesMap["description"]; found {
			fieldComment = fmt.Sprintf("%v", desc)
		}

		// Handle $ref to handle nested types
		if ref, found := propertiesMap["$ref"]; found {
			refType := ref.(string)
			refParts := strings.Split(refType, "/")
			fieldType = helpers.PrepareNameSafeStructName(refParts[len(refParts)-1])
		}

		fields = append(fields, StructField{Name: fieldName, Type: fieldType, JSONName: kv.Key, Comment: fieldComment})
	}

	return fields, nil
}
