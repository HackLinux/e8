package sand

import (
	"fmt"

	"github.com/h8liu/e8/printer"
)

const (
	TypPtr = iota
	TypBool
	TypUint8
	TypInt8
	TypUint16
	TypInt16
	TypUint32
	TypInt32
	TypFloat64
	TypFuncPtr
)

var typStrs = map[int]string{
	TypPtr:     "ptr",
	TypBool:    "bool",
	TypUint8:   "u8",
	TypInt8:    "i8",
	TypUint16:  "u16",
	TypInt16:   "i16",
	TypUint32:  "u32",
	TypInt32:   "i32",
	TypFloat64: "f64",
	TypFuncPtr: "fptr",
}

func TypStr(t int) string {
	if s, found := typStrs[t]; found {
		return s
	}
	return fmt.Sprintf("typ%d", t)
}

// Var defines a memory segment for storing data
type Var struct {
	typ  int // type of the memory
	size int
	init []byte
	name string

	onStack bool
	offset  int
}

func (self *Var) PrintTo(p printer.Interface) {
	p.Printf("%s %s(%d)", self.name, TypStr(self.typ), self.size)
}
