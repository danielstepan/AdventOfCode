package solution

import (
	"fmt"
	"strconv"

	"github.com/danielstepan/adventofcode/internal/grid"
	"github.com/danielstepan/adventofcode/internal/input"
)

func prepareInput(inputFilePath string) (string, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}

	return content[0], nil
}

func Part1(inputFilePath string) (string, error) {
	topographicMap, _ := grid.GetIntGrid(inputFilePath)

	sum := 0
	for y, row := range topographicMap {
		for x, cell := range row {
			if cell != 0 {
				continue
			}
			visitedPeaks := make(map[string]struct{})
			sequence := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			trails := findTrailsToUniquePeaks(topographicMap, x, y, sequence, visitedPeaks)
			sum += trails
		}
	}

	return strconv.Itoa(sum), nil
}

func findTrailsToUniquePeaks(topographicMap [][]int, x int, y int, sequence []int, visitedPeaks map[string]struct{}) int {
	if x < 0 || y < 0 || x >= len(topographicMap[0]) || y >= len(topographicMap) {
		return 0
	}

	if topographicMap[y][x] != sequence[0] {
		return 0
	}

	if _, exists := visitedPeaks[fmt.Sprintf("%d,%d", x, y)]; exists {
		return 0
	}

	if len(sequence) == 1 {
		visitedPeaks[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
		return 1
	}

	nextSequence := sequence[1:]

	count := findTrailsToUniquePeaks(topographicMap, x+1, y, nextSequence, visitedPeaks) +
		findTrailsToUniquePeaks(topographicMap, x-1, y, nextSequence, visitedPeaks) +
		findTrailsToUniquePeaks(topographicMap, x, y+1, nextSequence, visitedPeaks) +
		findTrailsToUniquePeaks(topographicMap, x, y-1, nextSequence, visitedPeaks)

	return count
}

func Part2(inputFilePath string) (string, error) {
	topographicMap, _ := grid.GetIntGrid(inputFilePath)

	sum := 0
	for y, row := range topographicMap {
		for x, cell := range row {
			if cell != 0 {
				continue
			}
			sequence := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			trails := findAllTrailsToPeaks(topographicMap, x, y, sequence)
			sum += trails
		}
	}

	return strconv.Itoa(sum), nil
}

func findAllTrailsToPeaks(topographicMap [][]int, x int, y int, sequence []int) int {
	if x < 0 || y < 0 || x >= len(topographicMap[0]) || y >= len(topographicMap) {
		return 0
	}

	if topographicMap[y][x] != sequence[0] {
		return 0
	}

	if len(sequence) == 1 {
		return 1
	}

	nextSequence := sequence[1:]

	count := findAllTrailsToPeaks(topographicMap, x+1, y, nextSequence) +
		findAllTrailsToPeaks(topographicMap, x-1, y, nextSequence) +
		findAllTrailsToPeaks(topographicMap, x, y+1, nextSequence) +
		findAllTrailsToPeaks(topographicMap, x, y-1, nextSequence)
	return count
}
