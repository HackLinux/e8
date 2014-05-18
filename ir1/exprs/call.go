package exprs

import (
	"bytes"
	"fmt"

	"github.com/h8liu/e8/ir1/decls"
	"github.com/h8liu/e8/ir1/types"
	"github.com/h8liu/e8/ir1/vars"
)

type CallExpr struct {
	f    decls.Func
	args []*vars.Var
}

func Call(f decls.Func, args ...*vars.Var) *CallExpr {
	return &CallExpr{f, args}
}

func (self *CallExpr) Type() types.Type {
	return self.f.Type()
}

func (self *CallExpr) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s(", self.f.Name())
	for i, arg := range self.args {
		if i > 0 {
			fmt.Fprint(buf, ", ")
		}
		fmt.Fprint(buf, arg.Name)
	}
	fmt.Fprint(buf, ")")

	return buf.String()
}
