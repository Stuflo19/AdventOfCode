package part1

import (
	"fmt"

	"stuflo19/aoc2024/day4"
	"stuflo19/aoc2024/helpers"
)

func Part1(filename string) {
	total := 0
	wordSearch := helpers.ReadFileWithSeparator(filename, "")

	bounds := day4.NewBounds(len(wordSearch[0]), -1, len(wordSearch), -1)

	for lineNo, line := range wordSearch {
		for charNo, char := range wordSearch[lineNo] {
			if char == "X" {
				// Check forwards
				if charNo+3 < bounds.MaxX {
					result := line[charNo] + line[charNo+1] + line[charNo+2] + line[charNo+3]

					if result == "XMAS" {
						total += 1
					}
				}

				// Check backwards
				if charNo-3 > bounds.MinX {
					result := line[charNo] + line[charNo-1] + line[charNo-2] + line[charNo-3]

					if result == "XMAS" {
						total += 1
					}
				}

				// Check up
				if lineNo-3 > bounds.MinY {
					result := line[charNo] + wordSearch[lineNo-1][charNo] + wordSearch[lineNo-2][charNo] + wordSearch[lineNo-3][charNo]

					if result == "XMAS" {
						total += 1
					}
				}

				// Check down
				if lineNo+3 < bounds.MaxY {
					result := line[charNo] + wordSearch[lineNo+1][charNo] + wordSearch[lineNo+2][charNo] + wordSearch[lineNo+3][charNo]

					if result == "XMAS" {
						total += 1
					}
				}

				// Check diagonal up right
				if lineNo-3 > bounds.MinY && charNo+3 < bounds.MaxX {
					result := line[charNo] + wordSearch[lineNo-1][charNo+1] + wordSearch[lineNo-2][charNo+2] + wordSearch[lineNo-3][charNo+3]

					if result == "XMAS" {
						total += 1
					}
				}

				// Check diagonal down right
				if lineNo+3 < bounds.MaxY && charNo+3 < bounds.MaxX {
					result := line[charNo] + wordSearch[lineNo+1][charNo+1] + wordSearch[lineNo+2][charNo+2] + wordSearch[lineNo+3][charNo+3]

					if result == "XMAS" {
						total += 1
					}
				}

				// Check diagonal up left
				if lineNo-3 > bounds.MinY && charNo-3 > bounds.MinX {
					result := line[charNo] + wordSearch[lineNo-1][charNo-1] + wordSearch[lineNo-2][charNo-2] + wordSearch[lineNo-3][charNo-3]

					if result == "XMAS" {
						total += 1
					}
				}

				// Check diagonal down left
				if lineNo+3 < bounds.MaxY && charNo-3 > bounds.MinX {
					result := line[charNo] + wordSearch[lineNo+1][charNo-1] + wordSearch[lineNo+2][charNo-2] + wordSearch[lineNo+3][charNo-3]

					if result == "XMAS" {
						total += 1
					}
				}
			}
		}
	}

	fmt.Println("Part1: ", total)
}
