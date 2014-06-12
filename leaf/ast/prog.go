package ast

type Program struct {
	Filename string
	Decls    []Node
}

func (self *Program) AddDecl(n Node) {
	self.Decls = append(self.Decls, n)
}
