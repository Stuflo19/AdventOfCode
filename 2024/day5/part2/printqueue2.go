package part2

import (
	"fmt"
	"maps"
	"reflect"
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
	queue []string,
	index int,
) int {
	for key := range maps.Keys(orderingRules) {
		if slices.Contains(orderingRules[key], item) && !slices.Contains(queue[:index+1], key) &&
			slices.Contains(queue, key) {
			return slices.Index(queue, key)
		}
	}

	return -1
}

func checkAfter(
	orderingRules map[string][]string,
	item string,
	queue []string,
	index int,
) int {
	for _, expectedAfter := range orderingRules[item] {
		if slices.Contains(queue[:index+1], expectedAfter) && slices.Contains(queue, item) {
			return slices.Index(queue, expectedAfter)
		}
	}

	return -1
}

func processPrintQueue(queue []string, orderingRules map[string][]string) (int, error) {
	hadError := false
	swapper := reflect.Swapper(queue)

	for index := range queue {
		invalid := true
		for invalid {
			beforeResult := checkBefore(orderingRules, queue[index], queue, index)
			if beforeResult != -1 {
				hadError = true
				swapper(index, beforeResult)
			}
			afterResult := checkAfter(orderingRules, queue[index], queue, index)
			if afterResult != -1 {
				hadError = true
				swapper(index, afterResult)
			}

			if beforeResult == -1 && afterResult == -1 {
				invalid = false
			}
		}
	}

	if hadError {
		return strconv.Atoi(queue[len(queue)/2])
	} else {
		return 0, nil
	}
}

func Part2(orderingRulesFile string, updatesFile string) {
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

	fmt.Println("Part 2:", total)
}
