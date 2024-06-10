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

	// Extract paths using the helper function
	paths, err := extractPaths(data)
	if err != nil {
		log.Fatalf("Failed to extract paths: %v", err)
	}

	// Save the paths to a new file called msgraphpaths.go
	pathsFilePath := filepath.Join(*exportPath, "msgraphpaths.go")
	err = extract.SavePathsToFile(paths, pathsFilePath)
	if err != nil {
		log.Fatalf("Failed to save paths to file: %v", err)
	}

	fmt.Println("Export successful")
}

// extractPaths is a helper function to extract paths with the specified parameters
func extractPaths(data []byte) ([]string, error) {
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
