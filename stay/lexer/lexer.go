package lexer

import (
	"fmt"
	"io"

	"github.com/h8liu/e8/stay/reporter"
	"github.com/h8liu/e8/stay/runes"
	"github.com/h8liu/e8/stay/scanner"
	"github.com/h8liu/e8/stay/tokens"
)

type Lexer struct {
	s        *scanner.Scanner
	reporter reporter.Interface

	line int
	col int

	illegal    bool  // illegal encountered
	insertSemi bool  // if treat end line as whitespace
	eof        bool  // end of file returned
	firstFail  error // lex error encountered
}

// Creates a new lexer
func New(in io.Reader) *Lexer {
	ret := new(Lexer)
	ret.s = scanner.New(in)

	return ret
}

// Use a particular error reporter
func (self *Lexer) ReportTo(r reporter.Interface) {
	self.reporter = r
}

func (self *Lexer) report(e error) {
	if self.reporter == nil {
		return
	}
	line, col := self.s.Pos()
	self.reporter.Report(line, col, e)
}

// Reports a lex error
func (self *Lexer) failf(f string, args ...interface{}) {
	e := fmt.Errorf(f, args...)
	self.report(e)

	if self.firstFail == nil {
		self.firstFail = e
	}
}

func (self *Lexer) skipWhites() {
	if self.insertSemi {
		self.s.SkipAnys(" \t\r")
	} else {
		self.s.SkipAnys(" \t\r\n")
	}
}

func (self *Lexer) _scanNumber(dotLed bool) (lit string, t int) {
	s := self.s

	if !dotLed {
		if s.Scan('0') {
			if s.Scan('x') || self.s.Scan('X') {
				if s.ScanHexDigits() == 0 {
					return s.Accept(), tokens.Illegal
				}
			} else {
				s.ScanOctDigits()
			}

			if s.ScanIdent() != 0 {
				return s.Accept(), tokens.Illegal
			}

			return s.Accept(), tokens.Int
		}

		s.ScanDigits()

		if s.ScanAny("eE") {
			s.ScanAny("-+")
			if s.ScanDigits() == 0 {
				return s.Accept(), tokens.Illegal
			}
			return s.Accept(), tokens.Float
		}

		if !s.Scan('.') {
			if s.ScanIdent() != 0 {
				return s.Accept(), tokens.Illegal
			}
			return s.Accept(), tokens.Int
		}

		s.ScanDigits()
	} else {
		if s.ScanDigits() == 0 {
			return s.Accept(), tokens.Illegal
		}
	}

	if s.ScanAny("eE") {
		s.ScanAny("-+")
		if s.ScanDigits() == 0 {
			return s.Accept(), tokens.Illegal
		}
	}

	return s.Accept(), tokens.Float
}

func (self *Lexer) scanNumber(dotLed bool) (lit string, t int) {
	lit, t = self._scanNumber(dotLed)
	if t == tokens.Illegal {
		self.failf("invalid number")
		t = tokens.Int
	}

	return
}

func (self *Lexer) scanEscape(q rune) {
	s := self.s

	if s.ScanAny("abfnrtv\\") {
		return
	}
	if s.Scan(q) {
		return
	}

	if s.Scan('x') {
		if !(s.ScanHexDigit() && s.ScanHexDigit()) {
			self.failf("invalid hex escape")
		}
		return
	}

	if s.ScanOctDigit() {
		if !(s.ScanOctDigit() && s.ScanOctDigit()) {
			self.failf("invalid octal escape")
		}
		return
	}

	self.failf("unknown escape char %q", s.Peek())
	s.Next()

	return
}

func (self *Lexer) scanChar() string {
	s := self.s
	n := 0
	for !s.Scan('\'') {
		if s.Peek() == '\n' || s.Closed() {
			self.failf("char not terminated")
			break
		}

		if s.Scan('\\') {
			self.scanEscape('\'')
		} else {
			s.Next()
		}
		n++
	}

	if n != 1 {
		self.failf("illegal char")
	}

	return s.Accept()
}

func (self *Lexer) scanComment() string {
	s := self.s

	if s.Scan('*') {
		for {
			if s.Scan('*') {
				if s.Scan('/') {
					return s.Accept()
				}
				continue
			}

			if s.Closed() {
				self.failf("incomplete block comment")
				return s.Accept()
			}
			s.Next()
		}
	}

	if s.Scan('/') {
		for {
			if s.Peek() == '\n' || s.Closed() {
				return s.Accept()
			}
			s.Next()
		}
	}

	panic("bug")
}

func (self *Lexer) scanOperator(r rune) int {
	s := self.s

	switch r {
	case '\n':
		self.insertSemi = false
		return tokens.Semicolon
	case ':':
		return tokens.Colon
	case '.':
		if s.Scan('.') {
			if s.Scan('.') {
				return tokens.Ellipsis
			}
			self.failf("two dots, expecting one more")
			return tokens.Illegal
		} else {
			return tokens.Dot
		}
	case ',':
		return tokens.Comma
	case ';':
		return tokens.Semicolon
	case '(':
		return tokens.Lparen
	case ')':
		return tokens.Rparen
	case '[':
		return tokens.Lbrack
	case ']':
		return tokens.Rbrack
	case '{':
		return tokens.Lbrace
	case '}':
		return tokens.Rbrack
	case '+':
		if s.Scan('+') {
			return tokens.Inc
		} else if s.Scan('=') {
			return tokens.AddAssign
		} else {
			return tokens.Add
		}
	case '-':
		if s.Scan('-') {
			return tokens.Dec
		} else if s.Scan('=') {
			return tokens.SubAssign
		}
		return tokens.Sub
	case '*':
		if s.Scan('=') {
			return tokens.MulAssign
		}
		return tokens.Mul
	case '/':
		if s.Scan('=') {
			return tokens.DivAssign
		}
		return tokens.Div
	case '%':
		if s.Scan('=') {
			return tokens.ModAssign
		}
		return tokens.Mod
	case '^':
		if s.Scan('=') {
			return tokens.XorAssign
		}
		return tokens.Xor
	case '<':
		if s.Scan('=') {
			return tokens.Leq
		} else if s.Scan('<') {
			if s.Scan('=') {
				return tokens.ShiftLeftAssign
			}
			return tokens.ShiftLeft
		}
		return tokens.Less
	case '>':
		if s.Scan('=') {
			return tokens.Geq
		} else if s.Scan('>') {
			if s.Scan('=') {
				return tokens.ShiftRightAssign
			}
			return tokens.ShiftRight
		}
		return tokens.Greater
	case '=':
		if s.Scan('=') {
			return tokens.Eq
		}
		return tokens.Assign
	case '!':
		if s.Scan('=') {
			return tokens.Neq
		}
		return tokens.Not
	case '&':
		if s.Scan('^') {
			if s.Scan('=') {
				return tokens.NandAssign
			} else {
				return tokens.Nand
			}
		}
		if s.Scan('=') {
			return tokens.AndAssign
		} else if s.Scan('&') {
			return tokens.Land
		}

		return tokens.And
	case '|':
		if s.Scan('=') {
			return tokens.OrAssign
		} else if s.Scan('|') {
			return tokens.Lor
		}
	}

	if !self.illegal {
		self.illegal = true
		self.failf("illegal character")
	}
	return tokens.Illegal
}

func (self *Lexer) Err() error {
	e := self.s.Err()
	if e != nil {
		return e
	}

	if self.firstFail != nil {
		return self.firstFail
	}

	return nil
}

func (self *Lexer) ScanErr() error {
	return self.s.Err()
}

func (self *Lexer) LexErr() error {
	return self.firstFail
}

var insertSemiTokens = []int{
	tokens.Ident,
	tokens.Int,
	tokens.Float,
	tokens.Break,
	tokens.Continue,
	tokens.Fallthrough,
	tokens.Return,
	tokens.Char,
	tokens.Rparen,
	tokens.Rbrack,
	tokens.Rbrace,
	tokens.Inc,
	tokens.Dec,
}

var insertSemiTokenMap = func() map[int]bool {
	ret := make(map[int]bool)
	for _, t := range insertSemiTokens {
		ret[t] = true
	}
	return ret
}()

func (self *Lexer) savePos() { self.line, self.col = self.s.Pos() }
func (self *Lexer) token(t int, lit string) *Token {
	return &Token{t, self.line, self.col, lit}
}

// Returns if the scanner has anything to return
func (self *Lexer) Scan() bool { return !self.eof }

// Returns the next token.
// t is the token code, p is the position code,
// and lit is the string literal.
// Returns tokens.EOF in t for the last token.
func (self *Lexer) Token() *Token {
	ret := self.scanToken()
	if ret.Token != tokens.Illegal {
		self.insertSemi = insertSemiTokenMap[ret.Token]
	}
	return ret
}

func (self *Lexer) scanToken() *Token {
	if self.eof {
		self.savePos()
		return self.token(tokens.EOF, "")
	}

	self.skipWhites()
	self.savePos()

	if self.s.Closed() {
		if self.insertSemi {
			self.insertSemi = false
			return self.token(tokens.Semicolon, ";")
		}
		self.eof = true

		e := self.s.Err()
		if e != nil {
			self.report(e)
		}
		return self.token(tokens.EOF, "")
	}

	s := self.s
	r := s.Peek()

	if runes.IsLetter(r) {
		s.ScanIdent()
		lit := s.Accept()
		t := tokens.IdentToken(lit)
		return self.token(t, lit)
	} else if runes.IsDigit(r) {
		lit, t := self.scanNumber(false)
		return self.token(t, lit)
	} else if r == '\'' {
		self.insertSemi = true
		s.Next()
		lit := self.scanChar()
		return self.token(tokens.Char, lit)
	}

	s.Next() // at this time, we will always make some progress

	if r == '.' && runes.IsDigit(s.Peek()) {
		lit, t := self.scanNumber(true)
		return self.token(t, lit)
	} else if r == '/' {
		r2 := s.Peek()
		if r2 == '/' || r2 == '*' {
			s := self.scanComment()
			return self.token(tokens.Comment, s)
		}
	}

	t := self.scanOperator(r)
	lit := s.Accept()
	if t == tokens.Semicolon {
		lit = ";"
	}

	return self.token(t, lit)
}
