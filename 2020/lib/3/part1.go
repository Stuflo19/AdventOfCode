package day3

import (
	"fmt"

	"aoc2020/lib/helpers"
)

const (
	MOVE_X = 3
	MOVE_Y = 1
	TREE   = "#"
)

func Part1(filename string) {
	fileContents := helpers.ReadFile(filename)

	grid := NewGrid(fileContents)

	for grid.sled.position.y < len(grid.forest)-1 {
		grid.slide()
	}

	fmt.Println("Trees hit: ", grid.sled.treesHit)
}
