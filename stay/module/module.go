package module

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

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
		e = ret.ScanImports()
		if e != nil {
			e = ret.SaveImports()
		}
	} else {
		e = ret.LoadImports()
	}

	if e != nil {
		return nil, e
	}

	maker.addModule(ret)
	return ret, nil
}

func (self *Module) ScanImports() error {
	srcpath := self.meta.srcpath
	files := self.files

	for _, file := range files {
		path := filepath.Join(srcpath, file)
		println(path)
		// TODO:
	}

	self.imports = make([]string, 0)
	return nil
}

func (self *Module) SaveImports() error {
	libpath := self.meta.libpath
	path := filepath.Join(libpath, "imports")

	buf := new(bytes.Buffer)
	for _, im := range self.imports {
		fmt.Fprintln(buf, im)
	}

	return ioutil.WriteFile(path, buf.Bytes(), 0644)
}

func (self *Module) LoadImports() error {
	libpath := self.meta.libpath
	path := filepath.Join(libpath, "imports")
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	s := string(bytes)
	lines := strings.Split(s, "\n")
	ret := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ret = append(ret, line)
	}

	self.imports = ret
	return nil
}
