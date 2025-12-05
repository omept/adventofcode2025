package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func LoadGridRuneFormat() [][]rune {
	// Open the file
	file, err := os.Open("paperroll.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	grid := [][]rune{}
	scanner := bufio.NewScanner(file)

	// Increase buffer size because your lines are VERY long
	const maxCapacity = 10 * 1024 * 1024 // 10MB
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		grid = append(grid, runes)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	// Print grid dimensions
	fmt.Printf("Loaded grid: %d rows, %d columns\n", len(grid), len(grid[0]))
	return grid
}
func LoadGridStringFormat() []string {
	// Open the file
	file, err := os.Open("paperroll.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	grid := []string{}
	scanner := bufio.NewScanner(file)

	// Increase buffer size because your lines are VERY long
	const maxCapacity = 10 * 1024 * 1024 // 10MB
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	// Print grid dimensions
	fmt.Printf("Loaded grid: %d rows, %d columns\n", len(grid), len(grid[0]))
	return grid
}
