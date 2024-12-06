package walker

import "fmt"

type WalkingDirection int

const (
	Up WalkingDirection = iota
	Right
	Down
	Left
)

type Guard struct {
	PositionX              int
	PositionY              int
	Direction              WalkingDirection
	Steps                  int
	WalksInLoop            bool
	VisitedLoationsWithDir map[string]struct{}
	UniqueVisitedLocations map[string]struct{}
}

func NewGuard(grid [][]string) *Guard {
	x, y := getGuardPosition(grid)
	return &Guard{
		PositionX:              x,
		PositionY:              y,
		Direction:              Up,
		Steps:                  0,
		WalksInLoop:            false,
		VisitedLoationsWithDir: make(map[string]struct{}),
		UniqueVisitedLocations: make(map[string]struct{}),
	}
}

func (g *Guard) Move(grid [][]string) bool {
	switch g.Direction {
	case Up:
		if isOutOfBound(grid, g.PositionX, g.PositionY-1) {
			return true
		}
		if isCollisionDetected(grid, g.PositionX, g.PositionY-1) {
			g.turnRight()
		} else {
			g.moveForward()
		}
	case Right:
		if isOutOfBound(grid, g.PositionX+1, g.PositionY) {
			return true
		}
		if isCollisionDetected(grid, g.PositionX+1, g.PositionY) {
			g.turnRight()
		} else {
			g.moveForward()
		}
	case Down:
		if isOutOfBound(grid, g.PositionX, g.PositionY+1) {
			return true
		}
		if isCollisionDetected(grid, g.PositionX, g.PositionY+1) {
			g.turnRight()
		} else {
			g.moveForward()
		}
	case Left:
		if isOutOfBound(grid, g.PositionX-1, g.PositionY) {
			return true
		}
		if isCollisionDetected(grid, g.PositionX-1, g.PositionY) {
			g.turnRight()
		} else {
			g.moveForward()
		}
	}

	return false
}

func isOutOfBound(grid [][]string, x, y int) bool {
	return x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid)
}

func (g *Guard) moveForward() {
	switch g.Direction {
	case Up:
		g.PositionY--
	case Right:
		g.PositionX++
	case Down:
		g.PositionY++
	case Left:
		g.PositionX--
	}
	g.recordVisit()
	g.Steps++
}

func (g *Guard) recordVisit() {
	positionWithDirection := fmt.Sprintf("%d,%d,%d", g.PositionX, g.PositionY, g.Direction)
	positionOnly := fmt.Sprintf("%d,%d", g.PositionX, g.PositionY)

	if _, exists := g.VisitedLoationsWithDir[positionWithDirection]; exists {
		g.WalksInLoop = true
	}

	g.VisitedLoationsWithDir[positionWithDirection] = struct{}{}
	g.UniqueVisitedLocations[positionOnly] = struct{}{}
}

func (g *Guard) turnRight() {
	g.Direction = (g.Direction + 1) % 4
}

func isCollisionDetected(grid [][]string, x, y int) bool {
	return grid[y][x] == "#"
}

func getGuardPosition(grid [][]string) (int, int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == "^" {
				return x, y
			}
		}
	}
	return 0, 0
}

func isGuardOutOfBounds(grid [][]string, x, y int) bool {
	return x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid)
}
