package parser

import (
	"github.com/h8liu/e8/stay/tokens"
)

func (self *Parser) parseProgram() {
	self.parseImports()

	s := self.s
	for {
		if !self.parseDecls() {
			if s.Closed() {
				self.failf("missing EOF token")
			} else if s.Accept(tokens.EOF) {
				break
			} else if s.Accept(tokens.Func) {
				// TODO: parse function here
				panic("todo")
			} else {
				self.failf("expect declaration")
				break
			}
		}

		if !s.Accept(tokens.Semicolon) {
			self.expectSemicolon()
		}
	}

	// just silently drain the rest
	for !s.Closed() {
		s.Next()
	}
}

func (self *Parser) parseDecls() bool {
	s := self.s
	if s.Accept(tokens.Const) {
		self.parseConsts()
	} else if s.Accept(tokens.Type) {
		// TODO: parse type
		panic("todo")
	} else if s.Accept(tokens.Var) {
		// TODO: parse variable
		panic("todo")
	} else {
		return false
	}

	return true
}
