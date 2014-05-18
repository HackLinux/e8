package ir1

import (
	"bytes"
	"fmt"

	"github.com/h8liu/e8/ir1/types"
	"github.com/h8liu/e8/ir1/vars"

	"github.com/h8liu/e8/printer"
)

// A combined continuous memory area of named memory slots
type Struct struct {
	vars    []*vars.Var
	nameMap map[string]*vars.Var
}

func NewStruct() *Struct {
	ret := new(Struct)
	ret.nameMap = make(map[string]*vars.Var)
	return ret
}

func (self *Struct) PrintTo(p printer.Iface) {
	for _, v := range self.vars {
		p.Print(v.String())
	}
}

func (self *Struct) List() string {
	buf := new(bytes.Buffer)

	fmt.Fprint(buf, "(")
	for i, v := range self.vars {
		if i > 0 {
			fmt.Fprint(buf, ",")
		}
		fmt.Fprintf(buf, "%s %s", v.Name, v.Type)
	}
	fmt.Fprint(buf, ")")
	return buf.String()
}

func (self *Struct) Var(v *vars.Var) {
	self.vars = append(self.vars, v)
	if v.Name == "_" {
		// add padding
		return
	}

	ret := self.nameMap[v.Name]
	if ret != nil {
		panic("duplicated variable in struct")
	}

	self.nameMap[v.Name] = v
}

func (self *Struct) Field(n string, t types.Type) *vars.Var {
	v := vars.NewVar(n, t)
	self.Var(v)
	return v
}

func (self *Struct) Find(n string) *vars.Var {
	if n == "_" {
		return nil
	}

	return self.nameMap[n]
}

func (self *Struct) Empty() bool {
	return len(self.vars) == 0
}
