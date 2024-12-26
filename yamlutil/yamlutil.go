package yamlutil

import (
	"gopkg.in/yaml.v3"
	"os"
)

// WriteFile writes data to a YAML file.
func WriteFile(filePath string, data any) error {
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, yamlBytes, 0664)
}

// ReadFile reads YAML file data into the provided struct.
func ReadFile(filePath string, v any) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return yaml.NewDecoder(file).Decode(v)
}

// IsYAML checks if a string is valid YAML (returns true if valid YAML).
func IsYAML(s string) bool {
	var temp any
	return yaml.Unmarshal([]byte(s), &temp) == nil
}
