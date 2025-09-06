package fileutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadFile reads the content of a file.
func ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	return data, nil
}

// WriteFile writes data to a file, creating the directory if it doesn't exist.
func WriteFile(dir, filename string, data []byte) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	filePath := filepath.Join(dir, filename)
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}