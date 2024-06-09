package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

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

// saveToFile saves the unmarshalled data to a file
func SaveToFile(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to encode data to file: %w", err)
	}

	return nil
}
