package ir1

import (
	"github.com/h8liu/e8/printer"
)

type Block struct {
	Stmts []Stmt
}

func NewBlock() *Block {
	ret := new(Block)
	return ret
}

func (self *Block) PrintTo(p printer.Interface) {
	for _, s := range self.Stmts {
		s.PrintTo(p)
	}
}

func (self *Block) S(s ...Stmt) {
	self.Stmts = append(self.Stmts, s...)
}
