package ir1

import (
	"fmt"
)

type BinExpr struct {
	V1, V2 *Var
	Op     Op
}

func (self *BinExpr) Type() Type {
	if self.Op >= OpEq {
		return Bool
	}

	return self.V1.Type
}

func (self *BinExpr) String() string {
	return fmt.Sprintf("%s %s %s",
		self.V1.Name,
		self.Op.String(),
		self.V2.Name,
	)
}

func Bexpr(v1 *Var, op Op, v2 *Var) *BinExpr {
	t1 := v1.Type
	t2 := v2.Type

	switch op {
	case OpNot:
		assert(false)
	case OpAnd, OpOr:
		assert(t1 == Bool && t2 == Bool)
	default:
		assert(t1 == t2)
		assert(t1 != Bool && t2 != Bool)
	}

	return &BinExpr{v1, v2, op}
}
