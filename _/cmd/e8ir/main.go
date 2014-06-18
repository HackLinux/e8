package main

import (
	. "e8vm.net/p/ir1"
	. "e8vm.net/p/ir1/exprs"
	. "e8vm.net/p/ir1/ops"
	. "e8vm.net/p/ir1/types"

	"e8vm.net/p/prt"
)

func main() {
	p := NewPackage("main")

	f := p.NewFunc("fabo", U32)
	f.AddArg("i", U32)
	f.Comment("calculating fabonaci numbers")

	c1 := f.NewConst(1, U32)
	c2 := f.NewConst(2, U32)
	t1 := f.NewTemp(f.Binary("i", G, c1))

	f.If(t1, "recur")
	f.AssignReturn(C(1, U32))
	f.Return()

	f.Label("recur")
	t2 := f.NewTemp(f.Binary("i", Sub, c1))
	f.Assign(t2, f.Call("fabo", t2))
	t3 := f.NewTemp(f.Binary("i", Sub, c2))
	f.Assign(t3, f.Call("fabo", t3))
	f.AssignReturn(f.Binary(t2, Add, t3))
	f.Return()

	prt.Print(f)
}
