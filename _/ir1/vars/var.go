package vars

import (
	"e8vm.net/e8/ir1/types"

	"fmt"
)

type Var struct {
	Name string // "_" for anonymous
	Type types.Type
}

func (self *Var) String() string {
	return fmt.Sprintf("%s %s", self.Name, self.Type)
}

func NewVar(n string, t types.Type) *Var {
	return &Var{
		Name: n,
		Type: t,
	}
}
