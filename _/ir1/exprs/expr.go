package exprs

import (
	"github.com/h8liu/e8/ir1/types"
)

type Expr interface {
	Type() types.Type
	String() string
}
