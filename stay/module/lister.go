package module

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Lister struct {
	Error          error
	ErrorHandler   func(e error)
	WarningHandler func(e error)

	list []string
}

func NewLister() *Lister {
	ret := new(Lister)
	ret.ErrorHandler = func(e error) {
		fmt.Fprintln(os.Stderr, "error:", e)
	}

	ret.WarningHandler = func(e error) {
		fmt.Fprintln(os.Stderr, "warning:", e)
	}

	return ret
}

func ListModules() ([]string, error) {
	lister := NewLister()
	return lister.List(), lister.Error
}

func (self *Lister) List() []string {
	self.list = make([]string, 0, 1024)
	self.scan("")

	ret := self.list
	self.list = nil
	return ret
}

func (self *Lister) err(e error) {
	self.Error = e
	self.ErrorHandler(e)
}

func (self *Lister) warn(e error) {
	self.WarningHandler(e)
}

func isDir(p string) (bool, error) {
	stat, e := os.Stat(p)
	if e != nil {
		return false, e
	}
	return stat.Mode().IsDir(), nil
}

func (self *Lister) scan(m string) {
	p := SrcPath(m)
	files, e := ioutil.ReadDir(p)
	if e != nil {
		self.err(e)
		return
	}

	for _, f := range files {
		name := f.Name()
		subm := path.Join(m, name)
		subp := SrcPath(subm)

		chk, e := isDir(subp)
		if e != nil {
			self.err(e)
		}
		if !chk {
			continue
		}
		if !IsValidName(name) {
			self.warn(fmt.Errorf("path %q has an invalid module name", subm))
			continue
		}

		self.list = append(self.list, subm)
		self.scan(subm)
	}
}
