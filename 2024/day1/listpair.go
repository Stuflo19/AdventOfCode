package day1

type ListPair struct {
	item1 string
	item2 string
}

func NewListPair(item1 string, item2 string) ListPair {
	return ListPair{
		item1,
		item2,
	}
}
