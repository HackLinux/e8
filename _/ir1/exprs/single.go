package exprs

import (
	"e8vm.net/e8/ir1/types"
	"e8vm.net/e8/ir1/vars"
)

type Single struct {
	V *vars.Var
}

func NewSingle(v *vars.Var) *Single {
	if v == nil {
		panic("bug")
	}

	return &Single{v}
}

func (self *Single) Type() types.Type {
	return self.V.Type
}

func (self *Single) String() string {
	return self.V.Name
}
