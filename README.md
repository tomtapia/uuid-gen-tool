# UUID Generator Tool

[![Go](https://github.com/tomtapia/uuid-gen-tool/actions/workflows/build.yml/badge.svg)](https://github.com/tomtapia/uuid-gen-tool/actions/workflows/build.yml)
[![Go Version](https://img.shields.io/badge/Go-1.25.3-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomtapia/uuid-gen-tool)](https://goreportcard.com/report/github.com/tomtapia/uuid-gen-tool)

A simple command-line tool written in Go to generate version 4 UUIDs (Universally Unique Identifiers).

## Description

This project generates random universally unique identifiers (UUID) using the UUID v4 standard. Each execution of the program produces a new unique UUID that can be used for database identifiers, session identifiers, unique file names, and more.

## Features

- ✨ UUID v4 generation (random)
- 🚀 Fast and efficient
- 📦 No complex dependencies
- 💻 Simple to use from the command line

## Requirements

- Go 1.25.3 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/tomtapia/uuid-gen-tool.git
cd uuid-gen-tool
```

2. Download dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build -o uuid-gen
```

## Usage

Run the program to generate a new UUID:

```bash
./uuid-gen
```

### Command-line Options

```bash
uuid-gen [options]
```

**Options:**
- `-h, --help` - Show help message with usage instructions
- `-v, --version` - Show version information
- `-n, --count N` - Generate N UUIDs (default: 1)
- `-u, --uppercase` - Output UUIDs in uppercase
- `-s, --simple` - Output UUIDs without hyphens (32 characters)

### Example output:
```
f47ac10b-58cc-4372-a567-0e02b2c3d479
```

### Examples

**Generate a single UUID (default):**
```bash
./uuid-gen
# Output: 550e8400-e29b-41d4-a716-446655440000
```

**Generate multiple UUIDs:**
```bash
./uuid-gen -n 5
# Output: 5 UUIDs, one per line
```

**Generate UUID in uppercase:**
```bash
./uuid-gen --uppercase
# Output: 550E8400-E29B-41D4-A716-446655440000
```

**Generate UUID without hyphens:**
```bash
./uuid-gen --simple
# Output: 550e8400e29b41d4a716446655440000
```

**Combine options:**
```bash
./uuid-gen -n 3 -u -s
# Output: 3 uppercase UUIDs without hyphens
```

**Show help:**
```bash
./uuid-gen --help
```

### Usage in scripts

You can use this tool in shell scripts to generate unique identifiers:

```bash
# Assign the UUID to a variable
UUID=$(./uuid-gen)
echo "Generated UUID: $UUID"

# Create a file with a unique name
touch "file-$UUID.txt"
```

## Dependencies

This project uses the following external dependency:

- [github.com/google/uuid](https://github.com/google/uuid) v1.6.0 - Google's official library for UUID generation

## What is a UUID?

A UUID (Universally Unique Identifier) is a 128-bit number used to uniquely identify information. A UUID v4 is generated randomly and has an extremely low probability of collision.

Standard format: `xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx`

Where:
- `x` is any hexadecimal digit
- The `4` indicates the version (v4)
- `y` is one of: 8, 9, A, or B

## Building for different platforms

### Linux
```bash
GOOS=linux GOARCH=amd64 go build -o uuid-gen-linux
```

### Windows
```bash
GOOS=windows GOARCH=amd64 go build -o uuid-gen.exe
```

### macOS
```bash
GOOS=darwin GOARCH=amd64 go build -o uuid-gen-mac
```

## Development

### Run without building
```bash
go run main.go
```

### Install globally
```bash
go install
```

This will install the binary in your `$GOPATH/bin`, making it accessible from any location.

## Use Cases

- Unique identifiers for database records
- Session identifiers
- Unique file names
- Transaction identifiers
- Unique keys in distributed systems
- Tracking tokens

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributions

Contributions are welcome. Please open an issue or pull request to suggest changes or improvements.

## Author

[tomtapia](https://github.com/tomtapia)

## Useful Links

- [RFC 4122 - UUID Specification](https://tools.ietf.org/html/rfc4122)
- [Google UUID Library Documentation](https://pkg.go.dev/github.com/google/uuid)
