package parser

import (
	"github.com/h8liu/e8/stay/lexer"
)

type Tokener interface {
	Scan() bool
	Token() *lexer.Token
}
