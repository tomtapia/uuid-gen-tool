package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	// Generate a new UUID V4 (Random)
	id := uuid.New()

	// Print the UUID as a string to the console
	fmt.Println(id.String())
}
