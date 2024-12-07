package part1

import "testing"

func TestPart1FakeInput(t *testing.T) {
	Part1("../test.txt")
}

func TestPart1RealInput(t *testing.T) {
	Part1("../input.txt")
}
