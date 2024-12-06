package types

type Bounds struct {
	maxX int
	maxY int
	minX int
	minY int
}

func NewBounds(maxX int, maxY int, minX int, minY int) Bounds {
	return Bounds{
		maxX,
		maxY,
		minX,
		minY,
	}
}
