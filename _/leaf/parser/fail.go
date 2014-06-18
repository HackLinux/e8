package parser

import (
	"fmt"

	"e8vm.net/p/leaf/reporter"
	"e8vm.net/p/leaf/token"
)

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

func (self *Parser) expect(t token.Token) bool {
	if self.s.CurIs(t) {
		self.s.Next()
		return true
	}

	self.failf("expect %s, got %s", t, self.s.Cur().Token)
	return false
}

func (self *Parser) failExpect(s string) {
	self.failf("expect %s, got %s", s, self.s.Cur().Token)
}

func (self *Parser) unexpectedEOF() bool {
	if self.s.Closed() {
		self.failf("unexpected end of file")
		return true
	}
	return false
}
