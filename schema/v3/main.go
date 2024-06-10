package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/extract"
	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/helpers"
)

func main() {
	// Define command-line arguments
	filePath := flag.String("file", "openapi.yaml", "Path to the OpenAPI YAML file")
	exportPath := flag.String("export-path", "exported", "Path to export the Go structs")

	// Parse command-line arguments
	flag.Parse()

	// Read the YAML file
	data, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Ensure the export folder exists
	err = helpers.CreateFolderIfNotExist(*exportPath)
	if err != nil {
		log.Fatalf("Failed to create export folder: %v", err)
	}

	// Print out the fields at various depths
	// err = helpers.PrintFieldsAtDepth(data, 0)
	// if err != nil {
	// 	log.Fatalf("Failed to print fields: %v", err)
	// }
	// err = helpers.PrintFieldsAtDepth(data, 1)
	// if err != nil {
	// 	log.Fatalf("Failed to print fields: %v", err)
	// }
	// // err = printFieldsAtDepth(data, 2)
	// // if err != nil {
	// // 	log.Fatalf("Failed to print fields: %v", err)
	// // }
	// // err = printFieldsAtDepth(data, 3)
	// // if err != nil {
	// // 	log.Fatalf("Failed to print fields: %v", err)
	// // }

	// Extract paths using the helper function
	paths, err := extract.ExtractURLPaths(data)
	if err != nil {
		log.Fatalf("Failed to extract paths: %v", err)
	}

	// Save the paths to a new file called msgraphpaths.go
	pathsFilePath := filepath.Join(*exportPath, "msgraphpaths.go")
	err = extract.SaveURLPathsToFile(paths, pathsFilePath)
	if err != nil {
		log.Fatalf("Failed to save paths to file: %v", err)
	}

	// Extract and print data models
	err = extractAndPrintProperties(data)
	if err != nil {
		log.Fatalf("Failed to extract and print data models: %v", err)
	}

	fmt.Println("Export successful")
}

// extractAndPrintProperties extracts and prints key-value pairs under values within properties
func extractAndPrintProperties(data []byte) error {
	// Define extraction parameters for properties
	fieldPath := "components.examples"
	fieldDepth := 0
	extractKey := true
	extractValue := true
	extractUniqueFieldsOnly := false
	sortFields := false
	delimiter := ""

	// Extract fields using ExtractField function
	extractedData, err := extract.ExtractField(data, fieldPath, fieldDepth, extractKey, extractValue, extractUniqueFieldsOnly, sortFields, delimiter)
	if err != nil {
		return fmt.Errorf("failed to extract %s: %w", fieldPath, err)
	}

	// Process and print the extracted fields
	for _, kv := range extractedData {
		fmt.Println(kv.Key)
		err := extractAndPrintNestedProperties(data, kv.Key)
		if err != nil {
			log.Fatalf("Failed to extract nested properties: %v", err)
		}
	}

	return nil
}

// extractAndPrintNestedProperties extracts and prints key-value pairs under 'value'
func extractAndPrintNestedProperties(data []byte, field string) error {
	fieldPath := fmt.Sprintf("components.examples.%s.value", field)
	fieldDepth := 0
	extractKey := true
	extractValue := true
	extractUniqueFieldsOnly := false
	sortFields := false
	delimiter := ""

	extractedNestedData, err := extract.ExtractField(data, fieldPath, fieldDepth, extractKey, extractValue, extractUniqueFieldsOnly, sortFields, delimiter)
	if err != nil {
		// Log the error and continue instead of returning an error
		log.Printf("No value field found for %s: %v", field, err)
		return nil
	}

	if len(extractedNestedData) == 0 {
		log.Printf("No value field found for %s", field)
		return nil
	}

	// Print the key and its nested values
	fmt.Printf("%s:\n", field)
	for _, kv := range extractedNestedData {
		fmt.Printf("  %s: %v\n", kv.Key, kv.Value)
	}

	return nil
}
