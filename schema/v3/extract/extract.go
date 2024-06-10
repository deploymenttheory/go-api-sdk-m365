package extract

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type KeyValue struct {
	Key   string
	Value interface{}
}

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
// ExtractField extracts specific fields from YAML data based on the provided parameters.
// Returns a slice of KeyValue structs containing the extracted fields, or an error if extraction fails.
func ExtractField(data []byte, fieldPath string, fieldDepth int, extractKey bool, extractValue bool, extractUniqueFieldsOnly bool, sortFields bool, delimiter string) ([]KeyValue, error) {
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

	// Convert the map to a slice of KeyValue structs
	var extractedSlice []KeyValue
	for k, v := range extractedFields {
		extractedSlice = append(extractedSlice, KeyValue{Key: k, Value: v})
	}

	// Sort the slice if required
	if sortFields {
		sort.Slice(extractedSlice, func(i, j int) bool {
			return extractedSlice[i].Key < extractedSlice[j].Key
		})
	}

	return extractedSlice, nil
}

// traversePath traverses the YAML structure to find the data at the specified path
func traversePath(data map[string]interface{}, path string) (interface{}, error) {
	current := data

	// Split the path into segments by dot
	segments := strings.Split(path, ".")

	// Iterate over each segment
	for i := 0; i < len(segments); i++ {
		segment := segments[i]
		log.Printf("Current path segment: %s", segment)

		// Check if the current segment exists in the current map
		if val, found := current[segment]; found {
			log.Printf("Traversing to: %s", segment)

			// If the value is a nested map, continue traversing
			if nestedMap, isMap := val.(map[string]interface{}); isMap {
				current = nestedMap
			} else {
				// If the value is not a map, return the value
				return val, nil
			}
		} else {
			// If the segment is not found, try combining it with subsequent segments
			combinedSegment := segment
			for j := i + 1; j < len(segments); j++ {
				combinedSegment += "." + segments[j]

				// Check if the combined segment exists
				if val, found := current[combinedSegment]; found {
					log.Printf("Traversing to combined segment: %s", combinedSegment)

					// If the value is a nested map, continue traversing
					if nestedMap, isMap := val.(map[string]interface{}); isMap {
						current = nestedMap
						i = j // Skip ahead to the end of the combined segment
						break
					} else {
						// If the value is not a map, return the value
						return val, nil
					}
				}
			}
			// If no combined segment is found, log and return an error
			if combinedSegment == segment {
				log.Printf("Path not found: %s", segment)
				return nil, fmt.Errorf("%s section not found in the YAML file", segment)
			}
		}
	}
	// Return the final nested map or value
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
