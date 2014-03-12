package sand

import (
	"github.com/h8liu/e8/printer"
)

type Call struct {
	name string // the function name, should be package.type.func
	f    *Func
	args []*Var
}

var _ Line = new(Call)

func (self *Call) AddArg(arg *Var) {
	self.args = append(self.args, arg)
}

func (self *Call) PrintTo(p printer.Interface) {
	name := self.name
	if self.f != nil {
		name = self.f.name
	}

	if len(self.args) > 0 {
		p.Printf("call %s {", name)
		p.ShiftIn()
		for _, arg := range self.args {
			arg.PrintTo(p)
		}
		p.ShiftOut("}")
	} else {
		p.Printf("call %s", name)
	}
}

func (self *Call) Compile(w *Writer) {

}
