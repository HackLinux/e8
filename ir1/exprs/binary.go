package exprs

import (
	"github.com/h8liu/e8/ir1/ops"
	"github.com/h8liu/e8/ir1/types"
	"github.com/h8liu/e8/ir1/vars"

	"fmt"
)

type Binary struct {
	Op     ops.Op
	V1, V2 *vars.Var
}

func (self *Binary) Type() types.Type {
	if self.Op >= ops.Eq {
		return types.Bool
	}

	return self.V1.Type
}

func (self *Binary) String() string {
	return fmt.Sprintf("%s %s %s",
		self.V1.Name,
		self.Op.String(),
		self.V2.Name,
	)
}

func NewBinary(v1 *vars.Var, op ops.Op, v2 *vars.Var) *Binary {
	t1 := v1.Type
	t2 := v2.Type

	switch op {
	case ops.Not:
		assert(false)
	case ops.And, ops.Or:
		assert(t1 == types.Bool && t2 == types.Bool)
	default:
		assert(t1 == t2)
		assert(t1 != types.Bool && t2 != types.Bool)
	}

	return &Binary{op, v1, v2}
}
