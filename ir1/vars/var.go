package vars

import (
	"github.com/h8liu/e8/ir1/types"

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

func V(n string, t types.Type) *Var {
	return NewVar(n, t)
}
