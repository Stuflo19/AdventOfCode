package types

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type Guard struct {
	position Position
	facing   int
	bounds   Bounds
}

func NewGuard(position Position, facing int, bounds Bounds) Guard {
	return Guard{
		position,
		facing,
		bounds,
	}
}

func (g *Guard) rotate() {
	if g.facing != 3 {
		g.facing += 1
	} else {
		g.facing = 0
	}
}

func (g *Guard) step(x int, y int) {
	g.position.x = x
	g.position.y = y
}
