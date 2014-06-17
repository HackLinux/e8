package parser

import (
	"io"

	// "github.com/h8liu/e8/leaf/ast"
	"github.com/h8liu/e8/leaf/lexer"
	// t "github.com/h8liu/e8/leaf/token"
)

type Parser struct {
	filename string
	in       io.Reader

	lex *lexer.Lexer
}

func NewParser(in io.Reader, filename string) *Parser {
	ret := new(Parser)
	ret.in = in
	ret.filename = filename
	ret.lex = lexer.New(in, filename)
	return ret
}
