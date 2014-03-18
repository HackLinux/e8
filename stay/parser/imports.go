package parser

import (
	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/tokens"
)

func (self *Parser) parseImports() {
	s := self.s

	for s.Accept(tokens.Import) {
		if s.Accept(tokens.Lparen) {
			for !s.Scan(tokens.Rparen) {
				if s.Scan(tokens.EOF) {
					self.failf("incomplete imports")
					return
				}

				self.parseImportSpec()
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
			if s.Scan(tokens.EOF) {
				self.failf("incomplete imports")
				return
			}

			if self.parseImportSpec() {
				if !s.Accept(tokens.Semicolon) {
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
	if s.Accept(tokens.Dot) {
		as = "."
	} else if s.Scan(tokens.Ident) {
		as = s.Cur().Lit
		s.Next()
	}

	if !s.Scan(tokens.String) {
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
