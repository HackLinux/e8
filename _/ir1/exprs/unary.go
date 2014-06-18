package exprs

import (
	"e8vm.net/e8/ir1/ops"
	"e8vm.net/e8/ir1/vars"
)

type Unary struct {
	Op ops.Op
	V  *vars.Var
}
