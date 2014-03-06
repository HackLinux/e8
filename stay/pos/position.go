package pos

import (
	"fmt"
)

type Position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}

func (self *Position) Valid() bool {
	return self.Line > 0
}

func (self *Position) String() string {
	s := self.Filename

	if !self.Valid() {
		if s != "" {
			return s
		}
		return "-"
	}

	// self.Valid()
	loc := fmt.Sprintf("%d:%d", self.Line, self.Column)
	if s == "" {
		return loc
	}
	return s + ":" + loc
}
