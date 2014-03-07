package lexer

import (
	"io"

	"github.com/h8liu/e8/stay/tokens"
)

type Lexer struct {
	*scanner
	es []error

	illegal    bool // illegal encountered
	insertSemi bool // if treat end line as whitespace
}

func New(in io.Reader) *Lexer {
	ret := new(Lexer)
	ret.scanner = newScanner(in)
	ret.es = make([]error, 0, 1000)

	return ret
}

func (self *Lexer) failf(f string, args ...interface{}) {
	self.es = append(self.es, self.errorf(f, args...))
}

func (self *Lexer) LexErrors() []error { return self.es }

const whites = " \t\r"

func (self *Lexer) skipWhites() {
	if self.insertSemi {
		self.skipAnys(whites)
	} else {
		self.skipAnys(whites + "\n")
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

func (self *Lexer) scanChar() string {
	panic("todo")
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

func (self *Lexer) Closed() bool {
	return !self.insertSemi && self.closed
}

func (self *Lexer) Err() error {
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

func (self *Lexer) Scan() (t int, p uint32, lit string) {
	t, p, lit = self.scanToken()

	if t != tokens.Illegal {
		self.insertSemi = insertSemiTokenMap[t]
	}

	return
}

func (self *Lexer) scanToken() (t int, p uint32, lit string) {
	self.skipWhites()

	r := self.peek()
	p = self.pos()

	if isLetter(r) {
		self.scanIdent()
		lit = self.accept()
		t = tokens.IdentToken(lit)

		return t, p, lit
	} else if isDigit(r) {
		lit, t = self.scanNumber(false)
		return t, p, lit
	} else if r == '\'' {
		self.insertSemi = true
		lit = self.scanChar()
		return tokens.Char, p, lit
	} else if self.closed {
		if self.insertSemi {
			self.insertSemi = false
			return tokens.Semicolon, self.pos(), ";"
		}
		return tokens.EOF, p, ""
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
