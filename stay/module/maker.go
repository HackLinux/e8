package module

type Maker struct {
	ForceRebuild     bool
	NotFollowImports bool

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

func (self *Maker) Get(p string) *Module {
	return self.modules[p]
}

func (self *Maker) Add(p string) (*Module, error) {
	m := self.Get(p)
	if m != nil {
		return m, nil
	}

	ret := new(Module)
	ret.path = p

	var e error
	ret.meta, e = NewMeta(p)
	if e != nil {
		return nil, e
	}
	ret.files = ret.meta.newMeta.files

	if self.ForceRebuild || ret.meta.Updated() {
		e = ret.meta.Save()
		if e != nil {
			return nil, e
		}
		e = ret.ScanImports()
		if e == nil {
			e = ret.SaveImports()
		}
	} else {
		e = ret.LoadImports()
	}

	if e != nil {
		return nil, e
	}

	self.addModule(ret)

	if !self.NotFollowImports {
		for _, imp := range ret.imports {
			_, e := self.Add(imp)
			if e != nil {
				return nil, e
			}
		}
	}

	return ret, nil
}

func (self *Maker) Make() {
	// TODO:
}
