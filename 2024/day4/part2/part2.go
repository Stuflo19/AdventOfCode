package part2

import (
	"fmt"

	"stuflo19/aoc2024/day4"
	"stuflo19/aoc2024/helpers"
)

func Part2(filename string) {
	total := 0
	wordSearch := helpers.ReadFileWithSeparator(filename, "")

	bounds := day4.NewBounds(len(wordSearch[0]), -1, len(wordSearch), -1)

	for lineNo := range wordSearch {
		for charNo, char := range wordSearch[lineNo] {
			if char == "A" {
				if charNo+1 < bounds.MaxX && charNo-1 > bounds.MinX && lineNo+1 < bounds.MaxY &&
					lineNo-1 > bounds.MinY {
					if wordSearch[lineNo+1][charNo+1] == "S" &&
						wordSearch[lineNo+1][charNo-1] == "S" &&
						wordSearch[lineNo-1][charNo+1] == "M" &&
						wordSearch[lineNo-1][charNo-1] == "M" {
						total += 1
					} else if wordSearch[lineNo+1][charNo+1] == "M" &&
						wordSearch[lineNo+1][charNo-1] == "M" &&
						wordSearch[lineNo-1][charNo+1] == "S" &&
						wordSearch[lineNo-1][charNo-1] == "S" {
						total += 1
					} else if wordSearch[lineNo+1][charNo+1] == "S" &&
						wordSearch[lineNo+1][charNo-1] == "M" &&
						wordSearch[lineNo-1][charNo+1] == "S" &&
						wordSearch[lineNo-1][charNo-1] == "M" {
						total += 1
					} else if wordSearch[lineNo+1][charNo+1] == "M" &&
						wordSearch[lineNo+1][charNo-1] == "S" &&
						wordSearch[lineNo-1][charNo+1] == "M" &&
						wordSearch[lineNo-1][charNo-1] == "S" {
						total += 1
					}
				}
			}
		}
	}

	fmt.Println("Part 2:", total)
}
