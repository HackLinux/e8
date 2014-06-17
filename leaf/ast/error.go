package ast

type Error struct {
	e      string
	tokens []Node
}

func NewError(e string) *Error {
	ret := new(Error)
	ret.e = e
	return ret
}

func (self *Error) Add(node Node) {
	self.tokens = append(self.tokens, node)
}
