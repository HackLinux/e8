package ir1

type Expr interface {
	Type() Type
	String() string
}
