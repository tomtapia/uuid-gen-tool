package main

import (
	"testing"

	"github.com/google/uuid"
)

// TestUUIDGeneration tests that UUID generation works correctly
func TestUUIDGeneration(t *testing.T) {
	// Generate a UUID
	id := uuid.New()

	// Test that the UUID is not nil
	if id == uuid.Nil {
		t.Error("Generated UUID should not be nil")
	}

	// Test that the UUID string is not empty
	uuidStr := id.String()
	if uuidStr == "" {
		t.Error("UUID string should not be empty")
	}

	// Test that the UUID string has the correct length (36 characters including hyphens)
	if len(uuidStr) != 36 {
		t.Errorf("UUID string length should be 36, got %d", len(uuidStr))
	}

	// Test that the UUID is version 4
	if id.Version() != 4 {
		t.Errorf("UUID should be version 4, got version %d", id.Version())
	}

	// Test that the UUID variant is RFC4122
	if id.Variant() != uuid.RFC4122 {
		t.Errorf("UUID variant should be RFC4122, got %v", id.Variant())
	}
}

// TestUUIDUniqueness tests that generated UUIDs are unique
func TestUUIDUniqueness(t *testing.T) {
	// Generate multiple UUIDs
	uuids := make(map[string]bool)
	iterations := 1000

	for i := 0; i < iterations; i++ {
		id := uuid.New()
		uuidStr := id.String()

		// Check if UUID already exists (collision)
		if uuids[uuidStr] {
			t.Errorf("Duplicate UUID generated: %s", uuidStr)
		}
		uuids[uuidStr] = true
	}

	// Verify we generated the expected number of unique UUIDs
	if len(uuids) != iterations {
		t.Errorf("Expected %d unique UUIDs, got %d", iterations, len(uuids))
	}
}

// TestUUIDFormat tests that the UUID format is correct
func TestUUIDFormat(t *testing.T) {
	id := uuid.New()
	uuidStr := id.String()

	// Check format: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
	// Position 14 should be '4' (version 4)
	if uuidStr[14] != '4' {
		t.Errorf("Character at position 14 should be '4', got '%c'", uuidStr[14])
	}

	// Hyphens should be at positions 8, 13, 18, 23
	hyphenPositions := []int{8, 13, 18, 23}
	for _, pos := range hyphenPositions {
		if uuidStr[pos] != '-' {
			t.Errorf("Character at position %d should be '-', got '%c'", pos, uuidStr[pos])
		}
	}
}

// BenchmarkUUIDGeneration benchmarks UUID generation performance
func BenchmarkUUIDGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uuid.New()
	}
}
