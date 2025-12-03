package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadLines(t *testing.T) {
	// Create a temporary file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")

	content := "line1\nline2\nline3"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test reading lines
	lines, err := ReadLines(tmpFile)
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}

	expected := []string{"line1", "line2", "line3"}
	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(lines))
	}

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

func TestReadFile(t *testing.T) {
	// Create a temporary file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")

	content := "test content\nwith multiple lines"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test reading file
	data, err := ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if data != content {
		t.Errorf("Expected %q, got %q", content, data)
	}
}

func TestSplitBlocks(t *testing.T) {
	input := "block1\nline2\n\nblock2\nline2\n\nblock3"
	blocks := SplitBlocks(input)

	expected := []string{"block1\nline2", "block2\nline2", "block3"}
	if len(blocks) != len(expected) {
		t.Errorf("Expected %d blocks, got %d", len(expected), len(blocks))
	}

	for i, block := range blocks {
		if block != expected[i] {
			t.Errorf("Block %d: expected %q, got %q", i, expected[i], block)
		}
	}
}
