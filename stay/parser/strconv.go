package parser

import (
	"math"
	"strconv"
)

func (self *Parser) unquote(s string) string {
	ret, e := strconv.Unquote(s)
	if e != nil {
		self.failf("invalid string literal")
		return ""
	}
	return ret
}

func (self *Parser) unquoteChar(s string) uint8 {
	n := len(s)
	if n < 3 {
		self.failf("invalid char literal")
		return 0
	}
	if s[0] != '\'' || s[n-1] != '\'' {
		self.failf("invalid quoting char literal")
		return 0
	}

	s = s[1 : n-1]
	ret, multi, tail, err := strconv.UnquoteChar(s, '\'')
	if multi {
		self.failf("multibyte char not allowed")
	} else if tail != "" {
		self.failf("char lit has a tail")
	} else if err != nil {
		self.failf("invalid char literal: %s, %v", s, err)
	} else if ret > math.MaxUint8 || ret < 0 {
		self.failf("invalid char value")
	}

	return uint8(ret)
}

func (self *Parser) parseInt(s string) int64 {
	ret, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		self.failf("invalid int literal")
	}
	return int64(ret)
}

func (self *Parser) parseFloat(s string) float64 {
	ret, err := strconv.ParseFloat(s, 64)
	if err != nil {
		self.failf("invalid float literal")
	}
	return ret
}
