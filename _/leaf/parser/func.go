package parser

import (
	"e8vm.net/e8/leaf/ast"
	"e8vm.net/e8/leaf/token"
)

func (self *Parser) parseFunc() {
	s := self.s

	if !s.CurIs(token.Ident) {
		self.expect(token.Ident)
		return
	}

	t := s.Cur()
	decl := &ast.FuncDecl{
		Name:     t.Lit,
		DeclLine: t.Line,
	}
	s.Next()

	if !self.expect(token.Lparen) {
		return
	}

	// TODO: parameter list

	if !self.expect(token.Rparen) {
		return
	}

	// TODO: return type

	decl.Body = self.parseBlockStmt()

	self.prog.AddFunc(decl)
}
