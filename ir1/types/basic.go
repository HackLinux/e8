package types

type Basic int

const (
	Void Basic = 0
	Bool Basic = 1
	I8   Basic = (1 << 1)
	U8   Basic = (1 << 1) + 1
	I16  Basic = (2 << 1)
	U16  Basic = (2 << 1) + 1
	I32  Basic = (4 << 1)
	U32  Basic = (4 << 1) + 1
)

var basicTypeNames = map[Basic]string{
	Void: "void",
	Bool: "bool",
	I8:   "i8",
	U8:   "u8",
	I16:  "i16",
	U16:  "u16",
	I32:  "i32",
	U32:  "u32",
}

func (t Basic) Size() uint32 {
	if t == Bool {
		return 1
	}

	return uint32(t) >> 1
}

func (t Basic) String() string {
	ret, found := basicTypeNames[t]
	if !found {
		return "<?>"
	}

	return ret
}

func (t Basic) Equal(other Type) bool {
	switch o := other.(type) {
	default:
		return false
	case Basic:
		return o == t
	}
}
