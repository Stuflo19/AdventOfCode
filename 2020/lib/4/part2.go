package day4

import (
	"fmt"

	"aoc2020/lib/helpers"
)

func Part2(filename string) {
	validPassports := 0
	fileContents := helpers.ReadFileWithoutSplit(filename)
	currPassport := ""

	for i := 0; i < len(fileContents); i++ {
		if len(fileContents[i]) == 0 {
			passport := NewPassport(currPassport)

			if passport.validatePassport2() {
				validPassports += 1
			}

			currPassport = ""
		} else {
			currPassport += " " + fileContents[i]
		}
	}

	fmt.Println("Valid Passports: ", validPassports)
}
