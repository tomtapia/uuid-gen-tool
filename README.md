# UUID Generator Tool

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

### Example output:
```
f47ac10b-58cc-4372-a567-0e02b2c3d479
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

This project is available under the license you deem appropriate.

## Contributions

Contributions are welcome. Please open an issue or pull request to suggest changes or improvements.

## Author

[tomtapia](https://github.com/tomtapia)

## Useful Links

- [RFC 4122 - UUID Specification](https://tools.ietf.org/html/rfc4122)
- [Google UUID Library Documentation](https://pkg.go.dev/github.com/google/uuid)
