package day3

type SledWithSlope struct {
	position Coordinates
	treesHit int
  slope Coordinates
}

func NewSledWithSlope(slope Coordinates) SledWithSlope {
	return SledWithSlope{
		position: NewCoordinates(0, 0),
		treesHit: 0,
    slope: slope,
	}
}

func (s *SledWithSlope) moveSled(width int) {
	s.position.x = (s.position.x + s.slope.x) % width
	s.position.y += s.slope.y
}

func (s *SledWithSlope) hitTree() {
	s.treesHit += 1
}

