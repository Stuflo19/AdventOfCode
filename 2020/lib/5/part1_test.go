package day5

import (
	"fmt"
	"testing"
)

func TestProcessRow(t *testing.T) {
	result := processRow("FBFBBFF")

	if result != 44 {
		panic("Unexpected")
	}

	result2 := processRow("BFFFBBF")

	if result2 != 70 {
		panic("Unexpected")
	}
}

func TestProcessColumn(t *testing.T) {
	result := processColumn("RLR")

	fmt.Println("Result:", result)
	if result != 5 {
		panic("Unexpected")
	}

	result2 := processColumn("RLL")
	fmt.Println("Result:", result2)

	if result2 != 4 {
		panic("Unexpected")
	}
}

func TestProcessTicket(t *testing.T) {
	result := processTicket("FBFBBFFRLR")

	if result != 357 {
		panic("Unexpected")
	}

	result2 := processTicket("BFFFBBFRRR")

	if result2 != 567 {
		panic("Unexpected")
	}

	result3 := processTicket("FFFBBBFRRR")

	if result3 != 119 {
		panic("Unexpected")
	}

	result4 := processTicket("BBFFBBFRLL")

	if result4 != 820 {
		panic("Unexpected")
	}
}

func TestPart1(t *testing.T) {
	Part1("input.txt")
}
