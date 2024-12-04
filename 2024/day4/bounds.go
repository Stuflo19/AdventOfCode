package day4

type Bounds struct {
	MaxX int
	MinX int
	MaxY int
	MinY int
}

func NewBounds(MaxX int, MinX int, MaxY int, MinY int) Bounds {
	return Bounds{
		MaxX,
		MinX,
		MaxY,
		MinY,
	}
}
