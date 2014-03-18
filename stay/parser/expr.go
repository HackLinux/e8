package parser

import (
	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/tokens"
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
	case tokens.Lparen:
		s.Next()
		return self.parseCallExpr(op)
	}

	return op
}

func (self *Parser) parseCallExpr(f ast.Expr) *ast.CallExpr {
	s := self.s

	ret := ast.NewCallExpr()
	ret.Func = f

	if !s.Scan(tokens.Rparen) {
		// non empty arg list
		for {
			exp := self.parseExpr()
			ret.AddArg(exp)

			if s.Scan(tokens.Rparen) {
				break
			}
			if !s.Scan(tokens.Comma) {
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
	case s.Scan(tokens.Ident):
		return &ast.Ident{lit}
	case s.Scan(tokens.Int):
		return &ast.IntLit{self.parseInt(lit)}
	case s.Scan(tokens.Float):
		return &ast.FloatLit{self.parseFloat(lit)}
	case s.Scan(tokens.String):
		return &ast.StringLit{self.unquote(lit)}
	case s.Scan(tokens.Char):
		return &ast.CharLit{self.unquoteChar(lit)}
	case s.Scan(tokens.Lparen):
		e := self.parseExpr()
		if !s.Scan(tokens.Rparen) {
			self.failf("expect right parenthesis")
		}
		return &ast.ParenExpr{e}
	}

	self.failf("expect operand")
	s.SkipUtil(tokens.Semicolon)
	return new(ast.BadExpr)
}
