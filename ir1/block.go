package ir1

import (
	"github.com/h8liu/e8/printer"
)

type Block struct {
	Stmts []Stmt
	Local *Struct
}

func NewBlock() *Block {
	ret := new(Block)
	ret.Local = NewStruct()
	return ret
}

func (self *Block) PrintTo(p printer.Iface) {
	for _, s := range self.Stmts {
		s.PrintTo(p)
	}
}

func (self *Block) S(s ...Stmt) {
	self.Stmts = append(self.Stmts, s...)
}
