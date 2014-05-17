package ir1

type Type interface {
	Size() uint32
	String() string
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

var basicTypeNames = map[BasicType]string{
	I8:  "i8",
	U8:  "u8",
	I16: "i16",
	U16: "u16",
	I32: "i32",
	U32: "u32",
}

func (t BasicType) Size() uint32 {
	return uint32(t) >> 1
}

func (t BasicType) String() string {
	ret, found := basicTypeNames[t]
	if !found {
		return "<?>"
	}

	return ret
}

type PointerType struct {
	of Type
}

const PointerSize = 4

func (p *PointerType) Size() uint32 {
	return PointerSize
}
