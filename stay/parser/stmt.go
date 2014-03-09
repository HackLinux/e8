package parser

import (
	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/tokens"
)

func (self *Parser) parseBlockStmt() *ast.BlockStmt {
	s := self.s

	if !s.Accept(tokens.Lbrace) {
		self.failf("expecting left brace")
	}

	ret := ast.NewBlock()

	// statement list
	for !s.Scan(tokens.Rbrace) {
		if s.Scan(tokens.EOF) || s.Closed() {
			self.failf("expecting rbrace, but reached eof")
			return ret
		}

		stmt := self.parseStmt()
		if stmt != nil {
			ret.Add(stmt)
		}

		if s.Scan(tokens.Rbrace) {
			break
		}

		if !s.Accept(tokens.Semicolon) {
			self.failf("missing semicolon")
		}
	}

	if !s.Accept(tokens.Rbrace) {
		self.failf("missing rbrace")
	}

	return ret
}

func (self *Parser) parseStmt() ast.Stmt {
	s := self.s

	t := s.Peek()
	switch t.tok {
	case tokens.Var, tokens.Type, tokens.Const:
		// TODO: parseDecls
	case tokens.Semicolon:
		// TODO: empty statement
	}

	exp := self.parseExpr()
	if exp != nil {
		return exp
	}

	return nil
}
