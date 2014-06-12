package parser

import (
	"io"

	"github.com/h8liu/e8/leaf/ast"
	"github.com/h8liu/e8/leaf/lexer"
	"github.com/h8liu/e8/leaf/token"
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
	return ret
}

func (self *Parser) Parse() *ast.Program {
	self.lex = lexer.New(self.in, self.filename)

	prog := new(ast.Program)
	prog.Filename = self.filename

	self.parseDecls(prog)

	return prog
}

func (self *Parser) parseDecls(prog *ast.Program) {
	decl := self.lex.Token()

	switch decl.Token {
	case token.Func:

	default:
		// TODO: report error here
		panic("expect declaration")
	}
}
