package ast

type ConstExpr struct {
	// TODO: a constant token
}

type CallExpr struct {
	Func  Node
	Paras []Node
}
