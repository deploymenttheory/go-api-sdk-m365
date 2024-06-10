package helpers

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// PrintFieldsAtDepth prints the fields of the YAML data at the specified depth
func PrintFieldsAtDepth(data []byte, depth int) error {
	var rawData map[string]interface{}
	err := yaml.Unmarshal(data, &rawData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	fmt.Printf("Fields at depth %d:\n", depth)
	printMapFields(rawData, depth, 0)
	return nil
}

// printMapFields prints the fields of a map at a specified depth
func printMapFields(m map[string]interface{}, targetDepth, currentDepth int) {
	if targetDepth == currentDepth {
		for key := range m {
			fmt.Println(key)
		}
		return
	}

	for key, value := range m {
		if subMap, ok := value.(map[string]interface{}); ok {
			fmt.Printf("%s:\n", key)
			printMapFields(subMap, targetDepth, currentDepth+1)
		}
	}
}
