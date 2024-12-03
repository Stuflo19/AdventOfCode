package day2

import "testing"

func TestValidIncrementing(t *testing.T) {
	report := NewReport([]string{"1", "3", "6", "7", "9"})

	result := report.validIncrementing()

	if !result {
		panic("Unexpected")
	}
}

func TestInvalidIncrementing(t *testing.T) {
	report := NewReport([]string{"1", "2", "7", "8", "9"})

	result := report.validIncrementing()

	if result {
		panic("Unexpected")
	}

	report = NewReport([]string{"1", "3", "2", "4", "5"})

	result = report.validIncrementing()

	if result {
		panic("Unexpected")
	}
}

func TestValidDecrementing(t *testing.T) {
	report := NewReport([]string{"7", "6", "4", "2", "1"})

	result := report.validDecrementing()

	if !result {
		panic("Unexpected")
	}
}

func TestInvalidDecrementing(t *testing.T) {
	report := NewReport([]string{"9", "7", "6", "2", "1"})

	result := report.validDecrementing()

	if result {
		panic("Unexpected")
	}

	report = NewReport([]string{"8", "6", "4", "4", "1"})

	result = report.validDecrementing()

	if result {
		panic("Unexpected")
	}
}
