package exprs

import (
	"bytes"
	"fmt"

	"e8vm.net/p/ir1/decls"
	"e8vm.net/p/ir1/types"
	"e8vm.net/p/ir1/vars"
)

type Call struct {
	F    decls.Func
	Args []*vars.Var
}

func NewCall(f decls.Func, args ...*vars.Var) *Call {
	return &Call{f, args}
}

func (self *Call) Type() types.Type {
	return self.F.Type()
}

func (self *Call) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s(", self.F.Name())
	for i, arg := range self.Args {
		if i > 0 {
			fmt.Fprint(buf, ", ")
		}
		fmt.Fprint(buf, arg.Name)
	}
	fmt.Fprint(buf, ")")

	return buf.String()
}
