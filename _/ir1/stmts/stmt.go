package stmts

import (
	"e8vm.net/e8/printer"
)

type Stmt interface {
	PrintTo(p printer.Iface)
}
