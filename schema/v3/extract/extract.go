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
	extractedFields := extractFromMap(fieldMap, fieldDepth, extractKey, extractValue, data)

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
func extractFromMap(data map[string]interface{}, depth int, extractKey bool, extractValue bool, fullData []byte) map[string]interface{} {
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
		switch nestedData := v.(type) {
		case map[string]interface{}:
			if properties, ok := nestedData["properties"]; ok {
				nestedResult := extractFromMap(properties.(map[string]interface{}), depth-1, extractKey, extractValue, fullData)
				for nk, nv := range nestedResult {
					if extractKey && extractValue {
						result[fmt.Sprintf("%s.%s", k, nk)] = nv
					} else if extractKey {
						result[fmt.Sprintf("%s.%s", k, nk)] = nil
					} else if extractValue {
						result[fmt.Sprintf("%v", nv)] = nil
					}
				}
			} else if allOf, ok := nestedData["allOf"]; ok {
				handleAllOfField(result, k, allOf.([]interface{}), fullData)
			} else {
				handlePrimitiveField(result, k, nestedData)
			}
		case []interface{}:
			handleArrayField(result, k, nestedData)
		case string, int, bool, float64:
			handlePrimitiveField(result, k, nestedData)
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

// handleAllOfField processes allOf fields
func handleAllOfField(result map[string]interface{}, key string, allOfData []interface{}, fullData []byte) {
	log.Printf("Processing allOf field: %s", key)
	for _, item := range allOfData {
		if ref, ok := item.(map[string]interface{})["$ref"]; ok {
			refType := parseRef(ref.(string))
			refPath := fmt.Sprintf("components.schemas.%s", refType)
			refData, err := ExtractField(fullData, refPath, 1, true, true, false, false, "")
			if err != nil {
				log.Printf("Failed to extract $ref field: %v", err)
			}
			for _, refKv := range refData {
				result[fmt.Sprintf("%s.%s", key, refKv.Key)] = refKv.Value
			}
		}
		if props, ok := item.(map[string]interface{})["properties"]; ok {
			nestedResult := extractFromMap(props.(map[string]interface{}), 1, true, true, fullData)
			for nk, nv := range nestedResult {
				result[fmt.Sprintf("%s.%s", key, nk)] = nv
			}
		}
	}
}

// handleArrayField processes array fields
func handleArrayField(result map[string]interface{}, key string, arrayData []interface{}) {
	log.Printf("Processing array field: %s", key)
	for i, item := range arrayData {
		itemKey := fmt.Sprintf("%s[%d]", key, i)
		if nestedMap, ok := item.(map[string]interface{}); ok {
			nestedResult := extractFromMap(nestedMap, 1, true, true, nil)
			for nk, nv := range nestedResult {
				result[fmt.Sprintf("%s.%s", itemKey, nk)] = nv
			}
		} else {
			result[itemKey] = item
		}
	}
}

// handlePrimitiveField processes primitive fields
func handlePrimitiveField(result map[string]interface{}, key string, value interface{}) {
	log.Printf("Processing primitive field: %s", key)
	result[key] = value
}

// mapValues maps OpenAPI types and properties to Go types and struct fields
func mapValues(fieldMap map[string]interface{}) map[string]interface{} {
	mapped := make(map[string]interface{})

	for k, v := range fieldMap {
		switch v := v.(type) {
		case map[string]interface{}:
			if t, ok := v["type"]; ok {
				switch t {
				case "string":
					mapped[k] = "string"
				case "integer":
					mapped[k] = "int"
				case "boolean":
					mapped[k] = "bool"
				case "array":
					if items, ok := v["items"]; ok {
						if ref, ok := items.(map[string]interface{})["$ref"]; ok {
							refType := parseRef(ref.(string))
							mapped[k] = fmt.Sprintf("[]%s", refType)
						} else if itemType, ok := items.(map[string]interface{})["type"]; ok {
							mapped[k] = fmt.Sprintf("[]%s", itemType)
						}
					}
				case "object":
					mapped[k] = "map[string]interface{}"
				}
			} else if ref, ok := v["$ref"]; ok {
				mapped[k] = parseRef(ref.(string))
			} else {
				mapped[k] = "interface{}"
			}
		case string:
			mapped[k] = v
		case int:
			mapped[k] = v
		case bool:
			mapped[k] = v
		default:
			mapped[k] = "interface{}"
		}
	}

	return mapped
}

func parseRef(ref string) string {
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}
