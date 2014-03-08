package parser

import (
	"io"
	"os"
	"strings"

	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/lexer"
	"github.com/h8liu/e8/stay/reporters"
)

type Parser struct {
	// the error reporter on the entire parsing process
	ErrReporter reporters.ErrReporter

	// token position will all be offset by this value on parsing
	PosOffset uint32
}

func New() *Parser {
	ret := new(Parser)
	ret.ErrReporter = reporters.Simple

	return ret
}

func (self *Parser) Parse(in io.Reader) (*ast.Ast, error) {
	lex := lexer.New(in)
	pipe := make(chan *Token, 1)

	go func() {
		for lex.Scan() {
			tok, pos, lit := lex.Token()
			pos += self.PosOffset
			pipe <- &Token{tok, pos, lit}
		}
		close(pipe)
	}()

	for _ = range pipe {
		// TODO: create the ast here
	}

	// wrap up the errors

	// fatal IO error on reading
	e := lex.ScanErr()
	if e != nil {
		return nil, e
	}

	// lexing error on parsing tokens
	if lex.FirstFail != nil {
		return nil, lex.FirstFail
	}

	// TODO: parsing error on creating ast

	return nil, nil
}

func ParseFile(path string) (*ast.Ast, error) {
	fin, e := os.Open(path)
	if e != nil {
		return nil, e
	}

	parser := New()
	parser.ErrReporter = reporters.NewPrefix(path)

	ret, e := parser.Parse(fin)
	if e != nil {
		return nil, e
	}

	e = fin.Close()
	if e != nil {
		return nil, e
	}

	return ret, nil
}

func ParseStr(s string) (*ast.Ast, error) {
	p := New()
	return p.Parse(strings.NewReader(s))
}
