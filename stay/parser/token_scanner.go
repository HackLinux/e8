package parser

import (
	"github.com/h8liu/e8/stay/lexer"
	"github.com/h8liu/e8/stay/token"
)

type TokenScanner struct {
	lexer *lexer.Lexer
	cur   *lexer.LexToken
}

func NewTokenScanner(lex *lexer.Lexer) *TokenScanner {
	ret := new(TokenScanner)
	ret.lexer = lex
	ret.Next()

	return ret
}

func (self *TokenScanner) Next() {
	if self.cur != nil && self.cur.Token == token.EOF {
		return
	}

	for self.lexer.Scan() {
		self.cur = self.lexer.Token()
		if self.cur.Token == token.Comment {
			continue
		}
		return
	}

	panic("EOF missing")
}

func (self *TokenScanner) Pos() (int, int) {
	return self.cur.Line, self.cur.Col
}

func (self *TokenScanner) Cur() *lexer.LexToken {
	return self.cur
}

func (self *TokenScanner) CurIs(t token.Token) bool {
	return self.cur.Token == t
}

func (self *TokenScanner) Scan(t token.Token) bool {
	if self.CurIs(t) {
		self.Next()
		return true
	}

	return false
}

func (self *TokenScanner) Closed() bool {
	return self.CurIs(token.EOF)
}

func (self *TokenScanner) SkipUtil(t token.Token) int {
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
