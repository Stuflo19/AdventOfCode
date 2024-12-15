package types

type Block struct {
	id int
}

func NewBlock(id int) Block {
	return Block{
		id: id,
	}
}

func (b *Block) GetID() int {
	return b.id
}
