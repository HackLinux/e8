package tokens

import (
	"fmt"
)

var tokenStr = map[int]string{
	Illegal: "Illegal",
	EOF:     "EOF",
	Comment: "Comment",

	Ident: "Ident",
	Int:   "Int",
	Float: "Float",
	Char:  "Char",

	Add: "+",
	Sub: "-",
	Mul: "*",
	Div: "/",
	Mod: "%",

	And:        "&",
	Or:         "|",
	Xor:        "^",
	ShiftLeft:  "<<",
	ShiftRight: ">>",
	Nand:       "&^",

	AddAssign: "+=",
	SubAssign: "-=",
	MulAssign: "*=",
	DivAssign: "/=",
	ModAssign: "%=",

	AndAssign:        "&=",
	OrAssign:         "|=",
	XorAssign:        "^=",
	ShiftLeftAssign:  "<<=",
	ShiftRightAssign: ">>=",
	NandAssign:       "&^=",

	Land: "&&",
	Lor:  "||",
	Inc:  "++",
	Dec:  "--",

	Eq:      "==",
	Less:    "<",
	Greater: ">",
	Assign:  "=",
	Not:     "!",

	Neq:      "!=",
	Leq:      "<=",
	Geq:      ">=",
	Ellipsis: "...",

	Lparen: "(",
	Lbrack: "[",
	Lbrace: "{",
	Comma:  ",",
	Dot:    ".",

	Rparen:    ")",
	Rbrack:    "]",
	Rbrace:    "}",
	Semicolon: ";",
	Colon:     ":",

	Break:       "break",
	Case:        "case",
	Const:       "const",
	Continue:    "continue",
	Default:     "default",
	Else:        "else",
	Fallthrough: "fallthrough",
	For:         "for",
	Func:        "func",
	Goto:        "goto",
	If:          "if",
	Import:      "import",
	Return:      "return",
	Struct:      "struct",
	Switch:      "switch",
	Type:        "type",
	Var:         "var",
}

func TokenStr(t int) string {
	if s, found := tokenStr[t]; found {
		return s
	}
	return fmt.Sprintf("<na-%d>", t)
}

var keywords = func() map[string]int {
	ret := make(map[string]int)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		s := tokenStr[i]
		ret[s] = i
	}
	return ret
}()

func IdentToken(s string) int {
	if i, found := keywords[s]; found {
		return i
	}
	return Ident
}
