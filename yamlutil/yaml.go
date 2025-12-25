package yamlutil

import (
	"os"

	"gopkg.in/yaml.v3"
)

// WriteFile writes data to a YAML file.
func WriteFile(filePath string, data interface{}) error {
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, yamlBytes, 0664)
}

// ReadFile reads YAML file data into the provided struct.
func ReadFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	return yaml.NewDecoder(file).Decode(v)
}

// IsYAML checks if a string is valid YAML (returns true if valid YAML).
func IsYAML(s string) bool {
	var temp interface{}
	return yaml.Unmarshal([]byte(s), &temp) == nil
}
