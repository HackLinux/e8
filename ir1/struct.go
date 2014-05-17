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

func (self *Struct) AddVar(v *Var) {
	self.vars = append(self.vars, v)
	if v.Name == "_" {
		return
	}

	ret := self.nameMap[v.Name]
	if ret != nil {
		panic("duplicated variable in struct")
	}

	self.nameMap[v.Name] = v
}

// add a field variable into the structure
func (self *Struct) Fv(v *Var) {
	self.AddVar(v)
}

func (self *Struct) F(n string, t Type) *Var {
	v := V(n, t)
	self.Fv(v)
	return v
}

func (self *Struct) Find(n string) *Var {
	if n == "_" {
		return nil
	}

	return self.nameMap[n]
}
