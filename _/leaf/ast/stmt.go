package ast

import (
	"e8vm.net/e8/printer"
)

type Stmt interface {
	printer.Printable
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

func (self *BlockStmt) PrintTo(p printer.Iface) {
	p.Println("{")
	p.ShiftIn()

	for _, s := range self.Stmts {
		s.PrintTo(p)
	}

	p.ShiftOut("}")
}
