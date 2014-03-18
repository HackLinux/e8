package ast

import (
	"github.com/h8liu/e8/printer"
)

type ImportDecl struct {
	As   string
	Path string
	Line int
}

func (self *ImportDecl) PrintTo(p printer.Interface) {
	if self.As == "" {
		p.Printf("%q\n", self.Path)
	} else {
		p.Printf("%s %q\n", self.As, self.Path)
	}
}
