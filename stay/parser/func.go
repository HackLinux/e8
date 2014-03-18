package parser

import (
	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/token"
)

func (self *Parser) parseFunc() {
	s := self.s

	if !s.CurIs(token.Ident) {
		self.failf("missing function name")
		return
	}

	t := s.Cur()
	decl := new(ast.FuncDecl)
	decl.Name = t.Lit
	decl.DeclLine = t.Line

	s.Next()

	// TODO: parameter list
	if !s.Scan(token.Lparen) {
		self.failf("expecting left parenthesis")
		return
	}

	if !s.Scan(token.Rparen) {
		self.failf("expecting right parenthesis")
		return
	}

	// TODO: return type

	decl.Body = self.parseBlockStmt()

	self.prog.AddFunc(decl)
}
