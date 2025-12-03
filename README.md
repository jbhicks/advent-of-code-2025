# Advent of Code 2025 - Go Edition

Doing it Go style with as little AI help as possible

## Project Structure

```
advent-of-code-2025/
├── main.go           # Entry point with CLI
├── go.mod            # Go module definition
├── Makefile          # Build and run commands
├── utils/            # Shared utility functions
│   └── input.go      # Input file reading utilities
├── day01/            # Day 1 solutions
│   ├── part1.go
│   ├── part2.go
│   └── input.txt
├── day02/            # Day 2 solutions
│   ├── part1.go
│   ├── part2.go
│   └── input.txt
├── ...
└── day12/            # Day 12 solutions
    ├── part1.go
    ├── part2.go
    └── input.txt
```

## Getting Started

### Prerequisites

- Go 1.20 or higher

### Building the Project

```bash
# Build the executable
make build

# Or manually
go build -o aoc2025 .
```

## Running Solutions

### Using Make (Recommended)

```bash
# Run a specific day (both parts)
make day DAY=1

# Run a specific day and part
make day DAY=1 PART=1

# Run all days
make run-all

# Format code
make fmt

# Run tests
make test

# See all available commands
make help
```

### Using Go Directly

```bash
# Run a specific day
go run . -day 1

# Run a specific part
go run . -day 1 -part 1

# Run all days
go run . -all
```

### Using the Built Executable

```bash
./aoc2025 -day 1
./aoc2025 -day 1 -part 2
./aoc2025 -all
```

## Adding Your Input

For each day, paste your puzzle input into the corresponding `dayXX/input.txt` file.

## Implementing Solutions

Each day has two files: `part1.go` and `part2.go`. Implement your solution in the respective functions:

```go
func Part1(inputFile string) (int, error) {
    lines, err := utils.ReadLines(inputFile)
    if err != nil {
        return 0, err
    }
    
    // Your solution here
    
    return result, nil
}
```

## Utility Functions

The `utils` package provides helpful functions:

- `ReadLines(filename string)` - Read file as slice of strings
- `ReadFile(filename string)` - Read entire file as string
- `SplitBlocks(input string)` - Split input by double newlines

## Development

```bash
# Format code
make fmt

# Run linter
make vet

# Clean build artifacts
make clean

# Tidy dependencies
make tidy
```
