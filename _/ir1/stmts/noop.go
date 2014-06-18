package stmts

import (
	"e8vm.net/p/printer"
)

type Noop struct {
	Comment string
}

func Comment(c string) *Noop {
	ret := new(Noop)
	ret.Comment = c
	return ret
}

func (s *Noop) PrintTo(p printer.Iface) {
	if s.Comment != "" {
		p.Printf("// %s", s.Comment)
	}
}
