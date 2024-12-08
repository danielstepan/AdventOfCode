package solution

import (
	"fmt"
	"strconv"

	gridInput "github.com/danielstepan/adventofcode/internal/grid"
)

var antiNodes map[string]struct{}

func Part1(inputFilePath string) (string, error) {
	return run(inputFilePath, placeSingleAntiNode)
}

func Part2(inputFilePath string) (string, error) {
	return run(inputFilePath, placeAntiNodesInLine)
}

func run(inputFilePath string, placeAntiNodeFunc func(grid [][]string, x1, y1, x2, y2 int)) (string, error) {
	grid, err := gridInput.GetGrid(inputFilePath)
	if err != nil {
		return "", fmt.Errorf("could not get grid: %w", err)
	}

	antiNodes = make(map[string]struct{})

	for y1, row := range grid {
		for x1, cell := range row {
			if cell == "." || cell == "#" {
				continue
			}
			antenna1 := grid[y1][x1]
			for y2, row := range grid {
				for x2, cell := range row {
					if cell == "." || cell == "#" {
						continue
					}
					antenna2 := grid[y2][x2]
					if antenna1 == antenna2 && !(x1 == x2 && y1 == y2) {
						placeAntiNodeFunc(grid, x1, y1, x2, y2)
					}
				}
			}
		}
	}
	return strconv.Itoa(len(antiNodes)), nil
}

func placeSingleAntiNode(grid [][]string, x1, y1, x2, y2 int) {
	stepX := x2 - x1
	stepY := y2 - y1

	newX, newY := x2+stepX, y2+stepY
	if newX < 0 || newX >= len(grid[0]) || newY < 0 || newY >= len(grid) {
		return
	}
	markAntiNode(grid, newX, newY)
}

func placeAntiNodesInLine(grid [][]string, x1, y1, x2, y2 int) {
	stepX := x2 - x1
	stepY := y2 - y1

	antiNodes[fmt.Sprintf("%d,%d", x1, y1)] = struct{}{}
	antiNodes[fmt.Sprintf("%d,%d", x2, y2)] = struct{}{}

	curX, curY := x2, y2
	for {
		curX += stepX
		curY += stepY
		if curX < 0 || curX >= len(grid[0]) || curY < 0 || curY >= len(grid) {
			break
		}
		markAntiNode(grid, curX, curY)
	}

	curX, curY = x1, y1
	for {
		curX -= stepX
		curY -= stepY
		if curX < 0 || curX >= len(grid[0]) || curY < 0 || curY >= len(grid) {
			break
		}
		markAntiNode(grid, curX, curY)
	}
}

func markAntiNode(grid [][]string, x, y int) {
	antiNodes[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
}
