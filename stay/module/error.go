package module

import (
	"fmt"
)

type Error struct {
	pos uint32
	error
}

func (self *Error) Str(filename string) string {
	lineno := uint16(self.pos >> 8)
	offset := uint8(self.pos)
	return fmt.Sprintf("%s:%d:%d: %v", filename,
		lineno, offset, self.error)
}
