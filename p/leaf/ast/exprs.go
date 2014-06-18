package ast

import (
	"e8vm.net/p/leaf/lexer"
)

type ConstExpr struct {
	// TODO: a constant token
}

type CallExpr struct {
	Func Node
	Args []Node
}

type Operand struct {
	Token *lexer.Token
}
