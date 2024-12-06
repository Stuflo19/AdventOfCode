package part2

import (
	"fmt"

	"stuflo19/aoc2024/day6/types"
	"stuflo19/aoc2024/helpers"
)

func Part2(filename string) {
	total := 0
	puzzleInput := helpers.ReadFileWithSeparator(filename, "")

	for y := range puzzleInput {
		for x := range puzzleInput {
			if puzzleInput[y][x] == "^" || puzzleInput[y][x] == "#" {
				continue
			}

			copyOfPuzzleInput := make([][]string, len(puzzleInput), len(puzzleInput[0]))

			for i, row := range puzzleInput {
				newRow := make([]string, len(row))
				copy(newRow, row)
				copyOfPuzzleInput[i] = newRow
			}

			copyOfPuzzleInput[y][x] = "#"

			grid := types.NewGrid(copyOfPuzzleInput)

			grid.WatchGuard()

			if grid.TotalWithoutNew == 500 {
				total += 1
			}
		}
	}

	fmt.Println("Part 2: ", total)
}
