package day01

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPart1(t *testing.T) {
	// Create a temporary test input file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_input.txt")

	// Example test input - replace with actual test case from the problem
	testInput := "# Add test input here"
	if err := os.WriteFile(tmpFile, []byte(testInput), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := Part1(tmpFile)
	if err != nil {
		t.Fatalf("Part1 failed: %v", err)
	}

	// Update with expected result from the problem
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	// Create a temporary test input file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_input.txt")

	// Example test input - replace with actual test case from the problem
	testInput := "# Add test input here"
	if err := os.WriteFile(tmpFile, []byte(testInput), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := Part2(tmpFile)
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}

	// Update with expected result from the problem
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
