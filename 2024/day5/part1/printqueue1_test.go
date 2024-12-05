package part1

import (
	"testing"
)

func TestPart1FakeInput(t *testing.T) {
	Part1("../test_ordering_rules.txt", "../test_update.txt")
}

func TestPart1RealInput(t *testing.T) {
	Part1("../ordering_rules.txt", "../updates.txt")
}
