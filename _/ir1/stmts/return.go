package stmts

import (
	"e8vm.net/p/printer"
)

type Return int

const Ret Return = 0

func (s Return) PrintTo(p printer.Iface) {
	p.Printf("return")
}
