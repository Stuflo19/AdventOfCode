package day5

import (
	"fmt"

	"aoc2020/lib/helpers"
)

func processRow(rowData string) int {
	searchSpace := NewSearchSpace(ROWS)

	for index, value := range rowData {
		if value == 'F' {
			if index == len(rowData)-1 {
				return searchSpace.lower
			}

			searchSpace.lowerSearchSpace()
		} else if value == 'B' {
			if index == len(rowData)-1 {
				return searchSpace.higher
			}
			searchSpace.higherSearchSpace()
		}
	}

	return -1
}

func processColumn(columnData string) int {
	searchSpace := NewSearchSpace(COLUMNS)

	for index, value := range columnData {
		if value == 'L' {
			if index == len(columnData)-1 {
				return searchSpace.lower
			}

			searchSpace.lowerSearchSpace()
		} else if value == 'R' {
			if index == len(columnData)-1 {
				return searchSpace.higher
			}

			searchSpace.higherSearchSpace()
		}
	}

	return -1
}

func processTicket(seatData string) int {
	rowData := seatData[:7]
	columnData := seatData[7:10]

	rowNumber := processRow(rowData)
	columnNumber := processColumn(columnData)

	return rowNumber*8 + columnNumber
}

func Part1(filename string) {
	maxSeatID := 0
	contents := helpers.ReadFileWithoutSplit(filename)

	for _, seatData := range contents {
		result := processTicket(seatData)

		if result > maxSeatID {
			maxSeatID = result
		}
	}

	fmt.Println("Max Seat ID: ", maxSeatID)
}
