package part1

import (
	"fmt"
	"regexp"
	"strconv"

	"stuflo19/aoc2024/helpers"
)

func Part1(filename string) {
	total := 0
	r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	contents := helpers.ReadFile(filename)

	for _, line := range contents {
		matchingConditions := r.FindAllStringSubmatch(line, -1)

		for _, valid := range matchingConditions {
			first, _ := strconv.Atoi(valid[1])
			second, _ := strconv.Atoi(valid[2])

			total += first * second
		}
	}

	fmt.Println("Part 1:", total)
}
