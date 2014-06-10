package decls

import (
	"github.com/h8liu/e8/ir1/types"
)

type Func interface {
	Name() string
	Type() types.Type
}
