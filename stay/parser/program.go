package parser

import (
	"fmt"

	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/tokens"
)

func (self *Parser) failf(s string, args ...interface{}) {
	e := fmt.Errorf(s, args)
	if len(self.errors) < MaxError {
		self.errors = append(self.errors, e)
	}
}

func (self *Parser) scanProgram() {
	self.scanImports()
}

func (self *Parser) scanImports() {
	s := self.s

	for s.Accept(tokens.Import) {
		if s.Accept(tokens.Lparen) {
			for self.scanImportSpec() {
				if !s.Accept(tokens.Semicolon) {
					if s.Scan(tokens.Rparen) {
						break
					}
					self.failf("missing semicolon")
				}
			}

			if !s.Accept(tokens.Rparen) {
				self.failf("missing right parenthesis")
			}
		} else {
			if self.scanImportSpec() {
				if !s.Accept(tokens.Semicolon) {
					self.failf("missing semicolon")
				}
			} else {
				self.failf("missing import spec")
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
		self.failf("missing import path")
		if as != "" {
			return true
		} else {
			return false
		}
	}

	t := s.Next()
	self.prog.AddImport(&ast.ImportDecl{as, t.lit, t.pos})

	return true
}
