package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/extract"
	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/generate"
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
	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Unmarshal the YAML data into the OpenAPISpec struct
	var openAPISpec openapi3.OpenAPISpec
	err = yaml.Unmarshal(data, &openAPISpec)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Ensure the export folder exists
	err = helpers.CreateFolderIfNotExist(*exportPath)
	if err != nil {
		log.Fatalf("Failed to create export folder: %v", err)
	}

	// Extract and Unmarshal Examples
	examples, err := extract.ExtractExamples(data)
	if err != nil {
		log.Fatalf("Failed to extract examples: %v", err)
	}

	// Generate Go structs from examples
	structs, err := generate.GenerateStructs(examples)
	if err != nil {
		log.Fatalf("Failed to generate structs: %v", err)
	}

	// Save the generated structs to a file
	structFilePath := filepath.Join(*exportPath, "examples_struct.go")
	err = ioutil.WriteFile(structFilePath, []byte(structs), 0644)
	if err != nil {
		log.Fatalf("Failed to save structs to file: %v", err)
	}

	fmt.Println("Export successful")
}
