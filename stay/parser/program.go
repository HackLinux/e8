package parser

import (
	"github.com/h8liu/e8/stay/token"
)

func (self *Parser) parseProgram() {
	self.parseImports()

	if self.ImportsOnly {
		return
	}

	/*
		s := self.s
		for {
			if !self.parseDecls() {
				if s.Closed() {
					self.failf("missing EOF token")
				} else if s.Accept(token.EOF) {
					break
				} else if s.Accept(token.Func) {
					self.parseFuncDecl()
				} else {
					self.failf("expect declaration")
					break
				}
			}

			if !s.Accept(token.Semicolon) {
				self.expectSemicolon()
			}
		}

		// just silently drain the rest
		for !s.Closed() {
			s.Next()
		}
	*/
}

func (self *Parser) parseDecls() bool {
	s := self.s

	switch {
	case s.Scan(token.Const):
		self.parseConstDecls()
	case s.Scan(token.Type):
		self.parseTypeDecls()
	case s.Scan(token.Var):
		self.parseVarDecls()
	default:
		return false
	}

	return true
}
