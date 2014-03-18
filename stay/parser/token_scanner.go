package parser

import (
	"github.com/h8liu/e8/stay/lexer"
	"github.com/h8liu/e8/stay/tokens"
)

type TokenScanner struct {
	lexer *lexer.Lexer
	cur   *lexer.Token
}

func NewTokenScanner(lex *lexer.Lexer) *TokenScanner {
	ret := new(TokenScanner)
	ret.lexer = lex
	ret.Next()

	return ret
}

func (self *TokenScanner) Next() *lexer.Token {
	// TODO: bind comments with tokens
	ret := self.cur

	for self.lexer.Scan() {
		self.cur = self.lexer.Token()
		if self.cur.Token == tokens.Comment {
			continue
		}

		break
	}

	return ret
}

func (self *TokenScanner) Pos() (int, int) {
	return self.cur.Line, self.cur.Col
}

func (self *TokenScanner) Peek() *lexer.Token {
	return self.cur
}

func (self *TokenScanner) Scan(t int) bool {
	if self.cur == nil {
		return false
	}
	return self.cur.Token == t
}

func (self *TokenScanner) Accept(t int) bool {
	if self.Scan(t) {
		self.Next()
		return true
	}

	return false
}

func (self *TokenScanner) Closed() bool {
	return self.cur == nil
}
