package shared_export_library

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

// cleanFilename sanitizes the filename by replacing illegal characters with underscores.
func CleanFilename(filename string) string {
	illegalCharacters := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range illegalCharacters {
		filename = strings.ReplaceAll(filename, char, "_")
	}
	return filename
}

// saveOutput saves the configuration data in the specified format.
func SaveOutput(format, path, filename string, data interface{}) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	fullPath := filepath.Join(path, filename)

	switch format {
	case "json":
		return SaveAsJSON(fullPath+".json", data)
	case "yaml":
		return SaveAsYAML(fullPath+".yaml", data)
	default:
		return fmt.Errorf("invalid output format: %s", format)
	}
}

// saveAsJSON saves data in JSON format.
func SaveAsJSON(filePath string, data interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating JSON file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // 4 spaces
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	return nil
}

// saveAsYAML saves data in YAML format.
func SaveAsYAML(filePath string, data interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating YAML file: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("error encoding YAML: %w", err)
	}

	return nil
}

// saveScript decodes base64 script content and saves it to a file.
func SaveScript(basePath, filename, scriptType, base64Content string) error {
	// Decode the base64 content
	decodedContent, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		return fmt.Errorf("error decoding script content: %w", err)
	}

	// Ensure the "Script Data" directory exists
	scriptPath := filepath.Join(basePath, "Script Data")
	if err := os.MkdirAll(scriptPath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating script directory: %w", err)
	}

	// Write the decoded content to a file
	filePath := filepath.Join(scriptPath, fmt.Sprintf("%s_%s.ps1", filename, scriptType))
	if err := os.WriteFile(filePath, decodedContent, 0644); err != nil {
		return fmt.Errorf("error writing script file: %w", err)
	}

	return nil
}

// RemoveKeys removes specific intune data keys from a map.
func RemoveKeys(data map[string]interface{}) map[string]interface{} {
	keysToRemove := map[string]bool{
		"id":                   true,
		"version":              true,
		"topicIdentifier":      true,
		"certificate":          true,
		"createdDateTime":      true,
		"lastModifiedDateTime": true,
		"isDefault":            true,
		"isAssigned":           true,
		"@odata.context":       true,
		"scheduledActionConfigurations@odata.context": true,
		"scheduledActionsForRule@odata.context":       true,
		"sourceId":                                    true,
		"supportsScopeTags":                           true,
		"companyCodes":                                true,
		"isGlobalScript":                              true,
		"highestAvailableVersion":                     true,
		"token":                                       true,
		"lastSyncDateTime":                            true,
		"isReadOnly":                                  true,
		"secretReferenceValueId":                      true,
		"isEncrypted":                                 true,
		"modifiedDateTime":                            true,
		"deployedAppCount":                            true,
	}

	// Iterate over the keys and remove them if they are in the map
	for key := range keysToRemove {
		delete(data, key)
	}

	return data
}

// ConvertStructToMap converts a struct to a map[string]interface{}, removing "omitempty" tags from JSON field tags.
func ConvertStructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := typ.Field(i)

		// Get the field tag value
		tag := typeField.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}

		// Remove "omitempty" from the tag if present
		tag = strings.Split(tag, ",")[0]

		// Add the field to the map
		result[tag] = field.Interface()
	}

	return result
}
