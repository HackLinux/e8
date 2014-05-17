package main

import (
	"github.com/h8liu/e8/ir1"
	"github.com/h8liu/e8/printer"
)

func main() {
	f := ir1.F("fabo")
	f.Arg("i", ir1.U32)
	f.Ret("ret", ir1.U32)
	f.Cm("some comment")
	f.Al("t", f.Vexpr("i"))

	printer.Print(f)
}
