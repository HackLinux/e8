package parser

import (
	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/tokens"
)

func (self *Parser) parseImports() {
	s := self.s

	for s.Scan(tokens.Import) {
		if s.Scan(tokens.Lparen) {
			for !s.CurIs(tokens.Rparen) {
				if s.CurIs(tokens.EOF) {
					self.failf("incomplete imports")
					return
				}

				self.parseImportSpec()
				if !s.Scan(tokens.Semicolon) {
					if s.CurIs(tokens.Rparen) {
						break
					}
					self.expectSemicolon()
				}
			}

			if s.Scan(tokens.Rparen) {
				if !s.Scan(tokens.Semicolon) {
					self.expectSemicolon()
				}
			} else {
				self.failf("expect right parenthesis")
			}
		} else {
			if s.CurIs(tokens.EOF) {
				self.failf("incomplete imports")
				return
			}

			if self.parseImportSpec() {
				if !s.Scan(tokens.Semicolon) {
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
	if s.Scan(tokens.Dot) {
		as = "."
	} else if s.CurIs(tokens.Ident) {
		as = s.Cur().Lit
		s.Next()
	}

	if !s.CurIs(tokens.String) {
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
