package part2

import (
	"testing"
)

func TestPart2FakeInput(t *testing.T) {
	Part2("../test_ordering_rules.txt", "../test_update.txt")
}

func TestPart2RealInput(t *testing.T) {
	Part2("../ordering_rules.txt", "../updates.txt")
}
