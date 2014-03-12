package sand

import (
	"io"

	"github.com/h8liu/e8/printer"
)

type Func struct {
	localMap map[string]*Var
	local []*Var

	lines []Line
}

type Line interface {
	printer.Printable

	Count() int // count instructions needs
	ApplyLabel()
	Compile(out io.Writer)
}
