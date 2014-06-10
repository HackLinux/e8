package module

import (
	"fmt"
	"path"
)

type Meta struct {
	path    string
	name    string
	srcpath string
	libpath string

	oldMeta *modMeta
	newMeta *modMeta
}

func NewMeta(p string) (*Meta, error) {
	ret := new(Meta)
	ret.path = p
	ret.name = path.Base(p)
	ret.srcpath = SrcPath(p)
	ret.libpath = LibPath(p)

	chk, e := isDir(ret.srcpath)
	if e != nil {
		return nil, e
	}
	if !chk {
		return nil, fmt.Errorf("%q is not a directory", ret.srcpath)
	}
	if !IsValidName(ret.name) {
		return nil, fmt.Errorf("%q has an invalid module name", p)
	}

	ret.oldMeta, e = loadModMeta(ret.libpath)
	if e != nil {
		return nil, e
	}

	ret.newMeta, e = newModMeta(ret.srcpath)
	if e != nil {
		return nil, e
	}

	return ret, nil
}

func (self *Meta) Save() error {
	if self.Updated() {
		self.newMeta.touch()
		e := self.newMeta.save(self.libpath)
		if e != nil {
			return e
		}
	}

	return nil
}

func (self *Meta) PrintFiles() {
	for _, f := range self.newMeta.files {
		fmt.Println("  ", f)
	}
}

func (self *Meta) Updated() bool {
	if self.oldMeta == nil {
		return true
	}
	if self.newMeta.modTime.After(self.oldMeta.modTime) {
		// some file changed or created after last check
		return true
	}
	if self.newMeta.fileList() != self.oldMeta.fileList() {
		// file list changed after last check
		return true
	}
	return false
}
