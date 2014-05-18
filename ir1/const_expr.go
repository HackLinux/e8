package ir1

import (
	"fmt"
)

type ConstExpr struct {
	v int64
	t BasicType
}

func (self *ConstExpr) Type() Type {
	return self.t
}

func (self *ConstExpr) String() string {
	return fmt.Sprintf("%s(%d)", self.t.String(), self.v)
}

func C(v int64, t BasicType) *ConstExpr {
	switch t {
	case Bool:
		if v != 0 {
			v = 1
		}
	case U8:
		v = int64(uint8(v))
	case I8:
		v = int64(int8(v))
	case U16:
		v = int64(uint16(v))
	case I16:
		v = int64(int16(v))
	case U32:
		v = int64(uint32(v))
	case I32:
		v = int64(int32(v))
	default:
		panic("unknown basic value")
	}

	return &ConstExpr{v, t}
}
