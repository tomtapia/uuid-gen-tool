package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/uuid"
)

const version = "1.0.0"

func printHelp() {
	fmt.Println("UUID Generator Tool v" + version)
	fmt.Println("\nA simple command-line tool to generate UUID v4 (random UUIDs)")
	fmt.Println("\nUsage:")
	fmt.Println("  uuid-gen [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  -h, --help       Show this help message")
	fmt.Println("  -v, --version    Show version information")
	fmt.Println("  -n, --count N    Generate N UUIDs (default: 1)")
	fmt.Println("  -u, --uppercase  Output UUIDs in uppercase")
	fmt.Println("  -s, --simple     Output UUIDs without hyphens")
	fmt.Println("\nExamples:")
	fmt.Println("  uuid-gen                    # Generate one UUID")
	fmt.Println("  uuid-gen -n 5               # Generate 5 UUIDs")
	fmt.Println("  uuid-gen --uppercase        # Generate UUID in uppercase")
	fmt.Println("  uuid-gen --simple           # Generate UUID without hyphens (32 chars)")
	fmt.Println("  uuid-gen -n 3 -u            # Generate 3 uppercase UUIDs")
	fmt.Println("\nRepository: https://github.com/tomtapia/uuid-gen-tool")
}

func printVersion() {
	fmt.Println("uuid-gen version " + version)
}

func main() {
	// Define flags
	help := flag.Bool("help", false, "Show help message")
	helpShort := flag.Bool("h", false, "Show help message")
	versionFlag := flag.Bool("version", false, "Show version")
	versionShort := flag.Bool("v", false, "Show version")
	count := flag.Int("count", 1, "Number of UUIDs to generate")
	countShort := flag.Int("n", 1, "Number of UUIDs to generate")
	uppercase := flag.Bool("uppercase", false, "Output in uppercase")
	uppercaseShort := flag.Bool("u", false, "Output in uppercase")
	simple := flag.Bool("simple", false, "Output without hyphens")
	simpleShort := flag.Bool("s", false, "Output without hyphens")

	flag.Parse()

	// Handle help flag
	if *help || *helpShort {
		printHelp()
		os.Exit(0)
	}

	// Handle version flag
	if *versionFlag || *versionShort {
		printVersion()
		os.Exit(0)
	}

	// Determine count (prioritize short flag if both are set)
	numUUIDs := *count
	if *countShort != 1 {
		numUUIDs = *countShort
	}

	// Validate count
	if numUUIDs < 1 {
		fmt.Fprintf(os.Stderr, "Error: count must be at least 1\n")
		os.Exit(1)
	}

	// Determine flags (short or long)
	isUppercase := *uppercase || *uppercaseShort
	isSimple := *simple || *simpleShort

	// Generate and print UUIDs
	for i := 0; i < numUUIDs; i++ {
		id := uuid.New()
		output := id.String()

		// Apply formatting
		if isSimple {
			output = fmt.Sprintf("%x", id[:])
		}
		if isUppercase {
			output = fmt.Sprintf("%X", id[:])
			if !isSimple {
				// Add hyphens back in uppercase format
				output = fmt.Sprintf("%s-%s-%s-%s-%s",
					output[0:8], output[8:12], output[12:16], output[16:20], output[20:32])
			}
		}

		fmt.Println(output)
	}
}
