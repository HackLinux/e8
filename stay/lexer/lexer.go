package lexer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/h8liu/e8/stay/pos"
)

type Lexer struct {
	reader *bufio.Reader

	r       rune
	rsize   int
	lineNo  int
	lineOff int
	eof     bool  // end of file reached, or some fatal error occured
	e       error // if any error

	whiteEndl bool // if treat end line as whilespace
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
	if r == ' ' || r == '\r' || r == '\t' {
		return true
	}
	if r == '\n' && self.whiteEndl {
		return true
	}
	return false
}

func (self *Lexer) acceptWhites() int {
	ret := 0

	for {
		r := self.peek()
		if self.isWhite(r) {
			ret++
			continue
		}
		break
	}

	return ret
}

func (self *Lexer) peek() rune {
	if self.eof {
		return rune(0)
	}

	return self.r
}

func (self *Lexer) Scan() (t int, p uint32, lit string) {
	self.acceptWhites()
	panic("todo")
}
