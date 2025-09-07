# Makefile for the Go JSON Sanitizer project

# --- Variables ---
# The name of the final executable binary
BINARY_NAME=json-sanitizer
# The path to the output directory, fetched from the settings file
# This makes it robust even if you change the config.
OUTPUT_DIR := $(shell grep outputDir settings.json | cut -d '"' -f 4)

# --- Main Commands ---

# The .PHONY directive tells make that these are not files,
# but command names. This is a best practice.
.PHONY: run clean build

# Runs the application directly
run:
	@go run main.go

# Builds the application into a single executable binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) main.go
	@echo "$(BINARY_NAME) built successfully."

# Cleans the project by removing the output directory and the binary
clean:
	@echo "Cleaning output directory: $(OUTPUT_DIR)"
	@# Check if the directory exists before trying to remove it
	@if [ -d "$(OUTPUT_DIR)" ]; then \
		rm -rf "$(OUTPUT_DIR)"; \
		echo "Removed $(OUTPUT_DIR)."; \
	else \
		echo "Directory $(OUTPUT_DIR) not found, nothing to clean."; \
	fi
	@echo "Cleaning binary: $(BINARY_NAME)"
	@if [ -f "$(BINARY_NAME)" ]; then \
		rm -f "$(BINARY_NAME)"; \
		echo "Removed $(BINARY_NAME)."; \
	fi