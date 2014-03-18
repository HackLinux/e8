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

func (self *TokenScanner) Next() {
	if self.cur != nil && self.cur.Token == tokens.EOF {
		return
	}

	for self.lexer.Scan() {
		self.cur = self.lexer.Token()
		if self.cur.Token == tokens.Comment {
			continue
		}
		return
	}

	panic("EOF missing")
}

func (self *TokenScanner) Pos() (int, int) {
	return self.cur.Line, self.cur.Col
}

func (self *TokenScanner) Cur() *lexer.Token {
	return self.cur
}

func (self *TokenScanner) CurIs(t int) bool {
	return self.cur.Token == t
}

func (self *TokenScanner) Scan(t int) bool {
	if self.CurIs(t) {
		self.Next()
		return true
	}

	return false
}

func (self *TokenScanner) Closed() bool {
	return self.CurIs(tokens.EOF)
}

func (self *TokenScanner) SkipUtil(t int) int {
	ret := 0
	for self.cur.Token != t {
		self.Next()
		ret++
		if self.Closed() {
			return ret
		}
	}
	return ret
}
