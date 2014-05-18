package exprs

import (
	"github.com/h8liu/e8/ir1/ops"
	"github.com/h8liu/e8/ir1/vars"
)

type Unary struct {
	V  *vars.Var
	Op ops.Op
}
