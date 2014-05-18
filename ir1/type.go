package ir1

type Type interface {
	Size() uint32
	String() string
	Equal(t Type) bool
}

type BasicType int

const (
	Bool BasicType = 1
	I8   BasicType = (1 << 1)
	U8   BasicType = (1 << 1) + 1
	I16  BasicType = (2 << 1)
	U16  BasicType = (2 << 1) + 1
	I32  BasicType = (4 << 1)
	U32  BasicType = (4 << 1) + 1
)

var basicTypeNames = map[BasicType]string{
	Bool: "bool",
	I8:   "i8",
	U8:   "u8",
	I16:  "i16",
	U16:  "u16",
	I32:  "i32",
	U32:  "u32",
}

func (t BasicType) Size() uint32 {
	if t == Bool {
		return 1
	}

	return uint32(t) >> 1
}

func (t BasicType) String() string {
	ret, found := basicTypeNames[t]
	if !found {
		return "<?>"
	}

	return ret
}

func (t BasicType) Equal(other Type) bool {
	switch o := other.(type) {
	default:
		return false
	case BasicType:
		return o == t
	}
}

type PointerType struct {
	of Type
}

const PointerSize = 4

func (t *PointerType) Size() uint32 {
	return PointerSize
}

func (t *PointerType) Equal(other Type) bool {
	switch o := other.(type) {
	default:
		return false
	case *PointerType:
		return t.of.Equal(o.of)
	}
}

func SameType(t1, t2 Type) bool {
	return t1.Equal(t2)
}

func (t *PointerType) String() string {
	return "*" + t.of.String()
}
