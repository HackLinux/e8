package stmts

import (
	"e8vm.net/p/printer"
)

type Label struct {
	label string
}

func NewLabel(n string) *Label {
	return &Label{n}
}

func (s *Label) PrintTo(p printer.Iface) {
	p.ShiftOut()
	p.Printf("%s:", s.label)
	p.ShiftIn()
}
