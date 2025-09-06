package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"json-sanitizer/config"
	"json-sanitizer/internal/fileutils"
	"json-sanitizer/internal/processor"
)

// logger will be a global-like variable for our main package
var logger *log.Logger

func main() {
	// --- Configuration ---
	cfg, err := config.LoadConfig("settings.json")
	if err != nil {
		log.Fatalf("FATAL: Could not load configuration. %v", err)
	}

	// --- Logging ---
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("FATAL: Failed to open log file: %v", err)
	}
	defer logFile.Close()
	logger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Application starting...")

	// --- Read Input File ---
	logger.Printf("Reading input file: %s\n", cfg.InputFile)
	inputData, err := fileutils.ReadFile(cfg.InputFile)
	if err != nil {
		logger.Fatalf("FATAL: %v", err)
	}

	// --- Process JSON ---
	var jsonData interface{} // Start with a generic interface to handle any valid JSON
	if err := json.Unmarshal(inputData, &jsonData); err != nil {
		logger.Fatalf("FATAL: Error unmarshaling JSON: %v", err)
	}

	// --- Sanitize and Write Output ---
	logger.Println("Processing JSON and writing to output files...")
	processNode(jsonData, "", cfg.OutputDir) // Start the recursive processing
	logger.Println("Processing complete.")
}

// processNode is a recursive function that traverses the JSON structure.
// It takes the current node (any), the path built so far, and the output directory.
func processNode(node interface{}, currentPath string, outputDir string) {
	// Use a type switch to determine what to do with the current node.
	switch value := node.(type) {
	case map[string]interface{}:
		// NODE IS AN OBJECT: This is the recursive step.
		// Loop through its keys and call this function again for each child.
		for key, childNode := range value {
			// Append the current key to the path for the next level down.
			newPath := currentPath + key + "-"
			processNode(childNode, newPath, outputDir)
		}
	case string:
		// NODE IS A STRING: This is our "base case". We've found text to process.
		// Remove the trailing "-" from the path to create a clean filename.
		fileName := strings.TrimSuffix(currentPath, "-")
		sanitizedFilename := sanitizeFilename(fileName) + ".txt"

		// Sanitize the text content itself.
		cleanedText := processor.SanitizeText(value)

		// Write the final file.
		err := fileutils.WriteFile(outputDir, sanitizedFilename, []byte(cleanedText))
		if err != nil {
			logger.Printf("ERROR: Failed to write file '%s': %v", sanitizedFilename, err)
		} else {
			logger.Printf("Successfully wrote file: %s\n", sanitizedFilename)
		}
	// We can add cases for []interface{} (arrays), numbers, etc., if needed.
	// For now, we simply ignore any node that isn't an object or a string.
	}
}

// sanitizeFilename replaces characters that are illegal in filenames.
func sanitizeFilename(name string) string {
	// Replace slash and other problematic characters with an underscore.
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		" ", "_", // Also replacing spaces for good measure
	)
	return replacer.Replace(name)
}