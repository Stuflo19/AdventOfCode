package types

import (
	"slices"
)

type Grid struct {
	guard           Guard
	grid            [][]string
	Total           int
	TotalWithoutNew int
}

func findGuard(grid [][]string) Guard {
	posX := -1
	posY := -1

	for index, row := range grid {
		result := slices.Index(row, "^")

		if result != -1 {
			posX = result
			posY = index
			break
		}
	}

	if posX == -1 || posY == -1 {
		panic("Failed to find guard")
	}

	bounds := NewBounds(len(grid[0]), len(grid), 0, 0)

	return NewGuard(NewPosition(posX, posY), 0, bounds)
}

func NewGrid(grid [][]string) Grid {
	guard := findGuard(grid)

	return Grid{
		guard,
		grid,
		0,
		0,
	}
}

func (g *Grid) hasVisited() {
	if g.grid[g.guard.position.y][g.guard.position.x] != "X" {
		g.grid[g.guard.position.y][g.guard.position.x] = "X"
		g.TotalWithoutNew = 0
		g.Total += 1
	} else {
		g.TotalWithoutNew += 1
	}
}

func (g *Grid) moveHorizontal(newX int) bool {
	if newX >= g.guard.bounds.maxX || newX < g.guard.bounds.minX {
		return false
	} else if g.grid[g.guard.position.y][newX] == "#" {
		g.guard.rotate()
		return true
	} else {
		g.guard.step(newX, g.guard.position.y)
		return true
	}
}

func (g *Grid) moveVertical(newY int) bool {
	if newY >= g.guard.bounds.maxY || newY < g.guard.bounds.minY {
		return false
	} else if g.grid[newY][g.guard.position.x] == "#" {
		g.guard.rotate()
		return true
	} else {
		g.guard.step(g.guard.position.x, newY)
		return true
	}
}

func (g *Grid) moveGuard() bool {
	switch g.guard.facing {
	case UP:
		return g.moveVertical(g.guard.position.y - 1)
	case DOWN:
		return g.moveVertical(g.guard.position.y + 1)
	case LEFT:
		return g.moveHorizontal(g.guard.position.x - 1)
	case RIGHT:
		return g.moveHorizontal(g.guard.position.x + 1)
	default:
		return false
	}
}

func (g *Grid) WatchGuard() {
	withinGrid := true

	for withinGrid && g.TotalWithoutNew < 500 {
		g.hasVisited()

		withinGrid = g.moveGuard()
	}
}
