package exprs

import (
	"e8vm.net/e8/ir1/types"
)

type Expr interface {
	Type() types.Type
	String() string
}
