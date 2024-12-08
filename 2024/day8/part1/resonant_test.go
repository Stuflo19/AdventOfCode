package part1

import "testing"

func TestExampleInputPart1(t *testing.T) {
	Part1("../example.txt")
}

func TestExampleInput2Part1(t *testing.T) {
	Part1("../example2.txt")
}

func TestFakeInputPart1(t *testing.T) {
	Part1("../test.txt")
}

func TestRealInputPart1(t *testing.T) {
	Part1("../input.txt")
}
