package token

type Token int

const (
	// misc
	Illegal = Token(iota)
	EOF
	Comment

	// literals
	literalBegin
	Ident
	Int
	Float
	Char
	String
	literalEnd

	// operators
	operatorBegin

	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %

	And        // &
	Or         // |
	Xor        // ^
	ShiftLeft  // <<
	ShiftRight // >>
	Nand       // &^

	AddAssign // +=
	SubAssign // -=
	MulAssign // *=
	DivAssign // /=
	ModAssign // %=

	AndAssign        // &=
	OrAssign         // |=
	XorAssign        // ^=
	ShiftLeftAssign  // <<=
	ShiftRightAssign // >>=
	NandAssign       // &^=

	Land // &&
	Lor  // ||
	Inc  // ++
	Dec  // --

	Eq      // ==
	Less    // <
	Greater // >
	Assign  // =
	Not     // !

	Neq      // !=
	Leq      // <=
	Geq      // >=
	Ellipsis // ...

	Lparen // (
	Lbrack // [
	Lbrace // {
	Comma  // ,
	Dot    // .

	Rparen    // )
	Rbrack    // ]
	Rbrace    // }
	Semicolon // ;
	Colon     // :

	operatorEnd

	// keywords
	keywordBegin

	Break
	Case
	Const
	Continue
	Default
	Else
	Fallthrough
	For
	Func
	Goto
	If
	Import
	Return
	Struct
	Switch
	Type
	Var

	keywordEnd
)
