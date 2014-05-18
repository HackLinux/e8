package ir1

import (
	"github.com/h8liu/e8/printer"
)

type AssignStmt struct {
	Alloc bool
	V     *Var
	E     Expr
}

func (s *AssignStmt) PrintTo(p printer.Iface) {
	if s.Alloc {
		p.Printf("%s := %s", s.V.Name, s.E.String())
	} else {
		p.Printf("%s = %s", s.V.Name, s.E.String())
	}
}

