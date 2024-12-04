package part1

import "testing"

func TestFakeInputPart1(t *testing.T) {
	Part1("../test.txt")
}

func TestFakeRevealedInputPart1(t *testing.T) {
	Part1("../test_revealed.txt")
}

func TestRealInputPart1(t *testing.T) {
	Part1("../input.txt")
}
