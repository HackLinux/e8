package lexer

import (
	"fmt"
	"io"

	"github.com/h8liu/e8/stay/reporters"
	"github.com/h8liu/e8/stay/tokens"
)

type Lexer struct {
	*scanner

	illegal    bool  // illegal encountered
	insertSemi bool  // if treat end line as whitespace
	eof        bool  // end of file returned
	FirstFail  error // lex error encountered
}

// Creates a new lexer
func New(in io.Reader) *Lexer {
	ret := new(Lexer)
	ret.scanner = newScanner(in)

	return ret
}

// Use a particular error reporter
func (self *Lexer) SetErrorReporter(reporter reporters.ErrReporter) {
	self.scanner.errReporter = reporter
}

// Reports a lex error
func (self *Lexer) failf(f string, args ...interface{}) {
	e := fmt.Errorf(f, args...)
	self.report(e)

	if self.FirstFail == nil {
		self.FirstFail = e
	}
}

func (self *Lexer) skipWhites() {
	if self.insertSemi {
		self.skipAnys(" \t\r")
	} else {
		self.skipAnys(" \t\r\n")
	}
}

func (self *Lexer) scanExpo() {
	self.scanAny("-+")
	self.scanDigits()
}

func (self *Lexer) _scanNumber(dotLed bool) (lit string, t int) {
	if !dotLed {
		if self.scan('0') {
			if self.scan('x') || self.scan('X') {
				if self.scanHexDigits() == 0 {
					return self.accept(), tokens.Illegal
				}
			} else {
				self.scanOctDigits()
			}

			if self.scanIdent() != 0 {
				return self.accept(), tokens.Illegal
			}

			return self.accept(), tokens.Int
		}

		self.scanDigits()

		if self.scanAny("eE") {
			self.scanExpo()
			if self.scanIdent() != 0 {
				return self.accept(), tokens.Illegal
			}
			return self.accept(), tokens.Float
		}

		if !self.scan('.') {
			if self.scanIdent() != 0 {
				return self.accept(), tokens.Illegal
			}
			return self.accept(), tokens.Int
		}

		self.scanDigits()
	} else {
		if self.scanDigits() == 0 {
			return self.accept(), tokens.Illegal
		}
	}

	if self.scanAny("eE") {
		self.scanAny("-+")
		if self.scanDigits() == 0 {
			return self.accept(), tokens.Illegal
		}
	}

	return self.accept(), tokens.Float
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
	if self.scanAny("abfnrtv\\") {
		return
	}
	if self.scan(q) {
		return
	}

	if self.scan('x') {
		if !(self.scanHexDigit() && self.scanHexDigit()) {
			self.failf("invalid hex escape")
		}
		return
	}

	if self.scanOctDigit() {
		if !(self.scanOctDigit() && self.scanOctDigit()) {
			self.failf("invalid oct escape")
		}
		return
	}

	self.failf("unknown escape char %q", self.peek())
	self.next()
	return
}

func (self *Lexer) scanChar() string {
	n := 0
	for !self.scan('\'') {
		if self.peek() == '\n' || self.closed {
			self.failf("char not terminated")
			break
		}

		if self.scan('\\') {
			self.scanEscape('\'')
		} else {
			self.next()
		}
		n++
	}

	if n != 1 {
		self.failf("illegal char")
	}

	return self.accept()
}

func (self *Lexer) scanComment() string {
	if self.scan('*') {
		for {
			if self.scan('*') {
				if self.scan('/') {
					return self.accept()
				}
				continue
			}

			if self.closed {
				self.failf("incomplete block comment")
				return self.accept()
			}
			self.next()
		}
	}

	if self.scan('/') {
		for {
			if self.peek() == '\n' || self.closed {
				return self.accept()
			}
			self.next()
		}
	}

	panic("bug")
}

func (self *Lexer) scanOperator(r rune) int {
	switch r {
	case '\n':
		self.insertSemi = false
		return tokens.Semicolon
	case ':':
		return tokens.Colon
	case '.':
		if self.scan('.') {
			if self.scan('.') {
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
		if self.scan('+') {
			return tokens.Inc
		} else if self.scan('=') {
			return tokens.AddAssign
		} else {
			return tokens.Add
		}
	case '-':
		if self.scan('-') {
			return tokens.Dec
		} else if self.scan('=') {
			return tokens.SubAssign
		}
		return tokens.Sub
	case '*':
		if self.scan('=') {
			return tokens.MulAssign
		}
		return tokens.Mul
	case '/':
		if self.scan('=') {
			return tokens.DivAssign
		}
		return tokens.Div
	case '%':
		if self.scan('=') {
			return tokens.ModAssign
		}
		return tokens.Mod
	case '^':
		if self.scan('=') {
			return tokens.XorAssign
		}
		return tokens.Xor
	case '<':
		if self.scan('=') {
			return tokens.Leq
		} else if self.scan('<') {
			if self.scan('=') {
				return tokens.ShiftLeftAssign
			}
			return tokens.ShiftLeft
		}
		return tokens.Less
	case '>':
		if self.scan('=') {
			return tokens.Geq
		} else if self.scan('>') {
			if self.scan('=') {
				return tokens.ShiftRightAssign
			}
			return tokens.ShiftRight
		}
		return tokens.Greater
	case '=':
		if self.scan('=') {
			return tokens.Eq
		}
		return tokens.Assign
	case '!':
		if self.scan('=') {
			return tokens.Neq
		}
		return tokens.Not
	case '&':
		if self.scan('^') {
			if self.scan('=') {
				return tokens.NandAssign
			} else {
				return tokens.Nand
			}
		}
		if self.scan('=') {
			return tokens.AndAssign
		} else if self.scan('&') {
			return tokens.Land
		}

		return tokens.And
	case '|':
		if self.scan('=') {
			return tokens.OrAssign
		} else if self.scan('|') {
			return tokens.Lor
		}
	}

	if !self.illegal {
		self.illegal = true
		self.failf("illegal character")
	}
	return tokens.Illegal
}

// Scanning error
func (self *Lexer) ScanErr() error {
	if self.err == io.EOF {
		return nil
	}
	return self.err
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

// Returns the next token.
// t is the token code, p is the position code,
// and lit is the string literal.
// Returns tokens.EOF in t for the last token.
func (self *Lexer) Token() (t int, p uint32, lit string) {
	t, p, lit = self.scanToken()

	if t != tokens.Illegal {
		self.insertSemi = insertSemiTokenMap[t]
	}

	return
}

// Returns if the scanner has anything to return
func (self *Lexer) Scan() bool { return !self.eof }

func (self *Lexer) scanToken() (t int, p uint32, lit string) {
	if self.eof {
		return tokens.EOF, self.pos(), ""
	}

	self.skipWhites()
	p = self.pos()

	if self.closed {
		if self.insertSemi {
			self.insertSemi = false
			return tokens.Semicolon, p, ";"
		}
		self.eof = true
		return tokens.EOF, p, ""
	}

	r := self.peek()

	if isLetter(r) {
		self.scanIdent()
		lit = self.accept()
		t = tokens.IdentToken(lit)

		return t, p, lit
	} else if isDigit(r) {
		lit, t = self.scanNumber(false)
		return t, p, lit
	} else if self.scan('\'') {
		self.insertSemi = true
		lit = self.scanChar()
		return tokens.Char, p, lit
	}

	self.next()

	if r == '.' && isDigit(self.peek()) {
		lit, t = self.scanNumber(true)
		return t, p, lit
	} else if r == '/' {
		r2 := self.peek()
		if r2 == '/' || r2 == '*' {
			s := self.scanComment()
			return tokens.Comment, p, s
		}
	}

	t = self.scanOperator(r)
	lit = self.accept()
	if t == tokens.Semicolon {
		lit = ";"
	}
	return t, p, lit
}
