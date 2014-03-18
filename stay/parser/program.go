package parser

import (
	"github.com/h8liu/e8/stay/tokens"
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
				} else if s.Accept(tokens.EOF) {
					break
				} else if s.Accept(tokens.Func) {
					self.parseFuncDecl()
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
	*/
}

func (self *Parser) parseDecls() bool {
	s := self.s

	switch {
	case s.Scan(tokens.Const):
		self.parseConstDecls()
	case s.Scan(tokens.Type):
		self.parseTypeDecls()
	case s.Scan(tokens.Var):
		self.parseVarDecls()
	default:
		return false
	}

	return true
}
