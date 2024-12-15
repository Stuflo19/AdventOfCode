package part2

import "testing"

func TestPart2Example(t *testing.T) {
	Part2("../example.txt", 5, 11, 7)
}

func TestPart2(t *testing.T) {
	Part2("../test.txt", 100, 11, 7)
}

func TestPart2Real(t *testing.T) {
	Part2("../input.txt", 10000, 101, 103)
}
