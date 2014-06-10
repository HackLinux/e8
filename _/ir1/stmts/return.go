package stmts

import (
	"github.com/h8liu/e8/printer"
)

type Return int

const Ret Return = 0

func (s Return) PrintTo(p printer.Iface) {
	p.Printf("return")
}
