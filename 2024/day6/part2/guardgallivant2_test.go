package part2

import "testing"

func TestPart2FakeInput(t *testing.T) {
	Part2("../test.txt")
}

func TestPart2RealInput(t *testing.T) {
	Part2("../input.txt")
}
