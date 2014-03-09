package parser

import (
	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/tokens"
)

func (self *Parser) scanImports() {
	s := self.s

	for s.Accept(tokens.Import) {
		if s.Accept(tokens.Lparen) {
			for !s.Scan(tokens.Rparen) {
				if s.Closed() {
					self.failf("incomplete imports")
					return
				}

				self.scanImportSpec()
				if !s.Accept(tokens.Semicolon) {
					if s.Scan(tokens.Rparen) {
						break
					}
					self.expectSemicolon()
				}
			}

			if s.Accept(tokens.Rparen) {
				if !s.Accept(tokens.Semicolon) {
					self.expectSemicolon()
				}
			} else {
				self.failf("expect right parenthesis")
			}
		} else {
			if s.Closed() {
				self.failf("incomplete imports")
				return
			}

			if self.scanImportSpec() {
				if !s.Accept(tokens.Semicolon) {
					self.expectSemicolon()
				}
			} else {
				self.failf("expect import spec")
			}
		}
	}
}

func (self *Parser) scanImportSpec() bool {
	s := self.s

	var as string
	if s.Accept(tokens.Dot) {
		as = "."
	} else if s.Scan(tokens.Ident) {
		t := s.Next()
		as = t.lit
	}

	if !s.Scan(tokens.String) {
		self.failf("expect import path")
		if as != "" {
			return true
		} else {
			return false
		}
	}

	t := s.Next()
	path := unquote(t.lit)
	self.prog.AddImport(&ast.ImportDecl{as, path, t.pos})

	return true
}
