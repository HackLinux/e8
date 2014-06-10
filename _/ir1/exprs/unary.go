package exprs

import (
	"github.com/h8liu/e8/ir1/ops"
	"github.com/h8liu/e8/ir1/vars"
)

type Unary struct {
	Op ops.Op
	V  *vars.Var
}
