package parser

import (
	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/token"
)

func (self *Parser) parseImports() {
	s := self.s

	for s.Scan(token.Import) {
		if s.Scan(token.Lparen) {
			for !s.CurIs(token.Rparen) {
				if s.CurIs(token.EOF) {
					self.failf("incomplete imports")
					return
				}

				self.parseImportSpec()
				if !s.Scan(token.Semicolon) {
					if s.CurIs(token.Rparen) {
						break
					}
					self.expectSemicolon()
				}
			}

			if s.Scan(token.Rparen) {
				if !s.Scan(token.Semicolon) {
					self.expectSemicolon()
				}
			} else {
				self.failf("expect right parenthesis")
			}
		} else {
			if s.CurIs(token.EOF) {
				self.failf("incomplete imports")
				return
			}

			if self.parseImportSpec() {
				if !s.Scan(token.Semicolon) {
					self.expectSemicolon()
				}
			} else {
				self.failf("expect import spec")
			}
		}
	}
}

func (self *Parser) parseImportSpec() bool {
	s := self.s

	var as string
	if s.Scan(token.Dot) {
		as = "."
	} else if s.CurIs(token.Ident) {
		as = s.Cur().Lit
		s.Next()
	}

	if !s.CurIs(token.String) {
		self.failf("expect import path")
		if as != "" {
			return true
		} else {
			return false
		}
	}

	t := s.Cur()
	path := self.unquote(t.Lit)
	self.prog.AddImport(&ast.ImportDecl{as, path, t.Line})
	s.Next()

	return true
}
