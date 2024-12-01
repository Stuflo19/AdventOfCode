package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func createLists(content []ListPair) ([]int, []int) {
	list1 := []int{}
	list2 := []int{}

	for _, listPair := range content {
		item1, _ := strconv.Atoi(listPair.item1)
		item2, _ := strconv.Atoi(listPair.item2)

		list1 = append(list1, item1)
		list2 = append(list2, item2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2
}

func Part1(filename string) {
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

	list1, list2 := createLists(contents)

	total := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			total += list1[i] - list2[i]
		} else {
			total += list2[i] - list1[i]
		}
	}

	fmt.Println("Total:", total)
}
