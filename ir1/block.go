package ir1

import (
	"github.com/h8liu/e8/printer"
)

type Block struct {
	Stmts []Stmt
	Local *Struct
}

func NewBlock() *Block {
	ret := new(Block)
	ret.Local = NewStruct()
	return ret
}

func (self *Block) PrintTo(p printer.Iface) {
	for _, s := range self.Stmts {
		s.PrintTo(p)
	}
}

func (self *Block) S(s ...Stmt) {
	self.Stmts = append(self.Stmts, s...)
}

func (self *Block) Cm(s string) {
	self.S(Cm(s))
}

func (self *Block) Al(n string, e Expr) {
	// TODO: first check in arg and ret for existense
	v := self.Local.F(n, e.Type())
	as := &AssignStmt{
		Alloc: true,
		V:     v,
		E:     e,
	}

	self.S(as)
}

func (self *Block) As(n string, e Expr) {
	// TODO: also find in arg and ret of the func
	v := self.Local.Find(n)

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

func (self *Block) V(n string) *Var {
	return self.Local.Find(n)
}

func (self *Block) Vexpr(n string) *VarExpr {
	return Ve(self.V(n))
}
