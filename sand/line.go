package sand

import (
	"github.com/h8liu/e8/printer"
	// "io"
)

type Line interface {
	printer.Printable
	Compile(w *Writer)

	/*
		Count() int // count instructions needs
		ApplyLabel()
	*/
}
