package ir1

import (
	"e8vm.net/e8/ir1/decls"
	"e8vm.net/e8/ir1/types"
)

type Package struct {
	Name  string
	Funcs map[string]*Func
}

func NewPackage(name string) *Package {
	ret := new(Package)
	ret.Name = name
	ret.Funcs = make(map[string]*Func)

	return ret
}

func (self *Package) NewFunc(name string, t types.Type) *Func {
	assert(self.Funcs[name] == nil)

	f := NewFunc(name, t)
	f.Pack = self
	self.Funcs[name] = f

	return f
}

func (self *Package) FindCall(name string) decls.Func {
	// TODO: search for imports
	return self.Funcs[name]
}

func (self *Package) FindFunc(name string) *Func {
	return self.Funcs[name]
}
