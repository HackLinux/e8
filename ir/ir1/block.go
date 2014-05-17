package ir1

type Block struct {
	stmts []*Stmt
}

func NewBlock() *Block {
	ret := new(Block)
	return ret
}
