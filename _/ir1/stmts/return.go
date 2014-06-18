package stmts

import (
	"e8vm.net/e8/printer"
)

type Return int

const Ret Return = 0

func (s Return) PrintTo(p printer.Iface) {
	p.Printf("return")
}
