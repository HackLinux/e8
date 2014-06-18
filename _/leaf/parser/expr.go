package parser

import (
	"e8vm.net/e8/leaf/ast"
	"e8vm.net/e8/leaf/token"
)

func (self *Parser) parseExpr() ast.Expr {
	return self.parseUnaryExpr()
}

func (self *Parser) parseUnaryExpr() ast.Expr {
	return self.parsePrimaryExpr()
}

func (self *Parser) parsePrimaryExpr() ast.Expr {
	s := self.s
	op := self.parseOperand()
	t := s.Cur()
	switch t.Token {
	case token.Lparen:
		s.Next()
		return self.parseCallExpr(op)
	}

	return op
}

func (self *Parser) parseCallExpr(f ast.Expr) *ast.CallExpr {
	s := self.s

	ret := ast.NewCallExpr()
	ret.Func = f

	if !s.Scan(token.Rparen) {
		// non empty arg list
		for {
			exp := self.parseExpr()
			ret.AddArg(exp)

			if s.Scan(token.Rparen) {
				break
			}
			if !s.Scan(token.Comma) {
				self.failf("expect comma")
			}
			if s.Closed() {
				self.failf("incomplete call expression")
				return ret
			}
		}
	}

	return ret
}

func (self *Parser) parseOperand() ast.Expr {
	s := self.s

	lit := s.Cur().Lit

	switch {
	case s.Scan(token.Ident):
		return &ast.Ident{lit}
	case s.Scan(token.Int):
		return &ast.IntLit{self.parseInt(lit)}
	case s.Scan(token.Float):
		return &ast.FloatLit{self.parseFloat(lit)}
	case s.Scan(token.String):
		return &ast.StringLit{self.unquote(lit)}
	case s.Scan(token.Char):
		return &ast.CharLit{self.unquoteChar(lit)}
	case s.Scan(token.Lparen):
		e := self.parseExpr()
		if !s.Scan(token.Rparen) {
			self.failf("expect right parenthesis")
		}
		return &ast.ParenExpr{e}
	}

	self.failf("expect operand")
	s.SkipUtil(token.Semicolon)
	return new(ast.BadExpr)
}
