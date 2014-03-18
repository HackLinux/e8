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

	for !s.Accept(tokens.Rparen) {
		exp := self.parseExpr()
		ret.AddArg(exp)

		if s.Accept(tokens.Rparen) {
			break
		}

		if !s.Accept(tokens.Comma) {
			self.failf("expect comma")
		}

		if s.Closed() || s.Scan(tokens.EOF) {
			self.failf("incomplete call expression")
			return ret
		}
	}

	return ret
}

func (self *Parser) parseOperand() ast.Expr {
	s := self.s

	t := s.Cur()
	lit := t.Lit
	switch t.Token {
	case tokens.Ident:
		s.Next()
		return &ast.Ident{lit}
	case tokens.Int:
		s.Next()
		return &ast.IntLit{self.parseInt(lit)}
	case tokens.Float:
		s.Next()
		return &ast.FloatLit{self.parseFloat(lit)}
	case tokens.String:
		s.Next()
		return &ast.StringLit{self.unquote(lit)}
	case tokens.Char:
		s.Next()
		return &ast.CharLit{self.unquoteChar(lit)}
	case tokens.Lparen:
		s.Accept(tokens.Lparen)
		e := self.parseExpr()
		if !s.Accept(tokens.Rparen) {
			self.failf("expect right parenthesis")
		}
		return &ast.ParenExpr{e}
	}

	self.failf("expect operand")

	for {
		if s.Scan(tokens.Semicolon) {
			break
		}
		if s.Scan(tokens.EOF) || s.Closed() {
			break
		}
	}

	return new(ast.BadExpr)
}
