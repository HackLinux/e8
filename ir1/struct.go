package ir1

import (
	"github.com/h8liu/e8/printer"
)

// a memory structure
type Struct struct {
	vars    []*Var
	nameMap map[string]*Var
}

func NewStruct() *Struct {
	ret := new(Struct)
	ret.nameMap = make(map[string]*Var)
	return ret
}

func (self *Struct) PrintTo(p printer.Iface) {
	for _, v := range self.vars {
		p.Print(v.String())
	}
}

func (self *Struct) AddVar(v *Var) bool {
	self.vars = append(self.vars, v)
	if v.Name == "_" {
		return true
	}

	ret := self.nameMap[v.Name]
	if ret != nil {
		self.nameMap[v.Name] = v
		return true
	}
	return false
}

// add a field variable into the structure
func (self *Struct) Fv(v *Var) bool {
	return self.AddVar(v)
}

func (self *Struct) F(n string, t Type) bool {
	return self.Fv(V(n, t))
}

func (self *Struct) Find(n string) *Var {
	if n == "_" {
		return nil
	}

	return self.nameMap[n]
}
