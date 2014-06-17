package ast

import (
	"github.com/h8liu/e8/leaf/lexer"
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
