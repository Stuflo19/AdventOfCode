package types

import (
	"image"
	"math"
)

type Machine struct {
	buttonA image.Point
	buttonB image.Point
	prize   image.Point
}

func NewMachine(buttonA image.Point, buttonB image.Point, prize image.Point) Machine {
	return Machine{
		buttonA: buttonA,
		buttonB: buttonB,
		prize:   prize,
	}
}

func NewMachine2(buttonA image.Point, buttonB image.Point, prize image.Point) Machine {
	return Machine{
		buttonA: buttonA,
		buttonB: buttonB,
		prize:   image.Point{X: prize.X + 10000000000000, Y: prize.Y + 10000000000000},
	}
}

func isWholeNumber(f float64) bool {
	return math.Mod(f, 1) == 0
}

func (m Machine) FindShortestPath() (int, int) {
	det := m.buttonA.X*m.buttonB.Y - m.buttonA.Y*m.buttonB.X
	if det == 0 {
		return 0, 0
	}

	invAX, invAY := m.buttonB.Y, -m.buttonA.Y
	invBX, invBY := -m.buttonB.X, m.buttonA.X

	a := (invAX*m.prize.X + invBX*m.prize.Y)
	b := (invAY*m.prize.X + invBY*m.prize.Y)

	nA := a / det
	nB := b / det

	if nA < 0 || nB < 0 || a != nA*det || b != nB*det {
		return 0, 0
	}

	if nA > 100 || nB > 100 {
		return 0, 0
	}

	return nA, nB
}

func (m Machine) FindShortestPath2() (int, int) {
	det := m.buttonA.X*m.buttonB.Y - m.buttonA.Y*m.buttonB.X
	if det == 0 {
		return 0, 0
	}

	invAX, invAY := m.buttonB.Y, -m.buttonA.Y
	invBX, invBY := -m.buttonB.X, m.buttonA.X

	a := (invAX*m.prize.X + invBX*m.prize.Y)
	b := (invAY*m.prize.X + invBY*m.prize.Y)

	nA := a / det
	nB := b / det

	if nA < 0 || nB < 0 || a != nA*det || b != nB*det {
		return 0, 0
	}

	return nA, nB
}
