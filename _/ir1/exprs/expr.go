package exprs

import (
	"e8vm.net/p/ir1/types"
)

type Expr interface {
	Type() types.Type
	String() string
}
