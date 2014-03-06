package pack

import (
	"fmt"
	"io"
	"os"

	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/parser"
	"github.com/h8liu/e8/stay/pos"
)

type Package struct {
	parser    *parser.Parser
	fileNames []string
	files     []*ast.Ast
}

func NewPackage() *Package {
	ret := new(Package)
	ret.parser = parser.New()
	ret.fileNames = make([]string, 0, pos.MaxFile)
	ret.files = make([]*ast.Ast, 0, pos.MaxFile)

	return ret
}

func (self *Package) Add(name string, in io.Reader) error {
	nfiles := len(self.files)
	if nfiles == pos.MaxFile {
		return fmt.Errorf("too many files in a package")
	}

	id := uint8(nfiles)

	tree, e := self.parser.Parse(id, in)
	if e != nil {
		return e // io error on parsing
	}

	self.fileNames = append(self.fileNames, name)
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
