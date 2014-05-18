package types

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
