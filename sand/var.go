package sand

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
)

// Var defines a memory segment for storing data
type Var struct {
	typ int // type of the memory
	size int 
	init []byte
}

func NewVar(t int, size int) *Var {
	ret := new(Var)
	ret.typ = t
	ret.size = size
	return ret
}
