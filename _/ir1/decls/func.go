package decls

import (
	"e8vm.net/e8/ir1/types"
)

type Func interface {
	Name() string
	Type() types.Type
}
