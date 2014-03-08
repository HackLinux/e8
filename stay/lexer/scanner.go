package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/h8liu/e8/stay/pos"
)

type scanner struct {
	reader      *bufio.Reader
	errReporter ErrReporter

	r rune

	tail *scanPos
	head *scanPos

	buf    *bytes.Buffer
	closed bool
	err    error
}

func newScanner(in io.Reader) *scanner {
	ret := new(scanner)
	ret.reader = bufio.NewReader(in)
	ret.head = newScanPos()
	ret.tail = newScanPos()
	ret.buf = new(bytes.Buffer)
	ret.errReporter = StderrReporter

	ret.next() // get ready for reading

	return ret
}

func (self *scanner) report(e error) {
	if self.errReporter == nil {
		return
	}
	self.errReporter.Report(
		uint16(self.tail.lineNo),
		uint8(self.tail.lineOffset),
		e,
	)
}

func (self *scanner) shutdown(e error) {
	if e != nil {
		self.report(e)
	}
	self.r = rune(-1)
	self.closed = true
	self.err = e
}

func (self *scanner) panicf(f string, args ...interface{}) {
	self.shutdown(fmt.Errorf(f, args...))
}

func (self *scanner) pos() uint32 {
	return self.tail.Pos()
}

func (self *scanner) next() rune {
	if self.closed {
		return rune(-1)
	}

	self.buf.WriteRune(self.r)

	var rsize int
	var e error
	self.r, rsize, e = self.reader.ReadRune()
	if e == io.EOF {
		self.shutdown(nil)
	} else if e != nil {
		self.shutdown(e)
	}

	if self.r == '\n' {
		self.head.NewLine()
		if self.head.lineNo >= pos.MaxLinePerFile {
			self.panicf("too many lines in a file")
		}
	} else {
		self.head.lineOffset += rsize
		if self.head.lineOffset >= pos.MaxRunePerLine {
			self.panicf("line too long")
		}
	}

	return self.r
}

func (self *scanner) peek() rune {
	return self.r
}

func (self *scanner) scan(r rune) bool {
	if self.r == r {
		self.next()
		return true
	}
	return false
}

func (self *scanner) scanDigits() int {
	ret := 0
	for isDigit(self.r) {
		ret++
		self.next()
	}

	return ret
}

func (self *scanner) scanHexDigit() bool {
	if isHexDigit(self.r) {
		self.next()
		return true
	}
	return false
}

func (self *scanner) scanOctDigit() bool {
	if isOctDigit(self.r) {
		self.next()
		return true
	}
	return false
}

func (self *scanner) scanHexDigits() int {
	ret := 0
	for isHexDigit(self.r) {
		ret++
		self.next()
	}
	return ret
}

func (self *scanner) scanOctDigits() int {
	ret := 0
	for isOctDigit(self.r) {
		ret++
		self.next()
	}
	return ret
}

func (self *scanner) scanIdent() int {
	ret := 0
	for isDigit(self.r) || isLetter(self.r) {
		ret++
		self.next()
	}
	return ret
}

func (self *scanner) scanAny(s string) bool {
	for _, r := range s {
		if r == self.r {
			self.scan(rune(r))
			return true
		}
	}

	return false
}

func (self *scanner) scanAnys(s string) int {
	ret := 0
	for self.scanAny(s) {
		ret++
	}
	return ret
}

func (self *scanner) sync() {
	self.buf.Reset()
	self.tail.SyncTo(self.head)
}

func (self *scanner) accept() string {
	ret := self.buf.String()
	self.sync()

	return ret
}

func (self *scanner) skipAny(s string) bool {
	ret := self.scanAny(s)
	self.sync()
	return ret
}

func (self *scanner) skipAnys(s string) int {
	ret := self.scanAnys(s)
	self.sync()
	return ret
}
