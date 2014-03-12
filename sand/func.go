package sand

import (
	"fmt"

	"github.com/h8liu/e8/printer"
)

type Func struct {
	name   string
	args   []*Var
	locals []*Var
	lines  []Line
}

func NewFunc(n string) *Func {
	ret := new(Func)
	ret.name = n
	ret.locals = make([]*Var, 0, 1024)
	ret.lines = make([]Line, 0, 1024)
	return ret
}

func (self *Func) NewVar(t, size int) *Var {
	n := len(self.locals)
	v := new(Var)
	v.typ = t
	v.size = size
	v.name = fmt.Sprintf("t%d", n)

	self.locals = append(self.locals, v)
	return v
}

func (self *Func) NewCall(name string, f *Func) *Call {
	call := new(Call)
	call.name = name
	call.f = f
	call.args = make([]*Var, 0, 8)

	self.lines = append(self.lines, call)
	return call
}

func (self *Func) PrintTo(p printer.Interface) {
	p.Printf("func %s() {", self.name)
	p.ShiftIn()

	if len(self.args) > 0 {
		p.Printf("arg {")
		p.ShiftIn()
		for _, v := range self.args {
			v.PrintTo(p)
		}
		p.ShiftOut("}")
	}

	if len(self.locals) > 0 {
		p.Printf("locals {")
		p.ShiftIn()
		for _, v := range self.locals {
			v.PrintTo(p)
		}
		p.ShiftOut("}")
	}

	for _, line := range self.lines {
		line.PrintTo(p)
	}

	p.ShiftOut("}")
}
