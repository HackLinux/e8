package parser

import (
	"e8vm.net/e8/leaf/lexer"
	"e8vm.net/e8/leaf/token"
)

// Scanner wraps a lexer and provides handy scanning API for the parser. It
// assumes that the lexer will always emit EOF token when it closed, due to
// either real EOF or an error. When it reads the first EOF, it stops
// scanning; if the EOF is missing, it panics.
type Scanner struct {
	lexer *lexer.Lexer
	cur   *lexer.LexToken
}

// Creates the scanner that wraps the lexer.
func NewScanner(lex *lexer.Lexer) *Scanner {
	ret := new(Scanner)
	ret.lexer = lex
	ret.Next()

	return ret
}

// Update the token buffer with the next token. If the current token
// is EOF, nothing will happen. It panics if the underlying lexer
// stops scanning before an EOF
func (self *Scanner) Next() {
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

// Returns the position of the current token.
func (self *Scanner) Pos() (int, int) {
	return self.cur.Line, self.cur.Col
}

// Returns the current token. This is a static buffer so the
// content will change if the scanner moves via
// Next(), Scan() or SkipUtil())
func (self *Scanner) Cur() *lexer.LexToken {
	return self.cur
}

// Returns true if the current token is t.
func (self *Scanner) CurIs(t token.Token) bool {
	return self.cur.Token == t
}

// Returns true and performs a Next() if the current token is t.
func (self *Scanner) Scan(t token.Token) bool {
	if self.CurIs(t) {
		self.Next()
		return true
	}

	return false
}

// Returns true if the current token is an EOF
func (self *Scanner) Closed() bool {
	return self.CurIs(token.EOF)
}

// Continue scanning until the current token is t. Returns the
// number of tokens skipped.
func (self *Scanner) SkipUtil(t token.Token) int {
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
