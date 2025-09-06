# Go JSON Sanitizer

A robust Go application designed to process complex, nested JSON files. It recursively extracts all string values, cleans them of formatting artifacts and escape sequences, and saves the clean text into individual output files.

---

## Features

-   **Recursive Processing**: Traverses any complex and deeply nested JSON structure to find all string values.
-   **Configuration Driven**: All paths and settings are managed via a simple `settings.json` file—no need for command-line arguments.
-   **Text Sanitization**: Cleans text of common escape sequences (`\n`, `\*`) and other artifacts.
-   **Robust Error Handling**: Includes detailed logging to an `app.log` file for easy debugging.
-   **Safe Filename Generation**: Automatically sanitizes JSON keys to create valid filenames, even if they contain illegal characters like `/`.
-   **Modular Codebase**: Organized into logical packages (`config`, `processor`, `fileutils`) for maintainability and scalability.

---

## Project Structure

```
json-sanitizer/
├── config/
│   └── config.go
├── internal/
│   ├── processor/
│   │   └── sanitizer.go
│   └── fileutils/
│       └── fileutils.go
├── logs/
│   └── app.log
├── .gitignore
├── go.mod
├── LICENSE
├── main.go
├── README.md
├── input.json      (Example input)
└── settings.json
```

---

## Setup and Usage

### Prerequisites

-   Go (version 1.18 or newer) installed on your system.
-   Git installed on your system.

### 1. Clone the Repository

```bash
git clone https://github.com/MTawhid7/json-sanitizer.git
cd json-sanitizer
```

### 2. Configure the Application

Open the `settings.json` file and modify the paths to match your environment. You can use relative or absolute paths.

```json
{
  "inputFile": "input.json",
  "outputDir": "output",
  "logLevel": 1
}
```
-   `inputFile`: The path to the source JSON file you want to process.
-   `outputDir`: The path to the folder where the cleaned `.txt` files will be saved.
-   `logLevel`: The verbosity of the logger (1: Info, 2: Debug, 3: Error).

### 3. Run the Program

Execute the following command from the root of the project directory:

```bash
go run main.go
```

The program will read your configuration, process the input file, and generate the cleaned text files in your specified output directory. Check `logs/app.log` for a detailed execution summary or any errors.