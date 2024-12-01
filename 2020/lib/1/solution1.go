package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution1() {
	readFile, err := os.Open("lib/1/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	err = readFile.Close()
	if err != nil {
		return
	}

	for _, subject := range fileLines {
		for _, line := range fileLines {
			if subject != line {
				subjectInt, _ := strconv.Atoi(subject)
				lineInt, _ := strconv.Atoi(line)
				result := subjectInt + lineInt

				if result == 2020 {
					println("Part 1 Answer: " + strconv.Itoa(lineInt*subjectInt))
					return
				}
			}
		}
	}
}
