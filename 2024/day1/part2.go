package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countOccurances(content []ListPair) ([]string, map[string]int) {
	list := []string{}
	occurances := make(map[string]int)

	for _, listPair := range content {
		list = append(list, listPair.item1)
		occurances[listPair.item2] += 1
	}

	return list, occurances
}

func Part2(filename string) {
	contents := []ListPair{}
	file, err := os.Open(filename)

	if err != nil {
		panic("File not found")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")
		contents = append(contents, NewListPair(split[0], split[1]))
	}

	if err := scanner.Err(); err != nil {
		panic("Error with scanner")
	}

	list, occurances := countOccurances(contents)

	total := 0
	for _, value := range list {
		integerValue, _ := strconv.Atoi(value)
		total += integerValue * occurances[value]
	}

	fmt.Println("Total:", total)
}
