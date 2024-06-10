package extract

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// ExtractField extracts a specific field from the YAML data based on the provided parameters
func ExtractField(rawData map[string]interface{}, fieldName string, fieldDepth int, extractKey bool, extractValue bool, extractUniqueFieldsOnly bool, sortFields bool) (map[string]interface{}, error) {
	fieldData, ok := rawData[fieldName]
	if !ok {
		return nil, fmt.Errorf("%s section not found in the YAML file", fieldName)
	}

	fieldMap := make(map[string]interface{})
	err := mapstructure.Decode(fieldData, &fieldMap)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", fieldName, err)
	}

	extractedFields := extractFromMap(fieldMap, fieldDepth, extractKey, extractValue)

	if extractUniqueFieldsOnly {
		extractedFields = getUniqueFields(extractedFields)
	}

	if sortFields {
		extractedFields = sortMapKeys(extractedFields)
	}

	return extractedFields, nil
}

// extractFromMap recursively extracts fields from the map based on depth and extraction parameters
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

// getUniqueFields filters out duplicate fields from the map
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

// sortMapKeys sorts the keys of the map and returns a new map with sorted keys
func sortMapKeys(data map[string]interface{}) map[string]interface{} {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedMap := make(map[string]interface{})
	for _, k := range keys {
		sortedMap[k] = data[k]
	}

	return sortedMap
}
