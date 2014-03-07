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
	eof     bool  // end of file reached, or some fatal error occured
	e       error // if any error

	insertSemi bool // if treat end line as whitespace
}

func New(in io.Reader) *Lexer {
	ret := new(Lexer)
	ret.reader = bufio.NewReader(in)
	ret.lineNo = 1

	return ret
}

func (self *Lexer) end(e error) rune {
	self.r = rune(0)
	self.rsize = 0
	self.eof = true
	self.e = e
	return rune(0)
}

func (self *Lexer) pos() uint32 {
	return uint32(self.lineNo)<<8 + uint32(self.lineOff)
}

func (self *Lexer) errorf(f string, args ...interface{}) error {
	return &Error{self.pos(), fmt.Errorf(f, args...)}
}

func (self *Lexer) next() rune {
	if self.eof {
		return rune(0)
	}

	self.lineOff += self.rsize
	if self.lineOff > pos.MaxCharPerLine {
		return self.end(self.errorf("line too long"))
	}

	if self.r == '\n' {
		if self.lineNo >= pos.MaxLinePerFile {
			return self.end(self.errorf("too many lines in a file"))
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

func (self *Lexer) scanNumber() (lit string, t int) {
	panic("todo")
}

func (self *Lexer) scanChar() string {
	panic("todo")
}

func (self *Lexer) scanString() string {
	panic("todo")
}

func (self *Lexer) peek() rune {
	if self.eof {
		return rune(0)
	}

	return self.r
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
		lit, t = self.scanNumber()
		return t, p, lit
	} else if self.eof && self.e == nil {
		if self.insertSemi {
			self.insertSemi = false
			return tokens.Semicolon, self.pos(), "\n"
		}
		return tokens.EOF, p, ""
	}

	self.accept(r)

	switch r {
	case '\n':
		self.insertSemi = false
		return tokens.Semicolon, p, "\n"
	/*
		case '"':
			self.insertSemi = true // why?
			lit = self.scanString()
			return tokens.String, p, lit
	*/
	case '\'':
		self.insertSemi = true
		lit = self.scanChar()
		return tokens.Char, p, lit

	}

	panic("todo")
}
