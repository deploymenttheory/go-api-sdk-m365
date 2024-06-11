package extract

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/helpers"
)

// StructField represents a field in a Go struct
type StructField struct {
	Name     string
	Type     string
	JSONName string
}

// GoStruct represents a Go struct
type GoStruct struct {
	Name   string
	Fields []StructField
}

func ExtractAndSaveStructs(data []byte, filePath string) error {
	// Define extraction parameters for properties
	fieldPath := "components.examples"
	fieldDepth := 0
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

	// Write the package declaration at the top of the file
	fmt.Fprintln(file, "package extractedmodels")

	// Define the template for the Go structs
	const structTemplate = `
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ .Type }} ` + "`json:\"{{ .JSONName }},omitempty\"`" + `
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

	return nil
}

// extractAndGetStructFields extracts and returns key-value pairs as struct fields
func extractAndGetStructFields(data []byte, field string) ([]StructField, error) {
	fieldPath := fmt.Sprintf("components.examples.%s.value", field)
	fieldDepth := 0
	extractKey := true
	extractValue := true
	extractUniqueFieldsOnly := false
	sortFields := false
	delimiter := ""

	extractedNestedData, err := ExtractField(data, fieldPath, fieldDepth, extractKey, extractValue, extractUniqueFieldsOnly, sortFields, delimiter)
	if err != nil {
		// Log the error and continue instead of returning an error
		log.Printf("No value field found for %s: %v", field, err)
		return nil, nil
	}

	if len(extractedNestedData) == 0 {
		log.Printf("No value field found for %s", field)
		return nil, nil
	}

	var fields []StructField
	for _, kv := range extractedNestedData {
		var fieldType string
		if nestedType, ok := kv.Value.(map[string]interface{}); ok && nestedType["@odata.type"] != nil {
			odataType := nestedType["@odata.type"].(string)
			fieldType = helpers.PrepareNameSafeStructName(odataType)
		} else {
			// Determine the type of the value using the helper function
			fieldType = helpers.ConvertOpenAPITypeToGoType(fmt.Sprintf("%v", kv.Value))
		}
		fieldName := helpers.PrepareNameSafeStructName(kv.Key)
		fields = append(fields, StructField{Name: fieldName, Type: fieldType, JSONName: kv.Key})
	}

	return fields, nil
}
