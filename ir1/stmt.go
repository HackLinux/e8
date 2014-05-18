package ir1

import (
	"github.com/h8liu/e8/printer"
)

type Stmt interface {
	PrintTo(p printer.Iface)
}
