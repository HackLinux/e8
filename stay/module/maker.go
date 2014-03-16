package module

type Maker struct {
	ForceRebuild bool

	modules map[string]*Module
}

func NewMaker() *Maker {
	ret := new(Maker)
	ret.modules = make(map[string]*Module)
	return ret
}

func (self *Maker) addModule(m *Module) {
	path := m.Path()
	self.modules[path] = m
}

func (self *Maker) FindModule(p string) *Module {
	return self.modules[p]
}
