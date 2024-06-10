package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/extract"
	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/helpers"
	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/models/openapi3"
	"gopkg.in/yaml.v3"
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

	// Unmarshal the YAML data into a map
	var rawData map[string]interface{}
	err = yaml.Unmarshal(data, &rawData)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Use the helper function to decode the map into OpenAPISpec and log the fields
	var openAPISpec openapi3.OpenAPISpec
	err = helpers.DecodeAndLog(rawData, &openAPISpec)
	if err != nil {
		log.Fatalf("Failed to decode and log OpenAPISpec: %v", err)
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

	extractedData, err := extract.ExtractField(data, fieldName, fieldDepth, extractKey, extractValue, extractUniqueFieldsOnly, sortFields)
	if err != nil {
		return nil, fmt.Errorf("failed to extract %s: %w", fieldName, err)
	}

	// Process the extracted data to get unique paths
	var paths []string
	for key := range extractedData {
		log.Printf("Path: %s", key)
		paths = append(paths, key)
	}

	return paths, nil
}
