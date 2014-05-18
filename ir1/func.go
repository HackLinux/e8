package ir1

import (
	"fmt"

	"github.com/h8liu/e8/printer"
)

type Func struct {
	Name  string  // the function name
	arg   *Struct // structure of func call arguments
	ret   *Struct // structure of return values
	local *Struct // structure of local variables

	nvar int

	Stmts []Stmt
}

func NewFunc(n string) *Func {
	ret := new(Func)
	ret.Name = n
	ret.arg = NewStruct()
	ret.ret = NewStruct()
	ret.local = NewStruct()

	return ret
}

func F(n string) *Func {
	return NewFunc(n)
}

func (self *Func) PrintTo(p printer.Iface) {
	p.Printf("func %s {", self.Name)
	p.ShiftIn()

	p.Printf("arg {")
	p.ShiftIn()
	self.arg.PrintTo(p)
	p.ShiftOut("}")

	p.Printf("ret {")
	p.ShiftIn()
	self.ret.PrintTo(p)
	p.ShiftOut("}")

	p.Printf("code {")
	p.ShiftIn()
	for _, s := range self.Stmts {
		s.PrintTo(p)
	}
	p.ShiftOut("}")

	p.ShiftOut("}")
}

func (self *Func) S(s ...Stmt) {
	self.Stmts = append(self.Stmts, s...)
}

func (self *Func) Cm(s string) {
	self.S(Cm(s))
}

func (self *Func) Arg(n string, t Type) *Var {
	if !self.ret.Empty() {
		panic("already added ret")
	}
	if !self.local.Empty() {
		panic("already added local")
	}

	return self.arg.F(n, t)
}

func (self *Func) Ret(n string, t Type) *Var {
	if self.arg.Find(n) != nil {
		panic("already added in arg")
	}
	if !self.local.Empty() {
		panic("already added local")
	}

	return self.ret.F(n, t)
}

func (self *Func) Var(n string, t Type) *Var {
	if self.arg.Find(n) != nil {
		panic("already added in arg")
	}
	if self.ret.Find(n) != nil {
		panic("already added in ret")
	}

	return self.local.F(n, t)
}

// Find a variable in the function scope
// return nil on not found
func (self *Func) Find(n string) *Var {
	if n == "_" {
		return nil
	}

	v := self.arg.Find(n)
	if v != nil {
		return v
	}

	v = self.ret.Find(n)
	if v != nil {
		return v
	}

	v = self.local.Find(n)
	if v != nil {
		return v
	}

	return nil
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

func (self *Func) V(n string) *Var {
	return self.Find(n)
}

func (self *Func) Vexpr(n string) *VarExpr {
	return Vexpr(self.V(n))
}

func (self *Func) Bexpr(n1 string, op Op, n2 string) *BinExpr {
	return Bexpr(self.V(n1), op, self.V(n2))
}

func (self *Func) Return() {
	self.S(Return)
}
