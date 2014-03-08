package lexer

type ErrReporter interface {
	Report(lineno uint16, offset uint8, e error)
}
