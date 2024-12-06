package types

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) Position {
	return Position{
		x, y,
	}
}
