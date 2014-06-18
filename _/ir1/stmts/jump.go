package stmts

import (
	"e8vm.net/e8/ir1/types"
	"e8vm.net/e8/ir1/vars"
	"e8vm.net/e8/printer"
)

type Jump struct {
	V     *vars.Var
	Label string
}

func If(v *vars.Var, lab string) *Jump {
	assert(v.Type == types.Bool)
	return &Jump{v, lab}
}

func Goto(lab string) *Jump {
	return &Jump{nil, lab}
}

func (s *Jump) PrintTo(p printer.Iface) {
	if s.V == nil {
		p.Printf("goto %s", s.Label)
	} else {
		p.Printf("if %s {", s.V.Name)
		p.ShiftIn()
		p.Printf("goto %s", s.Label)
		p.ShiftOut("}")
	}
}
