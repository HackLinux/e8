package stmts

import (
	"e8vm.net/p/ir1/exprs"
	"e8vm.net/p/ir1/vars"
	"e8vm.net/p/printer"
)

type Assign struct {
	Alloc bool
	V     *vars.Var
	E     exprs.Expr
}

func (s *Assign) PrintTo(p printer.Iface) {
	if s.Alloc {
		p.Printf("%s := %s", s.V.Name, s.E.String())
	} else {
		p.Printf("%s = %s", s.V.Name, s.E.String())
	}
}
