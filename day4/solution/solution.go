package solution

import (
	"strconv"

	"github.com/danielstepan/adventofcode/day4/internal/mas"
	"github.com/danielstepan/adventofcode/day4/internal/xmas"
	"github.com/danielstepan/adventofcode/internal/input"
)

func prepareInput(inputFilePath string) ([][]string, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return [][]string{}, err
	}

	letters := make([][]string, len(content))
	for i, line := range content {
		for _, char := range line {
			letters[i] = append(letters[i], string(char))
		}
	}

	return letters, nil
}

func Part1(inputFilePath string) (string, error) {
	letters, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	totalOccurrences := 0
	for i, line := range letters {
		for j, char := range line {
			if char == "X" {
				totalOccurrences += xmas.FoundXMASCount(letters, i, j)
			}
		}
	}

	return strconv.Itoa(totalOccurrences), nil
}

func Part2(inputFilePath string) (string, error) {
	letters, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	totalOccurrences := 0
	for i, line := range letters {
		for j, char := range line {
			if char == "A" {
				if mas.FoundMAS(letters, i, j) {
					totalOccurrences++
				}
			}
		}
	}

	return strconv.Itoa(totalOccurrences), nil
}
