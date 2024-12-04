package part2

import "testing"

func TestFakeInputPart2(t *testing.T) {
	Part2("../test.txt")
}

func TestFakeRevealedInputPart2(t *testing.T) {
	Part2("../test_revealed2.txt")
}

func TestRealInputPart2(t *testing.T) {
	Part2("../input.txt")
}
