package types

type Type interface {
	Size() uint32
	String() string
	Equal(t Type) bool
}
