package ir

const (
	typPointer = iota
	typInt32 
	typUint32
	typInt16
	typUint16
	typInt8
	typUint8
	typBool
	typFloat64
)

const (
	opAdd = iota
	opSub
	opInv
	opMul
	opDiv
	opMod
	opOr
	opAnd
	opXor
	opNand

	opEq
	opNeq
	opGeq

	opLor
	opLand
	opNot

	opConvert
	opCall
)

type Op struct {
	op int 
	typ int // integer type
	x *Op // the first operand
	y *Op // the second operand
	isConst bool
	v int64 // the const value
	vf float64 // the floating point value
	lab string // label lookup for const
	jump string // label to jump if true
}
