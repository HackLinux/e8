package parser

import (
	"strconv"
)

func unquote(s string) string {
	ret, e := strconv.Unquote(s)
	if e != nil {
		return ""
	}
	return ret
}
