package parser

import (
	"io"

	"github.com/h8liu/e8/stay/ast"
)

type Parser struct {
}

func New() *Parser {
	ret := new(Parser)
	panic("todo")

	return ret
}

func (self *Parser) Parse(id uint8, in io.Reader) (*ast.Ast, error) {
	panic("todo")
}
