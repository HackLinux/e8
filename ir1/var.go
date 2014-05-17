package ir1

import (
	"fmt"
)

type Var struct {
	Name string // "_" for anonymous
	Type Type
}

func (self *Var) String() string {
	return fmt.Sprintf("%s %s", self.Name, self.Type)
}

func NewVar(n string, t Type) *Var {
	return &Var{
		Name: n,
		Type: t,
	}
}

func V(n string, t Type) *Var {
	return NewVar(n, t)
}
