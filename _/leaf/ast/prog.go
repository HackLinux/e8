// Package ast defines the data structures for the abstract syntax tree.
package ast

import (
	"e8vm.net/e8/printer"
)

type Program struct {
	Imports []*ImportDecl
	Funcs   []*FuncDecl
}

func NewProgram() *Program {
	ret := new(Program)
	ret.Imports = make([]*ImportDecl, 0, 32)
	ret.Funcs = make([]*FuncDecl, 0, 1024)

	return ret
}

func (self *Program) AddImport(i *ImportDecl) {
	self.Imports = append(self.Imports, i)
}

func (self *Program) AddFunc(f *FuncDecl) {
	self.Funcs = append(self.Funcs, f)
}

func (self *Program) PrintTo(p printer.Iface) {
	if len(self.Imports) > 0 {
		p.Println("import (")
		p.ShiftIn()

		for _, imp := range self.Imports {
			imp.PrintTo(p)
		}

		p.ShiftOut(")")
	}

	for _, f := range self.Funcs {
		f.PrintTo(p)
	}
}
