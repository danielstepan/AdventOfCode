package solution

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielstepan/adventofcode/internal/input"
)

type Robot struct {
	posX, posY           int
	velocityX, velocityY int
}

func NewRobot(posX, posY, velocityX, velocityY int) *Robot {
	return &Robot{
		posX:      posX,
		posY:      posY,
		velocityX: velocityX,
		velocityY: velocityY,
	}
}

func (r *Robot) move(seconds int, modX, modY int) {
	r.posX += (r.velocityX * seconds)
	r.posY += (r.velocityY * seconds)

	r.posX = ((r.posX % modX) + modX) % modX
	r.posY = ((r.posY % modY) + modY) % modY
}

func getRobots(inputFilePath string) ([]*Robot, error) {
	content, err := input.ReadFileLines(inputFilePath)
	if err != nil {
		return nil, err
	}

	robots := make([]*Robot, 0)

	for _, line := range content {
		line = strings.ReplaceAll(line, "p=", "")
		line = strings.ReplaceAll(line, "v=", "")
		parts := strings.Fields(line)

		positions := strings.Split(parts[0], ",")
		velocities := strings.Split(parts[1], ",")

		posX, _ := strconv.Atoi(positions[0])
		posY, _ := strconv.Atoi(positions[1])

		velocityX, _ := strconv.Atoi(velocities[0])
		velocityY, _ := strconv.Atoi(velocities[1])

		robots = append(robots, NewRobot(
			posX, posY,
			velocityX, velocityY,
		))
	}

	return robots, nil

}

func Part1(inputFilePath string, width, height, seconds int) (string, error) {
	robots, err := getRobots(inputFilePath)
	if err != nil {
		return "", err
	}

	for _, robot := range robots {
		robot.move(seconds, width, height)
	}

	upperLeftQuadrant := make([]Robot, 0)
	upperRightQuadrant := make([]Robot, 0)
	lowerLeftQuadrant := make([]Robot, 0)
	lowerRightQuadrant := make([]Robot, 0)

	for _, robot := range robots {
		if robot.posX == width/2 || robot.posY == height/2 {
			continue
		}

		if robot.posX < width/2 && robot.posY < height/2 {
			upperLeftQuadrant = append(upperLeftQuadrant, *robot)
		} else if robot.posX > width/2 && robot.posY < height/2 {
			upperRightQuadrant = append(upperRightQuadrant, *robot)
		} else if robot.posX < width/2 && robot.posY > height/2 {
			lowerLeftQuadrant = append(lowerLeftQuadrant, *robot)
		} else {
			lowerRightQuadrant = append(lowerRightQuadrant, *robot)
		}
	}

	safetyFactor := len(upperLeftQuadrant) *
		len(upperRightQuadrant) *
		len(lowerLeftQuadrant) *
		len(lowerRightQuadrant)

	return strconv.Itoa(safetyFactor), nil
}

func Part2(inputFilePath string, width, height int) (string, error) {
	const seconds = 1

	robots, err := getRobots(inputFilePath)
	if err != nil {
		return "", err
	}

	counter := 0
	for !hasFull4x4Robots(robots, width, height) {
		for _, robot := range robots {
			robot.move(seconds, width, height)

		}
		counter++
	}

	printGridWithRobots(robots, width, height)

	return strconv.Itoa(counter), nil
}

func createGridWithRobots(robots []*Robot, width, height int) [][]rune {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, robot := range robots {
		grid[robot.posY][robot.posX] = 'X'
	}

	return grid
}

func printGridWithRobots(robots []*Robot, width, height int) {
	grid := createGridWithRobots(robots, width, height)

	for i := range grid {
		fmt.Println(string(grid[i]))
	}
}

func hasFull4x4Robots(robots []*Robot, width, height int) bool {
	grid := createGridWithRobots(robots, width, height)

	for y := 0; y <= height-4; y++ {
		for x := 0; x <= width-4; x++ {
			if is4x4Full(grid, x, y) {
				return true
			}
		}
	}

	return false
}

func is4x4Full(grid [][]rune, startX, startY int) bool {
	for y := startY; y < startY+4; y++ {
		for x := startX; x < startX+4; x++ {
			if grid[y][x] != 'X' {
				return false
			}
		}
	}
	return true
}
