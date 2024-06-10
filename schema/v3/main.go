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
	paths, err := extractURLPaths(data)
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

// extractURLPaths is a helper function to extract URL paths with the specified parameters
func extractURLPaths(data []byte) ([]string, error) {
	// Define extraction parameters
	fieldName := "paths"
	fieldDepth := 1
	extractKey := true
	extractValue := false
	extractUniqueFieldsOnly := true
	sortFields := true
	delimiter := "/"

	extractedData, err := extract.ExtractField(data, fieldName, fieldDepth, extractKey, extractValue, extractUniqueFieldsOnly, sortFields, delimiter)
	if err != nil {
		return nil, fmt.Errorf("failed to extract %s: %w", fieldName, err)
	}

	return extractedData, nil
}

// extractAndPrintProperties extracts and prints key-value pairs under values within properties
func extractAndPrintProperties(data []byte) error {
	// Define extraction parameters for properties
	fieldPath := "components.examples"
	fieldDepth := 1 // Assuming examples are nested under a level
	extractKey := true
	extractValue := false // Only keys are needed
	extractUniqueFieldsOnly := true
	sortFields := false
	delimiter := ""

	// Extract fields using ExtractField function
	extractedData, err := extract.ExtractField(data, fieldPath, fieldDepth, extractKey, extractValue, extractUniqueFieldsOnly, sortFields, delimiter)
	if err != nil {
		return fmt.Errorf("failed to extract %s: %w", fieldPath, err)
	}

	// Process and print the extracted keys
	for _, field := range extractedData {
		processedField := extract.ProcessField(field)
		fmt.Println(processedField)
	}

	return nil
}
