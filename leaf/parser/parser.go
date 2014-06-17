package parser

import (
	"io"

	"github.com/h8liu/e8/leaf/ast"
	"github.com/h8liu/e8/leaf/lexer"
	t "github.com/h8liu/e8/leaf/token"
)

type Parser struct {
	filename string
	in       io.Reader

	lex *lexer.Lexer
	*scanner
}

func NewParser(in io.Reader, filename string) *Parser {
	ret := new(Parser)
	ret.filename = filename
	ret.scanner = newScanner(in, filename)
	return ret
}

func (p *Parser) Parse() *ast.Program {
	ret := new(ast.Program)

	for !p.eof() {
		d := p.parseTopDecl()
		ret.AddDecl(d)
	}

	return ret
}

func (p *Parser) parseTopDecl() ast.Node {
	if p.ahead(t.Func) {
		return p.parseFunc()
	}

	return p.parseErrorDecl()
}

func (p *Parser) parseFunc() *ast.Func {
	ret := new(ast.Func)
	err := func() *ast.Func {
		p.skipUntil(t.Semi)
		return ret
	}

	if !p.expect(t.Func) {
		return err()
	}

	if !p.expect(t.Ident) {
		return err()
	}

	ret.Name = p.last.Lit

	// TODO: parse args and signature
	if !p.expect(t.Lparen) {
		return err()
	}

	if !p.expect(t.Rparen) {
		return err()
	}

	// TODO: parse block
	if !p.expect(t.Lbrace) {
		return err()
	}

	if !p.expect(t.Rbrace) {
		return err()
	}

	// TODO:
	return ret
}

func (p *Parser) parseErrorDecl() *ast.Error {
	ret := ast.NewError("invalid declaration")

	skipped := p.skipUntil(t.Semi)
	for _, n := range skipped {
		ret.Add(n)
	}

	return ret
}
