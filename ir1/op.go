package ir1

type Op int

const (
	OpAdd = iota
	OpSub
	OpMul
	OpDiv
	OpMod
	OpSl
	OpSr
	OpBand
	OpBor
	OpBxor

	// bellow will all result into Bool
	OpEq
	OpG
	OpGeq
	OpL
	OpLeq

	OpAnd
	OpOr
	OpNot
)

var opStr = map[Op]string{
	OpAdd:  "+",
	OpSub:  "-",
	OpMul:  "*",
	OpDiv:  "/",
	OpMod:  "%",
	OpSl:   ">>",
	OpSr:   "<<",
	OpBand: "&",
	OpBor:  "|",
	OpBxor: "^",

	OpEq:  "==",
	OpG:   ">",
	OpGeq: ">=",
	OpL:   "<",
	OpLeq: "<=",

	OpAnd: "&&",
	OpOr:  "||",
	OpNot: "!",
}

func (op Op) String() string {
	ret, found := opStr[op]
	if found {
		return ret
	}
	return "??"
}
