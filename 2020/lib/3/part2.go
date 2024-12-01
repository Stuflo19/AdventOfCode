package day3

import (
	"fmt"

	"aoc2020/lib/helpers"
)

func Part2(filename string) {
	fileContents := helpers.ReadFile(filename)

	firstSlope := NewCoordinates(1, 1)
	secondSlope := NewCoordinates(3, 1)
	thirdSlope := NewCoordinates(5, 1)
	fourthSlope := NewCoordinates(7, 1)
	fifthSlope := NewCoordinates(1, 2)

	sledsWithSlopes := []SledWithSlope{
		NewSledWithSlope(firstSlope),
		NewSledWithSlope(secondSlope),
		NewSledWithSlope(thirdSlope),
		NewSledWithSlope(fourthSlope),
		NewSledWithSlope(fifthSlope),
	}

	grid := NewSlopes(fileContents, sledsWithSlopes)

	for currSled := 0; currSled < len(grid.sledsWithSlopes); currSled++ {
		for grid.sledsWithSlopes[currSled].position.y < len(grid.forest)-1 {
			grid.slide(currSled)
		}
	}

	fmt.Println("Trees hit: ", grid.getTotal())
}
