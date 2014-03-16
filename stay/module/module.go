package module

const (
	MaxFile = 200
)

type Module struct {
	path    string
	meta    *Meta
	files   []string
	imports []string
}

func (self *Module) Path() string {
	return self.path
}

func (maker *Maker) Open(p string) (*Module, error) {
	m := maker.FindModule(p)
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

	if maker.ForceRebuild || ret.meta.Updated() {
		e = ret.meta.Save()
		if e != nil {
			return nil, e
		}
		ret.imports, e = ScanImports(ret.meta.srcpath, ret.files)
	} else {
		ret.imports, e = LoadImports(ret.meta.libpath)
	}

	if e != nil {
		return nil, e
	}

	maker.addModule(ret)
	return ret, nil
}
