package ir1

import (
	"github.com/h8liu/e8/printer"
)

// a memory structure
type Struct struct {
	Vars []*Var
}

func NewStruct() *Struct {
	ret := new(Struct)
	return ret
}

func (self *Struct) PrintTo(p printer.Interface) {
	for _, v := range self.Vars {
		p.Print(v.String())
	}
}

// add a field variable into the structure
func (self *Struct) Fv(v *Var) {
	self.Vars = append(self.Vars, v)
}

func (self *Struct) F(n string, t Type) {
	self.Fv(V(n, t))
}
