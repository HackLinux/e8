package ir1

type Type interface {
	Size() uint32
}

type BasicType int

const (
	I8  BasicType = (1 << 1)
	U8  BasicType = (1 << 1) + 1
	I16 BasicType = (2 << 1)
	U16 BasicType = (2 << 1) + 1
	I32 BasicType = (4 << 1)
	U32 BasicType = (4 << 1) + 1
)

func (t BasicType) Size() uint32 {
	return uint32(t) >> 1
}

type PointerType struct {
	of Type
}

const PointerSize = 4

func (p *PointerType) Size() uint32 {
	return PointerSize
}
