package exprs

import (
	"github.com/h8liu/e8/ir1/types"
	"github.com/h8liu/e8/ir1/vars"
)

type SingleVar struct {
	*vars.Var
}

func NewSingleVar(v *vars.Var) *SingleVar {
	if v == nil {
		panic("bug")
	}

	return &SingleVar{v}
}

func (self *SingleVar) Type() types.Type {
	return self.Var.Type
}

func (self *SingleVar) String() string {
	return self.Var.Name
}
