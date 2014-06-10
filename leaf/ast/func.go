package ast

import (
	"github.com/h8liu/e8/printer"
)

type FuncDecl struct {
	Name     string
	DeclLine int // XXX: fix me
	Body     *BlockStmt
}

func (self *FuncDecl) PrintTo(p printer.Iface) {
	p.Printf("func %s() {", self.Name)
	p.ShiftIn()

	for _, stmt := range self.Body.Stmts {
		stmt.PrintTo(p)
	}

	p.ShiftOut("}")
}
