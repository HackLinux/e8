package lexer

import (
	"github.com/h8liu/e8/leaf/token"
)

type Token struct {
	Token token.Token
	Line  int
	Col   int
	Lit   string
}
