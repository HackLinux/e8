package stmts

import (
	"github.com/h8liu/e8/printer"
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
