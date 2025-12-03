package utils

import (
	"bufio"
	"os"
	"strings"
)

// ReadLines reads a file and returns a slice of strings, one per line
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadFile reads the entire file as a string
func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// SplitBlocks splits input by double newlines
func SplitBlocks(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n\n")
}
