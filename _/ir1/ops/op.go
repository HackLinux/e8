package ops

type Op int

const (
	Add = iota
	Sub
	Mul
	Div
	Mod
	Sl
	Sr
	Band
	Bor
	Bxor

	// bellow will all result into Bool
	Eq
	G
	Geq
	L
	Leq

	And
	Or
	Not
)

var opStr = map[Op]string{
	Add:  "+",
	Sub:  "-",
	Mul:  "*",
	Div:  "/",
	Mod:  "%",
	Sl:   ">>",
	Sr:   "<<",
	Band: "&",
	Bor:  "|",
	Bxor: "^",

	Eq:  "==",
	G:   ">",
	Geq: ">=",
	L:   "<",
	Leq: "<=",

	And: "&&",
	Or:  "||",
	Not: "!",
}

func (op Op) String() string {
	ret, found := opStr[op]
	if found {
		return ret
	}
	return "??"
}
