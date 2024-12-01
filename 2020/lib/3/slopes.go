package day3

type Slopes struct {
  forest [][]string
  sledsWithSlopes []SledWithSlope
  width int
}
func (s *Slopes) checkCollision(currSled int) {
	if s.forest[s.sledsWithSlopes[currSled].position.y][s.sledsWithSlopes[currSled].position.x] == TREE {
		s.sledsWithSlopes[currSled].hitTree()
	}
}

func (s *Slopes) slide(currSled int) {
	s.sledsWithSlopes[currSled].moveSled(s.width)

	s.checkCollision(currSled)
}

func (s Slopes) getTotal() int {
  total := 0

  for currSled := 0; currSled < len(s.sledsWithSlopes); currSled++ {
    if total == 0 {
      total = s.sledsWithSlopes[currSled].treesHit
    } else {
      total *= s.sledsWithSlopes[currSled].treesHit
    }
  }

  return total
}

func NewSlopes(forest [][]string, sledsWithSlopes []SledWithSlope) Slopes {
	return Slopes{
		forest: forest,
		sledsWithSlopes: sledsWithSlopes,
		width:  len(forest[0]),
	}
}
