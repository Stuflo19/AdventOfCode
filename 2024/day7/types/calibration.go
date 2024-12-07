package types

import (
	"strconv"
)

type Calibration struct {
	expected int
	values   []int
	Total    int
}

func NewCalibration(expected int, values []int) Calibration {
	return Calibration{
		expected: expected,
		values:   values,
		Total:    values[0],
	}
}

func (c *Calibration) add(value int) {
	c.Total += value
}

func (c *Calibration) times(value int) {
	c.Total *= value
}

func (c *Calibration) append(value int) {
	strValue := strconv.Itoa(value)
	strTotal := strconv.Itoa(c.Total)

	c.Total, _ = strconv.Atoi(strTotal + strValue)
}

func (c *Calibration) FindTimesOrAddForExpected() int {
	operators := []string{"+", "*"}
	res := CartesianProduct(operators, len(c.values)-1)

	for _, op := range res {
		for index, value := range c.values[1:] {
			if op[index] == "+" {
				c.add(value)
			} else {
				c.times(value)
			}
		}

		if c.Total == c.expected {
			return c.Total
		} else {
			c.Total = c.values[0]
		}
	}

	return 0
}
func (c *Calibration) FindTimesOrAddForExpected2() int {
	operators := []string{"+", "*", "||"}
	res := CartesianProduct(operators, len(c.values)-1)

	for _, op := range res {
		for index, value := range c.values[1:] {
			if op[index] == "+" {
				c.add(value)
			} else if op[index] == "*" {
				c.times(value)

			} else if op[index] == "||" {
				c.append(value)
			}
		}

		if c.Total == c.expected {
			return c.Total
		} else {
			c.Total = c.values[0]
		}
	}

	return 0
}

// CartesianProduct generates the Cartesian product of the input array repeated n times.
func CartesianProduct(elements []string, n int) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	// Generate the Cartesian product for n-1
	subProduct := CartesianProduct(elements, n-1)
	var result [][]string

	// Append each element to each sub-product
	for _, sub := range subProduct {
		for _, el := range elements {
			combination := append([]string{}, sub...) // Copy sub to a new slice
			combination = append(combination, el)
			result = append(result, combination)
		}
	}

	return result
}
