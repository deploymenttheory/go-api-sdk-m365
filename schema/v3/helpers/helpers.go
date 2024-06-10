package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-m365/schema/v3/models/openapi3"
	"github.com/mitchellh/mapstructure"
)

// DecodeAndLog decodes the map into OpenAPISpec and logs the fields
func DecodeAndLog(rawData map[string]interface{}, spec *openapi3.OpenAPISpec) error {
	err := mapstructure.Decode(rawData, spec)
	if err != nil {
		return fmt.Errorf("failed to decode map into OpenAPISpec: %w", err)
	}

	logTopLevelFields(rawData)
	return nil
}

// logTopLevelFields logs the top-level fields and one level below
func logTopLevelFields(data map[string]interface{}) {
	log.Println("Logging top-level fields:")
	for key, value := range data {
		log.Printf("Top-level field: %s", key)
		if subMap, ok := value.(map[string]interface{}); ok {
			for subKey := range subMap {
				log.Printf("  Sub-field: %s.%s", key, subKey)
			}
		}
	}
}

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
