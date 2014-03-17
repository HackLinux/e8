package module

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/h8liu/e8/stay/parser"
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
