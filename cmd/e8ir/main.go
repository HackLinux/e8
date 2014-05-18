package main

import (
	. "github.com/h8liu/e8/ir1"
	"github.com/h8liu/e8/printer"
)

func main() {
	f := F("fabo")
	f.Arg("i", U32)
	f.Ret("ret", U32)
	f.Cm("calculating fabonaci numbers")

	c1 := f.Al("c1", C(1, U32))
	c2 := f.Al("c2", C(2, U32))
	_ = f.AlTmp(f.Bexpr("i", OpG, c1))
	// f.If(t1, "recur")
	f.As("ret", C(1, U32))
	f.Return()
	// f.Label("recur")
	t2 := f.AlTmp(f.Bexpr("i", OpSub, c1))
	// f.As(t2, f.Call("fabo", t2))
	t3 := f.AlTmp(f.Bexpr("i", OpSub, c2))
	// f.As(t3, f.Call("fabo", t3))
	f.As("ret", f.Bexpr(t2, OpAdd, t3))
	f.Return()

	printer.Print(f)
}
