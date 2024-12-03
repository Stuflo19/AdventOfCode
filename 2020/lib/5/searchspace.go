package day5

const (
	ROWS    = 127
	COLUMNS = 7
)

type SearchSpace struct {
	lower  int
	higher int
}

func NewSearchSpace(size int) SearchSpace {
	return SearchSpace{
		higher: size,
		lower:  0,
	}
}

func (s *SearchSpace) lowerSearchSpace() {
	newValue := (s.higher - s.lower) / 2
	s.higher = s.higher - newValue - 1
}

func (s *SearchSpace) higherSearchSpace() {
	newValue := (s.higher - s.lower) / 2
	s.lower = s.lower + newValue + 1
}
