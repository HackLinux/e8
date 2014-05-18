package ir1

import (
	"github.com/h8liu/e8/printer"
)

type LabelStmt struct {
	label string
}

func Label(n string) *LabelStmt {
	return &LabelStmt{n}
}

func (s *LabelStmt) PrintTo(p printer.Iface) {
	p.ShiftOut()
	p.Printf("%s:", s.label)
	p.ShiftIn()
}
