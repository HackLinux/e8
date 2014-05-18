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
	f := p.NewFunc("fabo", U32)
	f.Arg("i", U32)
	f.Comment("calculating fabonaci numbers")

	c1 := f.AssignNew("c1", C(1, U32))
	c2 := f.AssignNew("c2", C(2, U32))
	t1 := f.AssignNewTemp(f.Binary("i", G, c1))

	f.If(t1, "recur")
	f.AssignReturn(C(1, U32))
	f.Return()

	f.Label("recur")
	t2 := f.AssignNewTemp(f.Binary("i", Sub, c1))
	f.Assign(t2, f.Call("fabo", t2))
	t3 := f.AssignNewTemp(f.Binary("i", Sub, c2))
	f.Assign(t3, f.Call("fabo", t3))
	f.AssignReturn(f.Binary(t2, Add, t3))
	f.Return()

	printer.Print(f)
}
