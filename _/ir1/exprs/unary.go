package exprs

import (
	"e8vm.net/p/ir1/ops"
	"e8vm.net/p/ir1/vars"
)

type Unary struct {
	Op ops.Op
	V  *vars.Var
}
