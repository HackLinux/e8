package ast

type Ast struct {
}

type ImportDecl struct {
	As   string
	Path string
	Pos  uint32
}

type Program struct {
	Imports []*ImportDecl
}

func NewProgram() *Program {
	ret := new(Program)
	ret.Imports = make([]*ImportDecl, 32)

	return ret
}

func (self *Program) AddImport(i *ImportDecl) {
	self.Imports = append(self.Imports, i)
}
