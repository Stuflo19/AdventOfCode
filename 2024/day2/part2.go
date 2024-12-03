package day2

import (
	"fmt"

	"stuflo19/aoc2024/helpers"
)

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func filter[T any](slice []T, i int) []T {
	arr := []T{}

	for idx, value := range slice {
		if i == idx {
			continue
		}
		arr = append(arr, value)
	}

	return arr
}

func Part2(filename string) {
	totalValidReports := 0
	contents := helpers.ReadFileWithSeparator(filename, " ")

MainLoop:
	for _, reportData := range contents {
		report := NewReport(reportData)

		if report.validateReport() {
			fmt.Println("Passing report: ", reportData)
			totalValidReports += 1
		} else {
			for i := 0; i < len(reportData); i++ {
				currentData := filter(reportData, i)
				report = NewReport(currentData)

				if report.validateReport() {
					fmt.Println("Passing report after removal: ", currentData)
					totalValidReports += 1
					continue MainLoop
				}
			}
		}
	}

	fmt.Println("Total Valid: ", totalValidReports)
}
