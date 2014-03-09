package ast

type Program struct {
	Imports []*ImportDecl
}

func NewProgram() *Program {
	ret := new(Program)
	ret.Imports = make([]*ImportDecl, 0, 32)

	return ret
}

func (self *Program) AddImport(i *ImportDecl) {
	self.Imports = append(self.Imports, i)
}
