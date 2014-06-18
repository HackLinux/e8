package exprs

import (
	"e8vm.net/e8/ir1/types"
)

import (
	"fmt"
)

type Const struct {
	V int64
	T types.Basic
}

func (self *Const) Type() types.Type {
	return self.T
}

func (self *Const) String() string {
	return fmt.Sprintf("%s(%d)", self.T, self.V)
}

func C(v int64, t types.Basic) *Const {
	switch t {
	case types.Bool:
		if v != 0 {
			v = 1
		}
	case types.U8:
		v = int64(uint8(v))
	case types.I8:
		v = int64(int8(v))
	case types.U16:
		v = int64(uint16(v))
	case types.I16:
		v = int64(int16(v))
	case types.U32:
		v = int64(uint32(v))
	case types.I32:
		v = int64(int32(v))
	default:
		panic("unknown basic value")
	}

	return &Const{v, t}
}
