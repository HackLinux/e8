package stmts

import (
	"github.com/h8liu/e8/ir1/exprs"
	"github.com/h8liu/e8/ir1/vars"
	"github.com/h8liu/e8/printer"
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
