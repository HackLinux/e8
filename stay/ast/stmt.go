package ast

type Stmt interface {
}

type BlockStmt struct {
	Stmts []Stmt
}

func NewBlock() *BlockStmt {
	ret := new(BlockStmt)
	ret.Stmts = make([]Stmt, 0, 128)

	return ret
}

func (self *BlockStmt) Add(s Stmt) {
	self.Stmts = append(self.Stmts, s)
}

type CallStmt struct {
}
