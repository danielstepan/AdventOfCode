package solution

import (
	"strconv"

	"github.com/danielstepan/adventofcode/day6/internal/walker"
	gridUtil "github.com/danielstepan/adventofcode/internal/grid"
)

func prepareInput(inputFilePath string) ([][]string, error) {
	grid, err := gridUtil.GetGrid(inputFilePath)
	if err != nil {
		return nil, err
	}
	return grid, nil
}

func Part1(inputFilePath string) (string, error) {
	grid, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	guard := walker.NewGuard(grid)

	for {
		outOfBound := guard.Move(grid)
		if outOfBound {
			break
		}
	}

	return strconv.Itoa(len(guard.UniqueVisitedLocations)), nil
}

func Part2(inputFilePath string) (string, error) {
	grid, err := prepareInput(inputFilePath)
	if err != nil {
		return "", err
	}

	numOfObstructions := 0

	for y, row := range grid {
		for x := range row {
			obstructedGrid := gridUtil.CopyGrid(grid)
			guard := walker.NewGuard(obstructedGrid)
			if guard.PositionX == x && guard.PositionY == y {
				continue
			}
			obstructedGrid[y][x] = "#"

			for {
				outOfBound := guard.Move(obstructedGrid)
				if outOfBound {
					break
				}
				if guard.WalksInLoop {
					numOfObstructions++
					break
				}
			}
		}
	}

	return strconv.Itoa(numOfObstructions), nil
}
