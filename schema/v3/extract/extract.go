package extract

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// ExtractField extracts a specific field from the YAML data based on the provided parameters
// The function uses a map to ensure uniqueness and then converts the map keys to a slice for sorting
func ExtractField(data []byte, fieldName string, fieldDepth int, extractKey bool, extractValue bool, extractUniqueFieldsOnly bool, sortFields bool, delimiter string) ([]string, error) {
	var rawData map[string]interface{}
	err := yaml.Unmarshal(data, &rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode YAML: %w", err)
	}

	fieldData, ok := rawData[fieldName]
	if !ok {
		return nil, fmt.Errorf("%s section not found in the YAML file", fieldName)
	}

	fieldMap := make(map[string]interface{})
	err = mapstructure.Decode(fieldData, &fieldMap)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", fieldName, err)
	}

	// Extract the fields into a map to ensure uniqueness
	extractedFields := extractFromMap(fieldMap, fieldDepth, extractKey, extractValue)

	// Ensure uniqueness if required
	if extractUniqueFieldsOnly {
		extractedFields = getUniqueFields(extractedFields)
	}

	// Convert the map to a slice for sorting
	extractedSlice := mapKeysToSlice(extractedFields)

	// Sort the slice if required
	if sortFields {
		sort.Strings(extractedSlice)
	}

	return extractedSlice, nil
}

// extractFromMap recursively extracts fields from the map based on depth and extraction parameters
// Using a map here helps in ensuring that the keys (paths) are unique
func extractFromMap(data map[string]interface{}, depth int, extractKey bool, extractValue bool) map[string]interface{} {
	result := make(map[string]interface{})
	if depth == 0 {
		for k, v := range data {
			if extractKey && extractValue {
				result[k] = v
			} else if extractKey {
				result[k] = nil
			} else if extractValue {
				result[fmt.Sprintf("%v", v)] = nil
			}
		}
		return result
	}

	for k, v := range data {
		if nestedMap, ok := v.(map[string]interface{}); ok {
			nestedResult := extractFromMap(nestedMap, depth-1, extractKey, extractValue)
			for nk, nv := range nestedResult {
				if extractKey && extractValue {
					result[fmt.Sprintf("%s.%s", k, nk)] = nv
				} else if extractKey {
					result[fmt.Sprintf("%s.%s", k, nk)] = nil
				} else if extractValue {
					result[fmt.Sprintf("%v", nv)] = nil
				}
			}
		}
	}

	return result
}

// mapKeysToSlice converts the keys of a map to a slice
// This step is necessary for sorting the extracted fields
func mapKeysToSlice(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}

// getUniqueFields filters out duplicate fields from the map
// This function ensures that we only keep unique fields
func getUniqueFields(data map[string]interface{}) map[string]interface{} {
	uniqueFields := make(map[string]interface{})
	seen := make(map[string]struct{})

	for k := range data {
		baseField := strings.Split(k, ".")[0]
		if _, ok := seen[baseField]; !ok {
			seen[baseField] = struct{}{}
			uniqueFields[baseField] = nil
		}
	}

	return uniqueFields
}
