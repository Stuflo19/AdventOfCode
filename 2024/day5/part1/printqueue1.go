package part1

import (
	"fmt"
	"maps"
	"slices"
	"strconv"

	"stuflo19/aoc2024/helpers"
)

func parseOrderingRules(orderRules [][]string) map[string][]string {
	result := make(map[string][]string)

	for _, rule := range orderRules {
		result[rule[0]] = append(result[rule[0]], rule[1])
	}

	return result
}

func checkBefore(
	orderingRules map[string][]string,
	item string,
	processed []string,
	queue []string,
) bool {
	for key := range maps.Keys(orderingRules) {
		if slices.Contains(orderingRules[key], item) && !slices.Contains(processed, key) &&
			slices.Contains(queue, key) {
			//fmt.Println("checkBefore found printed wrong:", item)
			return false
		}
	}

	return true
}

func checkAfter(
	orderingRules map[string][]string,
	item string,
	processed []string,
	queue []string,
) bool {
	for _, expectedAfter := range orderingRules[item] {
		if slices.Contains(processed, expectedAfter) && slices.Contains(queue, item) {
			//fmt.Println("checkBefore found krinted wrong:", item)
			return false
		}
	}

	return true
}

func processPrintQueue(queue []string, orderingRules map[string][]string) (int, error) {
	processed := []string{}

	for _, item := range queue {
		if !checkBefore(orderingRules, item, processed, queue) {
			return 0, fmt.Errorf("wrong before")
		}
		if !checkAfter(orderingRules, item, processed, queue) {
			return 0, fmt.Errorf("wrong after")
		}

		processed = append(processed, item)
	}

	return strconv.Atoi(queue[len(queue)/2])
}

func Part1(orderingRulesFile string, updatesFile string) {
	total := 0

	orderingRules := helpers.ReadFileWithSeparator(orderingRulesFile, "|")
	updates := helpers.ReadFileWithSeparator(updatesFile, ",")

	orderingRulesMap := parseOrderingRules(orderingRules)

	for _, update := range updates {
		processed, err := processPrintQueue(update, orderingRulesMap)

		if err != nil {
			continue
		} else {
			total += processed
		}
	}

	fmt.Println("Part 1:", total)
}
