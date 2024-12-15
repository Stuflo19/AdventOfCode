package part1

import "testing"

func TestPartExampleInput1(t *testing.T) {
	total := Part1("0 1 10 99 999", 1)

	if total != 7 {
		t.Errorf("Expected 7, got %d", total)
	}
}

func TestPartTestInput1(t *testing.T) {
	total := Part1("125 17", 6)

	if total != 22 {
		t.Errorf("Expected 22, got %d", total)
	}
}

func TestPartTest2Input1(t *testing.T) {
	total := Part1("125 17", 25)

	if total != 55312 {
		t.Errorf("Expected 55312, got %d", total)
	}
}

func TestPartRealInput1(t *testing.T) {
	Part1("2 72 8949 0 981038 86311 246 7636740", 25)
}

func TestPartRealInput2(t *testing.T) {
	Part1("2 72 8949 0 981038 86311 246 7636740", 75)
}
