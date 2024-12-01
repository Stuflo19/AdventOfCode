package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution2() {
	total := 0

	readFile, err := os.Open("lib/2/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLine := strings.Split(line, " ")
		allowedUsages := strings.Split(splitLine[0], "-")
		pos1, _ := strconv.Atoi(allowedUsages[0])
		pos2, _ := strconv.Atoi(allowedUsages[1])

		if (splitLine[2][pos1-1] == splitLine[1][0] && splitLine[2][pos2-1] != splitLine[1][0]) ||
			(splitLine[2][pos1-1] != splitLine[1][0] && splitLine[2][pos2-1] == splitLine[1][0]) {
			total += 1
		}
	}

	totalString := strconv.Itoa(total)

	println("Solution 2 Answer: " + totalString)
}
