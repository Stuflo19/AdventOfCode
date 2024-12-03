package part2

import (
	"fmt"
	"regexp"
	"strconv"

	"stuflo19/aoc2024/helpers"
)

const (
	DO   = "do()"
	DONT = "don't()"
)

func Part2(filename string) {
	total := 0
	active := true

	r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|don't\(\)|do\(\)`)
	contents := helpers.ReadFile(filename)

	for _, line := range contents {
		matchingConditions := r.FindAllStringSubmatch(line, -1)

		for _, valid := range matchingConditions {
			if valid[0] == DO {
				active = true
			} else if valid[0] == DONT {
				active = false
			} else if active {
				first, _ := strconv.Atoi(valid[1])
				second, _ := strconv.Atoi(valid[2])

				total += first * second
			}
		}
	}

	fmt.Println("Part 2:", total)
}
