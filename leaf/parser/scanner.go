package parser

import (
	"io"

	"github.com/h8liu/e8/leaf/lexer"
	t "github.com/h8liu/e8/leaf/token"
)

type scanner struct {
	lexer *lexer.Lexer
	cur   *lexer.Token
	last  *lexer.Token
}

func newScanner(in io.Reader, filename string) *scanner {
	ret := new(scanner)
	ret.lexer = lexer.New(in, filename)
	ret.shift()

	return ret
}

// reads in the next token
// return false if the current token is already end-of-file
func (self *scanner) shift() bool {
	if self.cur.Token == t.EOF {
		return false
	}

	for self.lexer.Scan() {
		self.last = self.cur
		self.cur = self.lexer.Token()
		if self.cur.Token != t.Comment {
			return true
		}
	}

	panic("should never reach here")
}

func (self *scanner) ahead(tok t.Token) bool {
	return self.cur.Token == tok
}

func (self *scanner) expect(tok t.Token) bool {
	if tok == t.EOF {
		panic("cannot expect EOF")
	}

	if self.ahead(tok) {
		return self.shift()
	} else {
		// append error node
	}

	return false
}

func (self *scanner) eof() bool {
	return self.ahead(t.EOF)
}

func (self *scanner) skipUntil(tok t.Token) []*lexer.Token {
	var skipped []*lexer.Token

	for !self.ahead(tok) {
		skipped = append(skipped, self.cur)
		if !self.shift() {
			return skipped
		}
	}

	// shift the last one
	skipped = append(skipped, self.cur)
	self.shift()

	return skipped
}
