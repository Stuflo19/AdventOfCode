package part1

import "testing"

func TestPart1Example(t *testing.T) {
	Part1("../example.txt", 5, 11, 7)
}

func TestPart1(t *testing.T) {
	Part1("../test.txt", 100, 11, 7)
}

func TestPart1Real(t *testing.T) {
	Part1("../input.txt", 100, 101, 103)
}
