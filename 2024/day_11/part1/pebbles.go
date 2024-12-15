package part1

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInputToMap(input string) map[int]int {
	stonesMap := make(map[int]int)
	splitString := strings.Split(input, " ")

	for _, s := range splitString {
		i, _ := strconv.Atoi(s)
		stonesMap[i] = 1
	}

	return stonesMap
}

// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
func blink(stonesMap map[int]int) map[int]int {
	newMap := make(map[int]int)

	for key, value := range stonesMap {
		strKey := strconv.Itoa(key)
		isEvenNum := len(strKey)%2 == 0
		if key == 0 {
			newMap[1] += value
		} else if isEvenNum {
			left := strKey[:len(strKey)/2]
			right := strKey[len(strKey)/2:]
			leftInt, _ := strconv.Atoi(left)
			rightInt, _ := strconv.Atoi(right)
			newMap[leftInt] += value
			newMap[rightInt] += value
		} else {
			newNum := key * 2024
			newMap[newNum] += value
		}
	}

	return newMap
}

func Part1(input string, blinks int) int {
	stonesMap := parseInputToMap(input)

	for i := 0; i < blinks; i++ {
		stonesMap = blink(stonesMap)
	}

	total := 0
	for _, value := range stonesMap {
		total += value
	}

	fmt.Println("Part 1:", total)

	return total
}
