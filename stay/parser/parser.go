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

	ImportsOnly bool

	s        *TokenScanner
	e        error
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

func (self *Parser) fail(line, col int, e error) {
	self.Reporter.Report(line, col, e)
	if self.e == nil {
		self.e = e
	}
}

func (self *Parser) failf(f string, args ...interface{}) {
	if len(self.errors) < MaxError {
		e := fmt.Errorf(f, args...)
		line, col := self.s.Pos()
		if self.lastLine == line {
			return
		}
		self.lastLine = line
		self.errors = append(self.errors, &reporter.Error{line, col, e})
	}
}

func (self *Parser) expectSemicolon() {
	self.failf("expect semicolon")
}

func (self *Parser) Parse(in io.Reader) (*ast.Program, error) {
	lex := lexer.New(in)
	lex.ReportTo(self.Reporter)
	pipe := make(chan *Token, 1)

	// TODO: get another pipe for *Token recycle, so it won't
	// go to garbage collection

	// TODO: need to change this pipe structure since we might not need
	// to read the entire file.
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
	self.parseProgram()

	// return lex error first
	if self.e != nil {
		return nil, self.e
	}

	// return parse error
	if len(self.errors) > 0 {
		for _, re := range self.errors {
			self.Reporter.Report(re.Line, re.Col, re.E)
		}
		return nil, self.errors[0]
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
