package parser

import (
	"github.com/h8liu/e8/leaf/ast"
	"github.com/h8liu/e8/leaf/token"
)

func (self *Parser) parseImports() {
	s := self.s

	for s.Scan(token.Import) {
		self.parseImport()
	}
}

func (self *Parser) parseImport() {
	s := self.s

	if s.Scan(token.Lparen) {
		for !s.CurIs(token.Rparen) {
			if self.unexpectedEOF() {
				return
			}

			self.parseImportSpec()

			// skipping semicolon
			if s.CurIs(token.Rparen) {
				break
			}

			self.expect(token.Semicolon)
		}

		if self.expect(token.Rparen) {
			self.expect(token.Semicolon)
		}
	} else {
		self.parseImportSpec()
		self.expect(token.Semicolon)
	}
}

func (self *Parser) parseImportSpec() {
	s := self.s

	var as string
	if s.Scan(token.Dot) {
		as = "."
	} else if s.CurIs(token.Ident) {
		as = s.Cur().Lit
		s.Next()
	}

	if !s.CurIs(token.String) {
		self.failExpect("import path")
		return
	}

	t := s.Cur()
	path := self.unquote(t.Lit)
	self.prog.AddImport(&ast.ImportDecl{as, path, t.Line})
	s.Next()
}
