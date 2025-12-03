package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jbhicks/advent-of-code-2025/day01"
	"github.com/jbhicks/advent-of-code-2025/day02"
	"github.com/jbhicks/advent-of-code-2025/day03"
	"github.com/jbhicks/advent-of-code-2025/day04"
	"github.com/jbhicks/advent-of-code-2025/day05"
	"github.com/jbhicks/advent-of-code-2025/day06"
	"github.com/jbhicks/advent-of-code-2025/day07"
	"github.com/jbhicks/advent-of-code-2025/day08"
	"github.com/jbhicks/advent-of-code-2025/day09"
	"github.com/jbhicks/advent-of-code-2025/day10"
	"github.com/jbhicks/advent-of-code-2025/day11"
	"github.com/jbhicks/advent-of-code-2025/day12"
)

type DaySolution struct {
	Part1 func(string) (int, error)
	Part2 func(string) (int, error)
}

func main() {
	dayFlag := flag.Int("day", 0, "Run specific day (1-12)")
	partFlag := flag.Int("part", 0, "Run specific part (1 or 2)")
	allFlag := flag.Bool("all", false, "Run all days")
	flag.Parse()

	solutions := map[int]DaySolution{
		1:  {Part1: day01.Part1, Part2: day01.Part2},
		2:  {Part1: day02.Part1, Part2: day02.Part2},
		3:  {Part1: day03.Part1, Part2: day03.Part2},
		4:  {Part1: day04.Part1, Part2: day04.Part2},
		5:  {Part1: day05.Part1, Part2: day05.Part2},
		6:  {Part1: day06.Part1, Part2: day06.Part2},
		7:  {Part1: day07.Part1, Part2: day07.Part2},
		8:  {Part1: day08.Part1, Part2: day08.Part2},
		9:  {Part1: day09.Part1, Part2: day09.Part2},
		10: {Part1: day10.Part1, Part2: day10.Part2},
		11: {Part1: day11.Part1, Part2: day11.Part2},
		12: {Part1: day12.Part1, Part2: day12.Part2},
	}

	if *allFlag {
		runAll(solutions)
		return
	}

	if *dayFlag == 0 {
		fmt.Println("Usage: go run main.go -day <1-12> [-part <1|2>]")
		fmt.Println("       go run main.go -all")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *dayFlag < 1 || *dayFlag > 12 {
		fmt.Println("Day must be between 1 and 12")
		os.Exit(1)
	}

	solution, exists := solutions[*dayFlag]
	if !exists {
		fmt.Printf("Day %d not implemented\n", *dayFlag)
		os.Exit(1)
	}

	inputFile := getInputFile(*dayFlag)

	if *partFlag == 1 || *partFlag == 0 {
		runPart(*dayFlag, 1, solution.Part1, inputFile)
	}

	if *partFlag == 2 || *partFlag == 0 {
		runPart(*dayFlag, 2, solution.Part2, inputFile)
	}
}

func runAll(solutions map[int]DaySolution) {
	fmt.Println("Running all days...")
	fmt.Println(strings.Repeat("=", 60))

	for day := 1; day <= 12; day++ {
		solution := solutions[day]
		inputFile := getInputFile(day)

		fmt.Printf("\n📅 Day %d\n", day)
		fmt.Println(strings.Repeat("-", 60))

		runPart(day, 1, solution.Part1, inputFile)
		runPart(day, 2, solution.Part2, inputFile)
	}

	fmt.Println(strings.Repeat("=", 60))
}

func runPart(day, part int, fn func(string) (int, error), inputFile string) {
	start := time.Now()
	result, err := fn(inputFile)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("⚠️  Day %d, Part %d: Error - %v\n", day, part, err)
		return
	}

	fmt.Printf("✨ Day %d, Part %d: %d (took %v)\n", day, part, result, elapsed)
}

func getInputFile(day int) string {
	return filepath.Join(fmt.Sprintf("day%02d", day), "input.txt")
}
