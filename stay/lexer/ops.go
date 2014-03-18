package lexer

import (
	. "github.com/h8liu/e8/stay/token"
)

var simpleOps = map[rune]Token{
	':': Colon,
	',': Comma,
	';': Semicolon,
	'(': Lparen,
	')': Rparen,
	'[': Lbrack,
	']': Rbrack,
	'{': Lbrace,
	'}': Rbrace,
}

var eqOps = map[rune]*struct{ t, eqt Token }{
	'*': {Mul, MulAssign},
	'/': {Div, DivAssign},
	'%': {Mod, ModAssign},
	'^': {Xor, XorAssign},
	'=': {Assign, Eq},
	'!': {Not, Neq},
}

var xeqOps = map[rune]*struct {
	t, eqt, xt Token
	x          rune
}{
	'+': {Add, AddAssign, Inc, '+'},
	'-': {Sub, SubAssign, Dec, '-'},
	'|': {Or, OrAssign, Lor, '|'},
}

func (self *Lexer) scanOperator(r rune) Token {
	s := self.s

	if r == '\n' {
		self.insertSemi = false
		return Semicolon
	} else if r == '.' {
		if s.Scan('.') {
			if s.Scan('.') {
				return Ellipsis
			}
			self.failf("two dots, expecting one more")
			return Illegal
		}

		return Dot
	} else if ret, found := simpleOps[r]; found {
		return ret
	} else if o, found := eqOps[r]; found {
		if s.Scan('=') {
			return o.eqt
		}
		return o.t
	} else if o, found := xeqOps[r]; found {
		if s.Scan(o.x) {
			return o.xt
		} else if s.Scan('=') {
			return o.eqt
		}
		return o.t
	}

	switch r {
	case '<':
		if s.Scan('=') {
			return Leq
		} else if s.Scan('<') {
			if s.Scan('=') {
				return ShiftLeftAssign
			}
			return ShiftLeft
		}
		return Less
	case '>':
		if s.Scan('=') {
			return Geq
		} else if s.Scan('>') {
			if s.Scan('=') {
				return ShiftRightAssign
			}
			return ShiftRight
		}
		return Greater
	case '&':
		if s.Scan('^') {
			if s.Scan('=') {
				return NandAssign
			} else {
				return Nand
			}
		}
		if s.Scan('=') {
			return AndAssign
		} else if s.Scan('&') {
			return Land
		}

		return And
	}

	if !self.illegal {
		self.illegal = true
		self.failf("illegal character")
	}
	return Illegal
}
