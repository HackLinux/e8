// Package parser parses an input file into an abstract syntax tree.
package parser

import (
	"io"

	"github.com/h8liu/e8/leaf/ast"
	"github.com/h8liu/e8/leaf/lexer"
	"github.com/h8liu/e8/leaf/reporter"
	"github.com/h8liu/e8/leaf/token"
)

const (
	MaxError = 128
)

type Parser struct {
	// the error reporter on the entire parsing process
	Reporter reporter.Interface

	// token position will all be offset by this value on parsing
	PosOffset uint32

	ImportsOnly bool

	s        *Scanner
	prog     *ast.Program
	errors   []*reporter.Error
	lastLine int
}

func New() *Parser {
	ret := new(Parser)
	ret.Reporter = reporter.Simple
	ret.errors = make([]*reporter.Error, 0, MaxError)

	return ret
}

func (self *Parser) Parse(in io.Reader) (*ast.Program, error) {
	lex := lexer.New(in)
	lex.ReportTo(self.Reporter)

	// Prepare the scanner
	self.s = NewScanner(lex)

	// Parse the program now
	self.prog = ast.NewProgram()
	self.parseProgram()

	// Check lex error first
	e := lex.Err()
	if e != nil {
		return nil, e
	}

	// Return parse error, if any
	if len(self.errors) > 0 {
		for _, re := range self.errors {
			self.Reporter.Report(re.Line, re.Col, re.E)
		}
		return nil, self.errors[0]
	}

	return self.prog, nil
}

func (self *Parser) parseProgram() {
	self.parseImports()

	if self.ImportsOnly {
		return
	}

	s := self.s
	for !s.Closed() {
		switch {
		case s.Scan(token.Func):
			self.parseFunc()
		case s.Scan(token.Const):
			self.parseConst()
		case s.Scan(token.Var):
			self.parseVar()
		case s.Scan(token.Type):
			self.parseType()
		default:
			self.failExpect("declare")
			return
		}
	}
}
