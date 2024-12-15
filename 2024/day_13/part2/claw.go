package part2

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"stuflo19/aoc2024/day_13/types"
	"stuflo19/aoc2024/helpers"
)

func parseLine(line string, separator string) (int, int) {
	coords := strings.Split(line, ", ")

	splitX := strings.Split(coords[0], separator)
	splitY := strings.Split(coords[1], separator)

	x, _ := strconv.Atoi(splitX[1])
	y, _ := strconv.Atoi(splitY[1])

	return x, y
}

func Part2(filename string) {
	total := 0
	buttonA := image.Point{}
	buttonB := image.Point{}
	prize := image.Point{}

	fileContents := helpers.ReadFile(filename)

	for _, line := range fileContents {

		if len(line) == 0 {
			machine := types.NewMachine2(
				buttonA,
				buttonB,
				prize,
			)

			pressesA, pressesB := machine.FindShortestPath2()

			total += (pressesA * 3) + pressesB

		} else if cutA, ok := strings.CutPrefix(line, "Button A:"); ok {
			x, y := parseLine(cutA, "+")
			buttonA = image.Point{x, y}
		} else if cutB, ok := strings.CutPrefix(line, "Button B:"); ok {
			x, y := parseLine(cutB, "+")
			buttonB = image.Point{x, y}
		} else if cutP, ok := strings.CutPrefix(line, "Prize:"); ok {
			x, y := parseLine(cutP, "=")
			prize = image.Point{x, y}
		}
	}

	fmt.Println(total)
}
