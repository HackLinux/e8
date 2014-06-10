package parser

import (
	"github.com/h8liu/e8/leaf/ast"
	"github.com/h8liu/e8/leaf/token"
)

func (self *Parser) parseBlockStmt() *ast.BlockStmt {
	s := self.s

	if self.expect(token.Lbrace) {
		return nil
	}

	ret := ast.NewBlock()

	// statement list
	for !s.CurIs(token.Rbrace) {
		if self.unexpectedEOF() {
			return nil
		}

		stmt := self.parseStmt()
		if stmt != nil {
			ret.Add(stmt)
		}

		if s.CurIs(token.Rbrace) {
			break
		}

		self.expect(token.Semicolon)
	}

	if self.expect(token.Rbrace) {
		return nil
	}

	return ret
}

func (self *Parser) parseStmt() ast.Stmt {
	s := self.s

	t := s.Cur()
	switch t.Token {
	case token.Var, token.Type, token.Const:
		// TODO: parseDecls
	case token.Semicolon:
		// TODO: empty statement
	}

	exp := self.parseExpr()
	if exp != nil {
		return exp
	}

	return nil
}
