package ir1

import (
	"fmt"

	"github.com/h8liu/e8/printer"
)

type Func struct {
	name  string  // the function name
	arg   *Struct // structure of func call arguments
	ret   *Var    // structure of return values
	local *Struct // structure of local variables

	pack *Package

	nvar int

	Stmts []Stmt
}

func NewFunc(n string, t Type) *Func {
	ret := new(Func)
	ret.name = n
	ret.arg = NewStruct()
	ret.ret = &Var{"<ret>", t}
	ret.local = NewStruct()

	return ret
}

func F(n string, t Type) *Func {
	return NewFunc(n, t)
}

func (self *Func) Name() string { return self.name }
func (self *Func) Type() Type   { return self.ret.Type }

func (self *Func) PrintTo(p printer.Iface) {
	p.Printf("func %s %s {", self.name, self.ret.Type.String())
	p.ShiftIn()

	p.Printf("arg {")
	p.ShiftIn()
	self.arg.PrintTo(p)
	p.ShiftOut("}")

	p.Printf("code {")
	p.ShiftIn()
	for _, s := range self.Stmts {
		s.PrintTo(p)
	}
	p.ShiftOut("}")

	p.ShiftOut("}")
}

func (self *Func) Arg(n string, t Type) *Var {
	if !self.local.Empty() {
		panic("already added local")
	}
	if n == "<ret>" {
		panic("<ret> is reserved for return")
	}
	if t == Void || n == "_" {
		panic("bug")
	}

	return self.arg.F(n, t)
}

func (self *Func) Var(n string, t Type) *Var {
	if self.arg.Find(n) != nil {
		panic("already added in arg")
	}
	if n == "<ret>" {
		panic("<ret> is reserved for return")
	}

	if t == Void {
		if n != "_" {
			panic("cannot add named void type")
		}
		return nil
	}

	return self.local.F(n, t)
}

// Find a variable in the function scope
// return nil on not found
func (self *Func) Find(n string) *Var {
	if n == "_" {
		return nil
	}

	if n == "<ret>" {
		return self.ret
	}

	v := self.arg.Find(n)
	if v != nil {
		return v
	}

	v = self.local.Find(n)
	if v != nil {
		return v
	}

	return nil
}

// Append statements
func (self *Func) S(s ...Stmt) {
	self.Stmts = append(self.Stmts, s...)
}

// Append a comment statement
func (self *Func) Cm(s string) {
	self.S(Cm(s))
}

func (self *Func) Al(n string, e Expr) string {
	v := self.Var(n, e.Type())
	as := &AssignStmt{
		Alloc: true,
		V:     v,
		E:     e,
	}

	self.S(as)

	return n
}

func (self *Func) AlTmp(e Expr) (n string) {
	for {
		n = fmt.Sprintf("_%d", self.nvar)
		if self.Find(n) == nil {
			break
		}
		self.nvar++
	}

	return self.Al(n, e)
}

func (self *Func) As(n string, e Expr) {
	v := self.Find(n)

	if n != "_" && v == nil {
		panic("variable not found")
	}
	if !SameType(v.Type, e.Type()) {
		panic("wrong assignment type")
	}

	as := &AssignStmt{
		Alloc: false,
		V:     v,
		E:     e,
	}

	self.S(as)
}

func (self *Func) RetAs(e Expr) {
	self.As("<ret>", e)
}

func (self *Func) Return() {
	self.S(Return)
}

func (self *Func) Label(n string) {
	self.S(Label(n))
}

func (self *Func) If(n string, lab string) {
	self.S(If(self.V(n), lab))
}

func (self *Func) Goto(lab string) {
	self.S(Goto(lab))
}

func (self *Func) Call(f string, vars ...string) *CallExpr {
	var args []*Var

	for _, v := range vars {
		args = append(args, self.V(v))
	}

	fd := self.pack.FindCall(f)
	assert(fd != nil)

	return Call(fd, args...)
}

func (self *Func) V(n string) *Var {
	return self.Find(n)
}

func (self *Func) Vexpr(n string) *VarExpr {
	return Vexpr(self.V(n))
}

func (self *Func) Bexpr(n1 string, op Op, n2 string) *BinExpr {
	return Bexpr(self.V(n1), op, self.V(n2))
}
