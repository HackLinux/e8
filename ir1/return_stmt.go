package ir1

import (
	"github.com/h8liu/e8/printer"
)

type ReturnStmt int

const Return ReturnStmt = 0

func (s ReturnStmt) PrintTo(p printer.Iface) {
	p.Printf("return")
}
