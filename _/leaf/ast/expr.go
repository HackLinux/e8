package ast

import (
	"e8vm.net/p/printer"
)

type Expr interface {
	Stmt
}

type Ident struct {
	Ident string
}

type StringLit struct {
	Value string
}

type IntLit struct {
	Value int64
}

type FloatLit struct {
	Value float64
}

type CharLit struct {
	R uint8
}

type ParenExpr struct {
	X Expr
}

type BadExpr struct {
}

type CallExpr struct {
	Func    Expr
	ArgList []Expr
}

func NewCallExpr() *CallExpr {
	ret := new(CallExpr)
	ret.ArgList = make([]Expr, 0, 8)
	return ret
}

func (self *CallExpr) AddArg(e Expr) {
	self.ArgList = append(self.ArgList, e)
}

func (self *Ident) PrintTo(p printer.Iface) {
	p.Println("ident: ", self.Ident)
}

func (self *StringLit) PrintTo(p printer.Iface) {
	p.Printf("string: %q", self.Value)
}

func (self *IntLit) PrintTo(p printer.Iface) {
	p.Printf("int: %d", self.Value)
}

func (self *FloatLit) PrintTo(p printer.Iface) {
	p.Printf("float: %q", self.Value)
}

func (self *CharLit) PrintTo(p printer.Iface) {
	p.Printf("char: %q", rune(self.R))
}

func (self *ParenExpr) PrintTo(p printer.Iface) {
	p.Println("(")
	p.ShiftIn()
	self.X.PrintTo(p)
	p.ShiftOut(")")
}

func (self *BadExpr) PrintTo(p printer.Iface) {
	p.Println("BAD")
}

func (self *CallExpr) PrintTo(p printer.Iface) {
	p.Println("call {")
	p.ShiftIn()

	if ident, ok := self.Func.(*Ident); ok {
		p.Printf("func %s", ident.Ident)
	} else {
		p.Println("func {")
		self.Func.PrintTo(p)
		p.ShiftIn()
		self.Func.PrintTo(p)
		p.ShiftOut("}")
	}

	p.Println("args {")
	p.ShiftIn()
	for i, arg := range self.ArgList {
		if i > 0 {
			p.Println(",")
		}
		arg.PrintTo(p)
	}
	p.ShiftOut("}")

	p.ShiftOut("}")
}
