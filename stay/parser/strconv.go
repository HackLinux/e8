package parser

import (
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
