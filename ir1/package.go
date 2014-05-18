package ir1

type Package struct {
	Name  string
	Funcs map[string]*Func
}

func P(name string) *Package {
	ret := new(Package)
	ret.Name = name
	ret.Funcs = make(map[string]*Func)

	return ret
}

func (self *Package) F(name string, t Type) *Func {
	assert(self.Funcs[name] == nil)

	f := F(name, t)
	f.pack = self
	self.Funcs[name] = f

	return f
}

func (self *Package) FindCall(name string) FuncDecl {
	// TODO: search for imports
	return self.Funcs[name]
}

func (self *Package) FindFunc(name string) *Func {
	return self.Funcs[name]
}
