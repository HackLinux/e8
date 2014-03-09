package lexer

type Token struct {
	Token int
	Line  int
	Col   int
	Lit   string
}
