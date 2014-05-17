package ast

import (
	"github.com/h8liu/e8/printer"
)

type ImportDecl struct {
	As   string
	Path string
	Line int
}

func (self *ImportDecl) PrintTo(p printer.Iface) {
	if self.As == "" {
		p.Printf("%q", self.Path)
	} else {
		p.Printf("%s %q", self.As, self.Path)
	}
}
