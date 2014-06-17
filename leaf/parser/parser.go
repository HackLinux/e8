package parser

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/h8liu/e8/leaf/ast"
	"github.com/h8liu/e8/leaf/comperr"
	"github.com/h8liu/e8/leaf/lexer"
	t "github.com/h8liu/e8/leaf/token"
)

type Parser struct {
	filename string
	in       io.Reader

	lex *lexer.Lexer
	*scanner
	errors []error
}

func Parse(f string) (*ast.Program, []error) {
	p, e := Open(f)
	if e != nil {
		return nil, []error{e}
	}

	return p.Parse()
}

type TokenTree interface {
	PrintTree(io.Writer)
}

func ParseTree(f string) (TokenTree, []error) {
	p, e := Open(f)
	if e != nil {
		return nil, []error{e}
	}

	_, errs := p.Parse()
	return p, errs
}

func ParseStr(f string, prog string) (*ast.Program, []error) {
	in := strings.NewReader(prog)
	p := New(in, f)
	return p.Parse()
}

func New(in io.Reader, filename string) *Parser {
	ret := new(Parser)
	ret.filename = filename
	ret.scanner = newScanner(in, filename)
	return ret
}

func Open(filename string) (*Parser, error) {
	fin, e := os.Open(filename)
	if e != nil {
		return nil, e
	}

	return New(fin, filename), nil
}

func (p *Parser) Parse() (*ast.Program, []error) {
	p.push("source-file")
	defer p.pop()

	ret := new(ast.Program)

	for !p.eof() {
		d := p.parseTopDecl()
		if d != nil {
			ret.Decls = append(ret.Decls, d)
		}
	}

	return ret, p.Errors()
}

func (p *Parser) PrintTree(out io.Writer) {
	p.scanner.tracker.PrintTrack(out)
}

func (p *Parser) Errors() []error {
	if p.scanner.errors != nil {
		return p.scanner.errors
	}

	return p.errors
}

func (p *Parser) parseTopDecl() ast.Node {
	if p.ahead(t.Func) {
		return p.parseFunc()
	}

	p.parseErrorDecl()
	return nil
}

func (p *Parser) expect(tok t.Token) bool {
	if tok == t.EOF {
		panic("cannot expect EOF")
	}

	if p.ahead(tok) {
		assert(p.shift())
		return true
	}

	p.err(fmt.Sprintf("expect %s, got %s", tok, p.cur.Token))

	p.shift() // make progress anyway
	return false
}

func (p *Parser) parseFunc() *ast.Func {
	p.push("func-decl")
	defer p.pop()

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

	if !p.expect(t.Semi) {
		return err()
	}

	// TODO:
	return ret
}

func (p *Parser) err(s string) {
	e := new(comperr.Error)
	e.File = p.filename
	e.Err = fmt.Errorf("%s", s)
	e.Line = p.cur.Line
	e.Col = p.cur.Col

	p.errors = append(p.errors, e)
}

func (p *Parser) parseErrorDecl() {
	p.push("error-decl")
	defer p.pop()

	p.err("syntax error: expect declaration")
	p.skipUntil(t.Semi)
}
