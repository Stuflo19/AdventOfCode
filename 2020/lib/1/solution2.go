package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution2() {
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

	for _, i := range fileLines {
		for _, j := range fileLines {
			for _, k := range fileLines {
				if i != j && i != k && j != k {
					iInt, _ := strconv.Atoi(i)
					jInt, _ := strconv.Atoi(j)
					kInt, _ := strconv.Atoi(k)
					result := iInt + jInt + kInt

					if result == 2020 {
						println("Part 2 Answer: " + strconv.Itoa(iInt*jInt*kInt))
						return
					}
				}
			}
		}
	}
}
