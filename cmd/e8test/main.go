package main

import (
	"os"

	"github.com/h8liu/e8/printer"
	"github.com/h8liu/e8/sand"
)

func main() {
	f := sand.NewFunc("test")
	v := f.NewVar(sand.TypUint8, 1)

	c := f.NewCall("print", nil)
	c.AddArg(v)

	p := printer.New(os.Stdout)
	f.PrintTo(p)
}
