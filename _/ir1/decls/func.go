package decls

import (
	"e8vm.net/p/ir1/types"
)

type Func interface {
	Name() string
	Type() types.Type
}
