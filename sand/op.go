package sand

const (
	TypPointer = iota
	TypInt32 
	TypUint32
	TypInt16
	TypUint16
	TypInt8
	TypUint8
	TypBool
	TypFloat64
)

const (
	OpAdd = iota
	OpSub
	OpInv
	OpMul
	OpDiv
	OpMod
	OpOr
	OpAnd
	OpXor
	OpNand

	OpEq
	OpNeq
	OpGeq

	OpLor
	OpLand
	OpNot

	OpConvert
	OpCall
)

type Op struct {
	op int 
	typ int // integer type
	x *Op // the first operand
	y *Op // the second operand
	isConst bool
	v int64 // the const value
	vf float64 // the floating point value
	// lab string // label lookup for const
	jump string // label to jump if true
}
