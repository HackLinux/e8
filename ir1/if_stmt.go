package ir1

import (
	"github.com/h8liu/e8/printer"
)

type IfStmt struct {
	v     *Var
	label string
}

func If(v *Var, lab string) *IfStmt {
	assert(v.Type == Bool)
	return &IfStmt{v, lab}
}

func Goto(lab string) *IfStmt {
	return &IfStmt{nil, lab}
}

func (s IfStmt) PrintTo(p printer.Iface) {
	if s.v == nil {
		p.Printf("goto %s", s.label)
	} else {
		p.Printf("if %s goto %s", s.v.Name, s.label)
	}
}
