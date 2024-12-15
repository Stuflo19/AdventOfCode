package part1

import "testing"

func TestPart1(t *testing.T) {
	Part1("../test.txt")
}

func TestExamplePart1(t *testing.T) {
	Part1("../example.txt")
}

func TestPart1Real(t *testing.T) {
	Part1("../input.txt")
}
