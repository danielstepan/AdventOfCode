package solution

import (
	"strconv"

	"github.com/danielstepan/adventofcode/internal/grid"
)

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func Part1(inputFilePath string) (string, error) {
	return run(inputFilePath, calcPerimeterPart1)
}

func Part2(inputFilePath string) (string, error) {
	return run(inputFilePath, calcPerimeterPart2)
}

func run(inputFilePath string, calcPerimeter func(y, x int, garden [][]string) int) (string, error) {
	garden, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	visited := createVisited(len(garden), len(garden[0]))

	result := 0

	for y, row := range garden {
		for x := range row {
			area := 0
			perimeter := 0
			discoverRegion(y, x, garden, visited, &area, &perimeter, calcPerimeter)
			result += area * perimeter
		}
	}
	return strconv.Itoa(result), nil
}

func discoverRegion(y, x int, garden [][]string, visited [][]bool, area, perimeter *int, calcPerimeter func(y, x int, garden [][]string) int) {
	if visited[y][x] {
		return
	}

	visited[y][x] = true

	*area++
	*perimeter += calcPerimeter(y, x, garden)

	for _, dir := range directions {
		newY, newX := y+dir[0], x+dir[1]
		if newY >= 0 && newY < len(garden) && newX >= 0 && newX < len(garden[0]) &&
			garden[newY][newX] == garden[y][x] {
			discoverRegion(newY, newX, garden, visited, area, perimeter, calcPerimeter)
		}
	}
}

func calcPerimeterPart1(y, x int, garden [][]string) int {
	perimeter := 0
	currentChar := garden[y][x]

	for _, dir := range directions {
		newY, newX := y+dir[0], x+dir[1]
		if newY < 0 || newY >= len(garden) || newX < 0 || newX >= len(garden[0]) || garden[newY][newX] != currentChar {
			perimeter++
		}
	}

	return perimeter
}

// counts number of edge lines based on corners
func calcPerimeterPart2(y, x int, garden [][]string) int {
	rows, cols := len(garden), len(garden[0])
	currentChar := garden[y][x]

	isDifferent := func(newY, newX int) bool {
		return newY < 0 || newY >= rows || newX < 0 || newX >= cols || garden[newY][newX] != currentChar
	}

	perimeter := 0

	if isDifferent(y, x-1) && isDifferent(y-1, x) {
		perimeter++
	}

	if isDifferent(y, x+1) && isDifferent(y-1, x) {
		perimeter++
	}

	if isDifferent(y, x-1) && isDifferent(y+1, x) {
		perimeter++
	}

	if isDifferent(y, x+1) && isDifferent(y+1, x) {
		perimeter++
	}

	if !isDifferent(y, x-1) && !isDifferent(y-1, x) && isDifferent(y-1, x-1) {
		perimeter++
	}

	if !isDifferent(y, x+1) && !isDifferent(y-1, x) && isDifferent(y-1, x+1) {
		perimeter++
	}

	if !isDifferent(y, x-1) && !isDifferent(y+1, x) && isDifferent(y+1, x-1) {
		perimeter++
	}

	if !isDifferent(y, x+1) && !isDifferent(y+1, x) && isDifferent(y+1, x+1) {
		perimeter++
	}

	return perimeter
}

func prepareInput(inputFilePath string) ([][]string, error) {
	return grid.GetGrid(inputFilePath)
}

func createVisited(rows, cols int) [][]bool {
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	return visited
}
