package types

type FreeSpace struct {
	startPos int
	endPos   int
}

func NewFreeSpace(startPos int, endPos int) FreeSpace {
	return FreeSpace{
		startPos: startPos,
		endPos:   endPos,
	}
}

func (f *FreeSpace) GetStartPos() int {
	return f.startPos
}

func (f *FreeSpace) GetEndPos() int {
	return f.endPos
}

func (f *FreeSpace) ReduceSize(decrease int) {
	f.startPos += decrease
	f.endPos -= 1
}

func (f *FreeSpace) GetSize() int {
	return f.endPos - f.startPos + 1
}
