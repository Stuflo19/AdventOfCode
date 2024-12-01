package day3

type Grid struct {
	forest [][]string
	sled   Sled
	width  int
}

func (g *Grid) checkCollision() {
	if g.forest[g.sled.position.y][g.sled.position.x] == TREE {
		g.sled.hitTree()
	}
}

func (g *Grid) slide() {
	g.sled.moveSled(g.width)

	g.checkCollision()
}

func NewGrid(forest [][]string) Grid {
	return Grid{
		forest: forest,
		sled:   NewSled(),
		width:  len(forest[0]),
	}
}
