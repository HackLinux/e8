package main

import (
	"github.com/h8liu/e8/ir1"
	"github.com/h8liu/e8/printer"
)

func main() {
	f := ir1.F("fabo")
	f.Arg.F("i", ir1.U32)
	f.Ret.F("ret", ir1.U32)
	f.S(ir1.Cm("some comment"))
	f.Local.F("x", ir1.U32)
	f.Al("t", f.Vexpr("x"))

	printer.Print(f)
}
