package lexer

import (
	"github.com/h8liu/e8/stay/token"
)

type LexToken struct {
	Token token.Token
	Line  int
	Col   int
	Lit   string
}
