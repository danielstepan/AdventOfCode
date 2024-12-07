package grid

import (
	"fmt"

	"github.com/danielstepan/adventofcode/internal/input"
)

func GetGrid(inputFilePath string) ([][]string, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	grid := make([][]string, len(content))
	for i, line := range content {
		for _, char := range line {
			grid[i] = append(grid[i], string(char))
		}
	}

	return grid, nil
}

func CopyGrid(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]string, len(row))
		copy(newGrid[i], row)
	}
	return newGrid
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func CountElemens(grid [][]string, element string) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == element {
				count++
			}
		}
	}
	return count
}
