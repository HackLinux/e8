package parser

import (
	"io"
	"os"
	"strings"

	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/lexer"
	"github.com/h8liu/e8/stay/tokens"
)

type Parser struct {
}

func New() *Parser {
	ret := new(Parser)

	return ret
}

func (self *Parser) Parse(id uint8, in io.Reader) (*ast.Ast, error) {
	lex := lexer.New(in)

	for {
		to, _, _ := lex.Scan()
		if to == tokens.EOF {
			break
		}
	}

	e := lex.ScanErr()
	if e != nil {
		return nil, e
	}

	if lex.FirstFail != nil {
		return nil, lex.FirstFail
	}

	return nil, nil
}

func ParseFile(path string) (*ast.Ast, error) {
	fin, e := os.Open(path)
	if e != nil {
		return nil, e
	}

	parser := New()
	ret, e := parser.Parse(0, fin)
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
	return p.Parse(0, strings.NewReader(s))
}
