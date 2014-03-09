package ast

import (
	"github.com/h8liu/e8/printer"
)

type FuncDecl struct {
	Name    string
	DeclPos uint32
	Body    *BlockStmt
}

func (self *FuncDecl) PrintTo(p printer.Interface) {
	p.Printf("func %s() {", self.Name)
	p.ShiftIn()

	for _, stmt := range self.Body.Stmts {
		stmt.PrintTo(p)
	}

	p.ShiftOut("}")
}
