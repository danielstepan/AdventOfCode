package grid

import (
	"fmt"
	"strconv"

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

func GetIntGrid(inputFilePath string) ([][]int, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	grid := make([][]int, len(content))
	for i, line := range content {
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, fmt.Errorf("could not convert %q to int: %w", string(char), err)
			}
			grid[i] = append(grid[i], num)
		}
	}

	return grid, nil
}

func CopyGrid[T comparable](grid [][]T) [][]T {
	newGrid := make([][]T, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]T, len(row))
		copy(newGrid[i], row)
	}
	return newGrid
}

func PrintGrid[T any](grid [][]T) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func CountElements[T comparable](grid [][]T, element T) int {
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
