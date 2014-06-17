package parser

import (
	"github.com/h8liu/e8/leaf/lexer"
)

type Error struct {
	Pos *lexer.Token
	Err string
}
