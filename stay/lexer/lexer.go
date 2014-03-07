package lexer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/h8liu/e8/stay/pos"
	"github.com/h8liu/e8/stay/tokens"
)

type Lexer struct {
	reader *bufio.Reader

	r       rune
	rsize   int
	lineNo  int
	lineOff int
	illegal bool
	eof     bool  // end of file reached, or some fatal error occured
	e       error // if any error
	es      []error

	insertSemi bool // if treat end line as whitespace
}

func New(in io.Reader) *Lexer {
	ret := new(Lexer)
	ret.reader = bufio.NewReader(in)
	ret.lineNo = 1
	ret.es = make([]error, 0, 1000)

	return ret
}

func (self *Lexer) end(e error) rune {
	self.r = rune(0)
	self.rsize = 0
	self.eof = true
	self.e = e
	return rune(0)
}

func (self *Lexer) report(e error) {
	self.es = append(self.es, e)
}

func (self *Lexer) pos() uint32 {
	return uint32(self.lineNo)<<8 + uint32(self.lineOff)
}

func (self *Lexer) errorf(f string, args ...interface{}) error {
	return &Error{self.pos(), fmt.Errorf(f, args...)}
}

func (self *Lexer) panicf(f string, args ...interface{}) rune {
	return self.end(self.errorf(f, args...))
}

func (self *Lexer) failf(f string, args ...interface{}) {
	self.report(errorf(f, args...))
}

func (self *Lexer) next() rune {
	if self.eof {
		return rune(0)
	}

	self.lineOff += self.rsize
	if self.lineOff > pos.MaxCharPerLine {
		return self.panicf("line too long")
	}

	if self.r == '\n' {
		if self.lineNo >= pos.MaxLinePerFile {
			return self.panicf("too many lines in a file")
		}
		self.lineNo++
		self.lineOff = 0
	}

	var e error
	self.r, self.rsize, e = self.reader.ReadRune()
	if e == io.EOF {
		return self.end(nil)
	}
	if e != nil {
		return self.end(e)
	}

	return self.r
}

func (self *Lexer) accept(r rune) bool {
	if self.eof {
		return false
	}

	if self.r == r {
		self.next()
		return true
	}

	return false
}

func (self *Lexer) isWhite(r rune) bool {
	if r == '\n' && !self.insertSemi {
		return true
	}
	return isWhite(r)
}

func (self *Lexer) scanWhites() int {
	ret := 0

	for {
		r := self.peek()
		if self.isWhite(r) {
			ret++
			self.accept(r)
			continue
		}
		break
	}

	return ret
}

func (self *Lexer) scanLine() int {
	ret := 0

	for {
		r := self.peek()
		if r != '\n' {
			ret++
			self.accept(r)
			continue
		}
		break
	}

	if self.peek() == '\n' {
		self.accept('\n')
	}

	return ret
}

func (self *Lexer) scanIdent() string {
	r := self.peek()
	ret := string(r)
	self.accept(r)

	for {
		r := self.peek()
		if isLetter(r) || isDigit(r) {
			ret += string(r)
			self.accept(r)
			continue
		}

		break
	}

	return ret
}

func (self *Lexer) scanWildDigits() (lit string, max int) {
	for {
		r := self.peek()
		v := digitVal(r)
		if v < 0 {
			return
		}
		lit += string(r)
		self.accept(r)
		if v > max {
			max = v
		}
	}
}

func (self *Lexer) scanDigits() string {
	ret := ""
	for {
		r := self.peek()
		if !isDigit(r) {
			break
		}
		ret += string(r)
		self.accept(r)
	}
	return ret
}

func (self *Lexer) scanNumber(dotLed bool) (lit string, t int) {
	ret := ""

	if !dotLed {
		if self.accept('0') {
			lit += "0"
			r := self.peek()
			if r == 'x' || r == 'X' {
				lit += string(r)
				self.accept(r)

				s, m := self.scanWildDigits()
				lit += s
				if m >= 16 {
					self.failf("invalid hex number")
					return lit, tokens.Illegal
				}
			} else {
				s, m := self.scanWildDigits()
				lit += s
				if m >= 8 {
					self.failf("invalid octal number")
					return lit, tokens.Illegal
				}
			}

			return lit, tokens.Int
		}

		s, m := self.scanDigits()
		lit += s
		if !self.accept('.') {
			return lit, tokens.Int
		}
	}

	lit += "." + self.scanDigits()
	if self.scanAny("eE") {
		self.scanAny("-+")
		self.scanDigits()
	}

	return self.accept(), tokens.Float
}

func (self *Lexer) scanChar() string {
	panic("todo")
}

func (self *Lexer) scanComment(r rune) string {
	panic("todo")
}

func (self *Lexer) peek() rune {
	if self.eof {
		return rune(0)
	}

	return self.r
}

func (self *Lexer) scanSymbol(r rune) int {
	switch r {
	case '\n':
		self.insertSemi = false
		return tokens.Semicolon
	case ':':
		return tokens.Colon
	case '.':
		if self.accept('.') {
			if self.accept('.') {
				return tokens.Ellipsis
			}
			self.failf("two dots, expecting one more")
			return tokens.Illegal
		} else {
			return tokens.Period
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
		if self.accept('+') {
			return tokens.Inc
		} else if self.accept('=') {
			return tokens.AddAssign
		} else {
			return tokens.Add
		}
	case '-':
		if self.accept('-') {
			return tokens.Dec
		} else if self.accept('=') {
			return tokens.SubAssign
		}
		return tokens.Sub
	case '*':
		if self.accept('=') {
			return tokens.MulAssign
		}
		return tokens.Mul
	case '/':
		if self.accept('=') {
			return tokens.DivAssign
		}
		return tokens.Div
	case '%':
		if self.accept('=') {
			return tokens.ModAssign
		}
		return tokens.Mod
	case '^':
		if self.accept('=') {
			return tokens.XorAssign
		}
		return tokens.Xor
	case '<':
		if self.accept('=') {
			return tokens.Leq
		} else if self.accept('<') {
			if self.accept('=') {
				return tokens.ShiftLeftAssign
			}
			return tokens.ShiftLeft
		}
		return tokens.Less
	case '>':
		if self.accept('=') {
			return tokens.Geq
		} else if self.accept('>') {
			if self.accept('=') {
				return tokens.ShiftRightAssign
			}
			return tokens.ShiftRight
		}
		return tokens.Greater
	case '=':
		if self.accept('=') {
			return tokens.Eq
		}
		return tokens.Assign
	case '!':
		if self.accept('=') {
			return tokens.Neq
		}
		return tokens.Not
	case '&':
		if self.accept('^') {
			if self.accept('=') {
				return tokens.NandAssign
			} else {
				return tokens.Nand
			}
		}
		if self.accept('=') {
			return tokens.AndAssign
		} else if self.accept('&') {
			return tokens.Land
		}

		return tokens.And
	case '|':
		if self.accept('=') {
			return tokens.OrAssign
		} else if self.accept('|') {
			return tokens.Lor
		}
	}

	if !self.illegal {
		self.illegal = true
		self.errorf("illegal character")
	}
	return tokens.Illegal
}

func (self *Lexer) Scan() (t int, p uint32, lit string) {
	self.scanWhites()

	r := self.peek()
	p = self.pos()

	if isLetter(r) {
		lit = self.scanIdent()
		// handle keywords
		return tokens.Ident, p, lit
	} else if isDigit(r) {
		lit, t = self.scanNumber(false)
		return t, p, lit
	} else if r == '\'' {
		self.insertSemi = true
		lit = self.scanChar()
		return tokens.Char, p, lit
	} else if self.eof && self.e == nil {
		if self.insertSemi {
			self.insertSemi = false
			return tokens.Semicolon, self.pos(), "\n"
		}
		return tokens.EOF, p, ""
	}

	self.accept(r)

	if r == '.' && isDigit(self.peek()) {
		lit, t = self.scanNumber(true)
		return t, p, lit
	} else if r == '/' {
		r2 := self.peek()
		if r2 == '/' || r2 == '*' {
			self.accept(r2)
			lit = self.scanComment(r2)
			return tokens.Comment, p, lit
		}
	}

	t = self.scanSymbol(r)
	return t, p, ""
}
