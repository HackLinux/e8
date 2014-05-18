package types

type Type interface {
	Size() uint32
	String() string
	Equal(t Type) bool
}

func IsSame(t1, t2 Type) bool {
	return t1.Equal(t2)
}
