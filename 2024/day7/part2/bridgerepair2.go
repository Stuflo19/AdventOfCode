package part2

import (
	"fmt"
	"strconv"
	"strings"

	"stuflo19/aoc2024/day7/types"
	"stuflo19/aoc2024/helpers"
)

func parseLine(line string) (int, []int) {
	split := strings.Split(line, ": ")

	expected, _ := strconv.Atoi(split[0])

	numbersSplit := strings.Split(split[1], " ")

	numbers := make([]int, len(numbersSplit))
	for i, number := range numbersSplit {
		numbers[i], _ = strconv.Atoi(number)
	}

	return expected, numbers
}

func Part2(filename string) {
	fileContents := helpers.ReadFile(filename)
	total := 0

	for _, line := range fileContents {
		expected, numbers := parseLine(line)

		calibration := types.NewCalibration(expected, numbers)
		total += calibration.FindTimesOrAddForExpected2()
	}

	fmt.Printf("Part 2: %v\n", total)
}
