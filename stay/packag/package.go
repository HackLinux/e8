package packag

import (
	"fmt"
	"io"
	"os"

	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/parser"
)

type Package struct {
	parser    *parser.Parser
	filenames []string
	files     []*ast.Program
}

const (
	MaxFile = 200
)

func NewPackage() *Package {
	ret := new(Package)
	ret.parser = parser.New()
	ret.filenames = make([]string, 0, MaxFile)
	ret.files = make([]*ast.Program, 0, MaxFile)

	return ret
}

func (self *Package) QueryFilename(id uint8) string {
	return self.filenames[id]
}

func (self *Package) Add(name string, in io.Reader) error {
	nfiles := len(self.files)
	if nfiles == MaxFile {
		return fmt.Errorf("too many files in a package")
	}

	fid := uint8(nfiles)

	self.parser.PosOffset = uint32(fid) << 32
	tree, e := self.parser.Parse(in)
	if e != nil {
		return e // io error on parsing
	}

	self.filenames = append(self.filenames, name)
	self.files = append(self.files, tree)

	return nil
}

func (self *Package) AddFile(path string) error {
	fin, e := os.Open(path)
	if e != nil {
		return e
	}

	defer fin.Close() // read only, so okay to defer

	e = self.Add(path, fin)
	if e != nil {
		return e
	}

	return nil
}
