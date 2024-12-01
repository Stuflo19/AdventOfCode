package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution1() {
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
		min, _ := strconv.Atoi(allowedUsages[0])
		max, _ := strconv.Atoi(allowedUsages[1])

		usages := strings.Count(splitLine[2], string(splitLine[1][0]))

		if min <= usages && max >= usages {
			total += 1
		}
	}

	totalString := strconv.Itoa(total)

	println("Solution 1 Answer: " + totalString)
}
