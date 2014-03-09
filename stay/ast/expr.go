package ast

type Expr interface {
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
	ret.ArgList = make([]Expr, 8)
	return ret
}

func (self *CallExpr) AddArg(e Expr) {
	self.ArgList = append(self.ArgList, e)
}
