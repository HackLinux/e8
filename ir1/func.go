package ir1

import (
	"github.com/h8liu/e8/printer"
)

type Func struct {
	Name   string  // the function name
	Arg    *Struct // structure of func call arguments
	Ret    *Struct // structure of return values
	*Block         // code block
}

func NewFunc(n string) *Func {
	ret := new(Func)
	ret.Name = n
	ret.Arg = NewStruct()
	ret.Ret = NewStruct()
	ret.Block = NewBlock()

	return ret
}

func F(n string) *Func {
	return NewFunc(n)
}

func (self *Func) PrintTo(p printer.Iface) {
	p.Printf("func %s {", self.Name)
	p.ShiftIn()

	p.Printf("arg {")
	p.ShiftIn()
	self.Arg.PrintTo(p)
	p.ShiftOut("}")

	p.Printf("ret {")
	p.ShiftIn()
	self.Ret.PrintTo(p)
	p.ShiftOut("}")

	p.Printf("code {")
	p.ShiftIn()
	self.Block.PrintTo(p)
	p.ShiftOut("}")

	p.ShiftOut("}")
}
