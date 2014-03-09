package parser

import (
	"github.com/h8liu/e8/stay/tokens"
)

type TokenScanner struct {
	cur *Token
	c   <-chan *Token
}

func NewTokenScanner(c <-chan *Token) *TokenScanner {
	ret := new(TokenScanner)
	ret.c = c
	ret.Next()
	return ret
}

func (self *TokenScanner) Next() *Token {
	// TODO: bind comments with tokens
	ret := self.cur

	for {
		self.cur = <-self.c
		if self.cur.tok != tokens.Comment {
			break
		}
	}

	return ret
}

func (self *TokenScanner) Peek() *Token {
	return self.cur
}

func (self *TokenScanner) Scan(t int) bool {
	if self.cur == nil {
		return false
	}
	return self.cur.tok == t
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
