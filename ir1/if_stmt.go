package ir1

import (
	"github.com/h8liu/e8/ir1/types"
	"github.com/h8liu/e8/ir1/vars"
	"github.com/h8liu/e8/printer"
)

type IfStmt struct {
	v     *vars.Var
	label string
}

func If(v *vars.Var, lab string) *IfStmt {
	assert(v.Type == types.Bool)
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
