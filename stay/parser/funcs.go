package parser

import (
	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/tokens"
)

func (self *Parser) parseFuncDecl() {
	s := self.s

	if !s.CurIs(tokens.Ident) {
		self.failf("missing function name")
		return
	}

	t := s.Cur()
	decl := new(ast.FuncDecl)
	decl.Name = t.Lit
	decl.DeclLine = t.Line

	s.Next()

	// TODO: parameter list
	if !s.Scan(tokens.Lparen) {
		self.failf("expecting left parenthesis")
		return
	}

	if !s.Scan(tokens.Rparen) {
		self.failf("expecting right parenthesis")
		return
	}

	// TODO: return type

	decl.Body = self.parseBlockStmt()

	self.prog.AddFunc(decl)
}
