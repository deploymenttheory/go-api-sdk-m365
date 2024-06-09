package extract

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func ExtractExamples(data []byte) (map[string]interface{}, error) {
	var openAPISpec map[string]interface{}
	err := yaml.Unmarshal(data, &openAPISpec)
	if err != nil {
		return nil, fmt.Errorf("failed to decode YAML: %w", err)
	}

	// Navigate to the examples section in the YAML file
	examples, ok := openAPISpec["examples"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("examples section not found in the YAML file")
	}

	return examples, nil
}
