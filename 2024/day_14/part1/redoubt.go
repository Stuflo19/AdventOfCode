package part1

import (
	"fmt"

	day14 "stuflo19/aoc2024/day_14"
	"stuflo19/aoc2024/helpers"
)

func Part1(filename string, iterations int, maxX int, maxY int) {
	contents := helpers.ReadFile(filename)
	quadrants := map[int]int{}
	robots := []day14.Robot{}

	for _, line := range contents {
		robots = append(robots, day14.NewRobot(line))
	}

	for _, r := range robots {
		r.MoveTimes(iterations, maxX, maxY)
		q := r.GetQuadrant(maxX, maxY)

		quadrants[q]++
	}

	total := 1
	for k, v := range quadrants {
		if k == -1 {
			continue
		}
		total = total * v
	}

	fmt.Println("Part 1:", total)
}
