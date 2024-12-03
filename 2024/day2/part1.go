package day2

import (
	"fmt"

	"stuflo19/aoc2024/helpers"
)

func Part1(filename string) {
	totalValidReports := 0
	contents := helpers.ReadFileWithSeparator(filename, " ")

	for _, reportData := range contents {
		report := NewReport(reportData)

		if report.validateReport() {
			totalValidReports += 1
		}
	}

	fmt.Println("Total Valid: ", totalValidReports)
}
