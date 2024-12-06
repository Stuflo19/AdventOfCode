package part1

import (
	"fmt"

	"stuflo19/aoc2024/day6/types"
	"stuflo19/aoc2024/helpers"
)

func Part1(filename string) {
	puzzleInput := helpers.ReadFileWithSeparator(filename, "")

	grid := types.NewGrid(puzzleInput)

	grid.WatchGuard()

	fmt.Println("Part 1: ", grid.Total)
}
