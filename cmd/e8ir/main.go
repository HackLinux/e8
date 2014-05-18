package main

import (
	. "github.com/h8liu/e8/ir1"
	. "github.com/h8liu/e8/ir1/exprs"
	. "github.com/h8liu/e8/ir1/ops"
	. "github.com/h8liu/e8/ir1/types"

	"github.com/h8liu/e8/printer"
)

func main() {
	p := P("main")
	f := p.F("fabo", U32)
	f.Arg("i", U32)
	f.Cm("calculating fabonaci numbers")

	c1 := f.Al("c1", C(1, U32))
	c2 := f.Al("c2", C(2, U32))
	t1 := f.AlTmp(f.Bexpr("i", G, c1))
	f.If(t1, "recur")
	f.RetAs(C(1, U32))
	f.Return()
	f.Label("recur")
	t2 := f.AlTmp(f.Bexpr("i", Sub, c1))
	f.As(t2, f.Call("fabo", t2))
	t3 := f.AlTmp(f.Bexpr("i", Sub, c2))
	f.As(t3, f.Call("fabo", t3))
	f.RetAs(f.Bexpr(t2, Add, t3))
	f.Return()

	printer.Print(f)
}
