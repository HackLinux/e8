package ast

type Expr interface {
}

type CallExpr struct {
	Func Expr
}

type IdentExpr struct {
	Ident string
}

type StringLit struct {
	Value string
}

type IntLit struct {
	Type  int
	Value int64
}
