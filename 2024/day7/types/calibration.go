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

//func (c *Calibration) FindTimesOrAddForExpected2() int {
//	operators := []string{"+", "*", "||"}
//	res := CartesianProduct(operators, len(c.values)-1)
//
//	fmt.Println("All Operators: ", res)
//
//	for _, op := range res {
//		joined := make([]int, 0, len(c.values))
//		withoutOr := make([]string, 0, len(op))
//
//		if slices.Contains(op, "||") {
//			for idx, o := range op {
//				if o == "||" {
//					curr := strconv.Itoa(c.values[idx])
//					next := strconv.Itoa(c.values[idx+1])
//
//					res, _ := strconv.Atoi(curr + next)
//					joined = append(joined, res)
//				} else {
//					joined = append(joined, c.values[idx])
//					withoutOr = append(withoutOr, o)
//				}
//			}
//		} else {
//			withoutOr = append(withoutOr, op...)
//			joined = append(joined, c.values...)
//		}
//
//		c.Total = joined[0]
//
//		fmt.Println("Found expected")
//		fmt.Println("Joined:", joined)
//		fmt.Println("With or:", op)
//		fmt.Println("Without or:", withoutOr)
//		if c.Total == c.expected {
//			//			fmt.Println("Found expected")
//			//			fmt.Println("Joined:", joined)
//			//			fmt.Println("Without or:", withoutOr)
//			return c.Total
//		} else if len(withoutOr) == 0 {
//			continue
//		}
//
//		for index, value := range joined[1:] {
//			if withoutOr[index] == "+" {
//				c.add(value)
//			} else if withoutOr[index] == "*" {
//				c.times(value)
//			}
//		}
//
//		if c.Total == c.expected {
//			//	fmt.Println("Found expected")
//			//	fmt.Println("Joined:", joined)
//			//	fmt.Println("Without or:", withoutOr)
//			return c.Total
//		} else {
//			c.Total = c.values[0]
//		}
//	}
//
//	return 0
//}

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
