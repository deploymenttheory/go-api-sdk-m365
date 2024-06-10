package extract

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// ExtractField extracts specific fields from YAML data based on the provided parameters.
// The function supports traversing nested paths within the YAML structure, extracting keys, values,
// or both, and can ensure uniqueness and sorting of the extracted fields.
//
// Parameters:
// - data: The raw YAML data as a byte slice.
// - fieldPath: A dot-separated string representing the path to the nested field (e.g., "components.schemas").
// - fieldDepth: The depth to which fields should be extracted within the nested structure. A depth of 0 means only direct fields are extracted.
// - extractKey: A boolean indicating whether to extract the keys of the fields.
// - extractValue: A boolean indicating whether to extract the values of the fields.
// - extractUniqueFieldsOnly: A boolean indicating whether to ensure the extracted fields are unique.
// - sortFields: A boolean indicating whether to sort the extracted fields alphabetically.
// - delimiter: A string used to join nested field keys (not used in the current implementation).
//
// Returns:
// - A slice of strings containing the extracted fields, or an error if extraction fails.
func ExtractField(data []byte, fieldPath string, fieldDepth int, extractKey bool, extractValue bool, extractUniqueFieldsOnly bool, sortFields bool, delimiter string) ([]string, error) {
	var rawData map[string]interface{}
	err := yaml.Unmarshal(data, &rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode YAML: %w", err)
	}

	fieldData, err := traversePath(rawData, fieldPath)

	if err != nil {
		return nil, err
	}

	fieldMap := make(map[string]interface{})
	err = mapstructure.Decode(fieldData, &fieldMap)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", fieldPath, err)
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

// traversePath traverses the YAML structure to find the data at the specified path
func traversePath(data map[string]interface{}, path string) (interface{}, error) {
	current := data
	regex := regexp.MustCompile(`^microsoft\.graph\.[^.]+`)

	// Function to intelligently split path
	splitPath := func(path string) []string {
		var segments []string
		for {
			if match := regex.FindString(path); match != "" {
				segments = append(segments, match)
				path = strings.TrimPrefix(path, match)
				if len(path) > 0 && path[0] == '.' {
					path = path[1:] // Remove leading dot
				}
			} else {
				segments = append(segments, strings.Split(path, ".")...)
				break
			}
		}
		return segments
	}

	segments := splitPath(path)
	for _, segment := range segments {
		log.Printf("Current path segment: %s", segment)

		if val, found := current[segment]; found {
			log.Printf("Traversing to: %s", segment)
			if nestedMap, isMap := val.(map[string]interface{}); isMap {
				current = nestedMap
			} else {
				return val, nil
			}
		} else {
			log.Printf("Path not found: %s", segment)
			return nil, fmt.Errorf("%s section not found in the YAML file", segment)
		}
	}
	return current, nil
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
		log.Printf("Processing key: %s", k)
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
