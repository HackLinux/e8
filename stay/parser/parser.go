package parser

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/lexer"
	"github.com/h8liu/e8/stay/reporter"
)

const (
	MaxLine  = 50000 // 16 bit
	MaxCol   = 250   // 8 bit
	MaxError = 128
)

type Parser struct {
	// the error reporter on the entire parsing process
	Reporter reporter.Interface

	// token position will all be offset by this value on parsing
	PosOffset uint32

	s      *TokenScanner
	e      error
	prog   *ast.Program
	errors []error
}

func New() *Parser {
	ret := new(Parser)
	ret.Reporter = reporter.Simple
	ret.errors = make([]error, 0, MaxError)

	return ret
}

func (self *Parser) fail(line, col int, e error) {
	self.Reporter.Report(line, col, e)
	if self.e == nil {
		self.e = e
	}
}

func (self *Parser) Parse(in io.Reader) (*ast.Program, error) {
	lex := lexer.New(in)
	pipe := make(chan *Token, 1)

	// TODO: get another pipe for *Token recycle, so it won't
	// go to garbage collection

	go func() {
		for lex.Scan() {
			t := lex.Token()

			if t.Line > MaxLine {
				self.fail(t.Line, t.Col, fmt.Errorf("too many lines"))
				break
			} else if t.Col > MaxCol {
				self.fail(t.Line, t.Col, fmt.Errorf("line too long"))
				break
			}

			pos := self.PosOffset + (uint32(t.Line) << 8) + uint32(t.Col)
			pipe <- &Token{t.Token, pos, t.Lit}
		}

		if self.e == nil {
			self.e = lex.Err()
		}

		close(pipe)
	}()

	self.s = NewTokenScanner(pipe)

	self.prog = ast.NewProgram()
	self.scanProgram()

	// return lex error first
	if self.e != nil {
		return nil, self.e
	}

	return self.prog, nil
}

func ParseFile(path string) (*ast.Program, error) {
	fin, e := os.Open(path)
	if e != nil {
		return nil, e
	}

	parser := New()
	parser.Reporter = reporter.NewPrefix(path)

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

func ParseStr(s string) (*ast.Program, error) {
	p := New()
	return p.Parse(strings.NewReader(s))
}
