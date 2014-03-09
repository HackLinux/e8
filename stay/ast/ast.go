package ast

type Program struct {
	Imports []*ImportDecl
	Funcs   []*FuncDecl
}

func NewProgram() *Program {
	ret := new(Program)
	ret.Imports = make([]*ImportDecl, 0, 32)
	ret.Funcs = make([]*FuncDecl, 0, 1024)

	return ret
}

func (self *Program) AddImport(i *ImportDecl) {
	self.Imports = append(self.Imports, i)
}

func (self *Program) AddFunc(f *FuncDecl) {
	self.Funcs = append(self.Funcs, f)
}
