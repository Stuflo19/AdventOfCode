package day2

import (
	"strconv"
)

const (
	INCREMENTING = iota
	DECREMENTING
)

type Report struct {
	reportData []string
}

func NewReport(reportData []string) Report {
	return Report{
		reportData,
	}
}

func (r Report) getDirection() int {
	firstInt, _ := strconv.Atoi(r.reportData[0])
	secondInt, _ := strconv.Atoi(r.reportData[1])

	if firstInt > secondInt {
		return DECREMENTING
	} else {
		return INCREMENTING
	}
}

func (r Report) validIncrementing() bool {
	for index := 0; index < len(r.reportData)-1; index++ {
		currInt, _ := strconv.Atoi(r.reportData[index])
		nextInt, _ := strconv.Atoi(r.reportData[index+1])
		diff := nextInt - currInt

		if currInt > nextInt || diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func (r Report) validDecrementing() bool {
	for index := 0; index < len(r.reportData)-1; index++ {
		currInt, _ := strconv.Atoi(r.reportData[index])
		nextInt, _ := strconv.Atoi(r.reportData[index+1])
		diff := currInt - nextInt

		if currInt < nextInt || diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func (r Report) validateReport() bool {
	direction := r.getDirection()

	if direction == INCREMENTING {
		return r.validIncrementing()
	} else {
		return r.validDecrementing()
	}
}
