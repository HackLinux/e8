package ir1

import (
	"bytes"
	"fmt"
)

type CallExpr struct {
	f    FuncDecl
	args []*Var
}

func Call(f FuncDecl, args ...*Var) *CallExpr {
	return &CallExpr{f, args}
}

func (self *CallExpr) Type() Type {
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
