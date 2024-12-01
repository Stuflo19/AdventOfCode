package day3

type Sled struct {
	position Coordinates
	treesHit int
}

func NewSled() Sled {
	return Sled{
		position: NewCoordinates(0, 0),
		treesHit: 0,
	}
}

func (s *Sled) moveSled(width int) {
	s.position.x = (s.position.x + MOVE_X) % width
	s.position.y += MOVE_Y
}

func (s *Sled) hitTree() {
	s.treesHit += 1
}
