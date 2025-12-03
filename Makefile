.PHONY: all build run run-all test clean fmt vet day help

# Default target
all: build

# Build the project
build:
	@echo "Building Advent of Code 2025..."
	@go build -o aoc2025 .

# Run a specific day (usage: make day DAY=1 PART=1)
day:
	@if [ -z "$(DAY)" ]; then \
		echo "Usage: make day DAY=<1-12> [PART=<1|2>]"; \
		exit 1; \
	fi
	@if [ -z "$(PART)" ]; then \
		go run . -day $(DAY); \
	else \
		go run . -day $(DAY) -part $(PART); \
	fi

# Run all days
run-all:
	@go run . -all

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f aoc2025
	@go clean

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

# Help
help:
	@echo "Advent of Code 2025 - Go Edition"
	@echo ""
	@echo "Available targets:"
	@echo "  make build          - Build the project"
	@echo "  make day DAY=N      - Run both parts of day N (1-12)"
	@echo "  make day DAY=N PART=P - Run part P of day N"
	@echo "  make run-all        - Run all days"
	@echo "  make test           - Run all tests"
	@echo "  make fmt            - Format all Go files"
	@echo "  make vet            - Run go vet"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make tidy           - Tidy go.mod dependencies"
	@echo ""
	@echo "Examples:"
	@echo "  make day DAY=1           # Run day 1, both parts"
	@echo "  make day DAY=3 PART=2    # Run day 3, part 2 only"
