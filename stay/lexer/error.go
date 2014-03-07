package lexer

// Have to match with a package to create more meaningful error message
type Error struct {
	Pos uint32
	error
}
