package stmts

import (
	"e8vm.net/p/printer"
)

type Stmt interface {
	PrintTo(p printer.Iface)
}
