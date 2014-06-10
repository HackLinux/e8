// Package modules defines module organizing structure and project making procedure.
package module

import (
	"path/filepath"

	"github.com/h8liu/e8/leaf/parser"
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

func (self *Module) ScanImports() error {
	srcpath := self.meta.srcpath
	files := self.files

	for _, file := range files {
		if file == ".imports" {
			continue
		}

		path := filepath.Join(srcpath, file)

		// println("  ", file)
		_, e := parser.ParseFile(path)
		if e != nil {
			return e
		}

		// TODO:
	}

	self.imports = make([]string, 0)
	return nil
}

func (self *Module) importsPath() string {
	return filepath.Join(self.meta.libpath, "imports")
}

func (self *Module) SaveImports() error {
	return writeLines(self.importsPath(), self.imports)
}

func (self *Module) LoadImports() (e error) {
	self.imports, e = readLines(self.importsPath())
	return e
}
