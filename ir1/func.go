package ir1

import (
	"fmt"

	"github.com/h8liu/e8/ir1/exprs"
	"github.com/h8liu/e8/ir1/ops"
	"github.com/h8liu/e8/ir1/stmts"
	"github.com/h8liu/e8/ir1/types"
	"github.com/h8liu/e8/ir1/vars"
	"github.com/h8liu/e8/printer"
)

type Func struct {
	name  string    // the function name
	Arg   *Struct   // structure of func call arguments
	Ret   *vars.Var // structure of return values
	Local *Struct   // structure of local variables

	Pack *Package

	Stmts []stmts.Stmt

	nVar int
}

func NewFunc(n string, t types.Type) *Func {
	ret := new(Func)
	ret.name = n
	ret.Arg = NewStruct()
	ret.Ret = vars.NewVar("ret", t)
	ret.Local = NewStruct()
	ret.nVar = 1

	return ret
}

func (self *Func) Name() string     { return self.name }
func (self *Func) Type() types.Type { return self.Ret.Type }

func (self *Func) Sig() string {
	ret := fmt.Sprintf("func %s%s", self.name, self.Arg.List())
	if self.Ret.Type != types.Void {
		ret += fmt.Sprintf(" (ret %s)", self.Ret.Type.String())
	}
	return ret
}

func (self *Func) PrintTo(p printer.Iface) {
	p.Printf("%s {", self.Sig())
	p.ShiftIn()

	for _, s := range self.Stmts {
		s.PrintTo(p)
	}

	p.ShiftOut("}")
}

func (self *Func) AddArg(n string, t types.Type) *vars.Var {
	if !self.Local.Empty() {
		panic("already added local")
	}
	if n == "<ret>" {
		panic("<ret> is reserved for return")
	}
	if t == types.Void || n == "_" {
		panic("bug")
	}

	return self.Arg.Field(n, t)
}

func (self *Func) newLocal(n string, t types.Type) *vars.Var {
	if self.Arg.Find(n) != nil {
		panic("already added in arg")
	}
	if n == "<ret>" {
		panic("<ret> is reserved for return")
	}

	if t == types.Void {
		if n != "_" {
			panic("cannot add named void type")
		}
		return nil
	}

	return self.Local.Field(n, t)
}

// Find a variable in the function scope
// return nil on not found
func (self *Func) V(n string) *vars.Var {
	if n == "_" {
		return nil
	}

	if n == "<ret>" {
		return self.Ret
	}

	v := self.Arg.Find(n)
	if v != nil {
		return v
	}

	v = self.Local.Find(n)
	if v != nil {
		return v
	}

	return nil
}

// Append statements
func (self *Func) State(s ...stmts.Stmt) {
	self.Stmts = append(self.Stmts, s...)
}

// Append a comment statement
func (self *Func) Comment(s string) {
	self.State(stmts.Comment(s))
}

func (self *Func) AssignNew(n string, e exprs.Expr) string {
	v := self.newLocal(n, e.Type())
	as := &stmts.Assign{
		Alloc: true,
		V:     v,
		E:     e,
	}

	self.State(as)

	return n
}

func (self *Func) AssignNewTemp(e exprs.Expr) (n string) {
	for {
		n = fmt.Sprintf("_%d", self.nVar)
		if self.V(n) == nil {
			break
		}
		self.nVar++
	}

	return self.AssignNew(n, e)
}

func (self *Func) Assign(n string, e exprs.Expr) {
	v := self.V(n)

	if n != "_" && v == nil {
		panic("variable not found")
	}
	if !types.IsSame(v.Type, e.Type()) {
		panic("wrong assignment type")
	}

	as := &stmts.Assign{
		Alloc: false,
		V:     v,
		E:     e,
	}

	self.State(as)
}

func (self *Func) AssignReturn(e exprs.Expr) {
	self.Assign("<ret>", e)
}

func (self *Func) Return() {
	self.State(stmts.Ret)
}

func (self *Func) Label(n string) {
	self.State(stmts.NewLabel(n))
}

func (self *Func) If(n string, lab string) {
	self.State(stmts.If(self.V(n), lab))
}

func (self *Func) Goto(lab string) {
	self.State(stmts.Goto(lab))
}

func (self *Func) Call(f string, vs ...string) *exprs.Call {
	var args []*vars.Var

	for _, v := range vs {
		args = append(args, self.V(v))
	}

	fd := self.Pack.FindCall(f)
	assert(fd != nil)

	return exprs.NewCall(fd, args...)
}

func (self *Func) Single(n string) *exprs.Single {
	return exprs.NewSingle(self.V(n))
}

func (self *Func) Binary(n1 string, op ops.Op, n2 string) *exprs.Binary {
	return exprs.NewBinary(self.V(n1), op, self.V(n2))
}
