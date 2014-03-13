package packag

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/h8liu/e8/stay/ast"
	"github.com/h8liu/e8/stay/parser"
	"github.com/h8liu/e8/stay/reporter"
)

type Package struct {
	Name      string
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

var stayPath string
func init() {
	stayPath = os.Getenv("STAYPATH")
	n := len(stayPath)
	if n > 0 && stayPath[n-1] == os.PathSeparator {
		stayPath = stayPath[:n-1]
	}
}

func packagePath(p string) (string, error) {
	p, e := filepath.Abs(p)
	if e != nil {
		return "", e
	}
	p, e = filepath.EvalSymlinks(p)
	if e != nil {
		return "", e
	}

	p = filepath.Clean(p)

	if !strings.HasPrefix(p, stayPath) {
		return "", fmt.Errorf("not in STAYPATH")
	}

	p = p[len(stayPath):]
	return p, nil
}

func LoadPackage(p string) (*Package, error) {
	p, e := packagePath(p)
	if e != nil {
		return nil, e
	}

	ret := NewPackage()
	ret.Name = p
	if !strings.HasSuffix(p, string(os.PathSeparator)) {
		p = p + string(os.PathSeparator)
	}

	abs := filepath.Join(stayPath, p)

	files, e := ioutil.ReadDir(abs)
	if e != nil {
		return nil, e
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		name := f.Name()
		if !strings.HasSuffix(name, ".stay") {
			continue
		}

		e = ret.AddFile(filepath.Join(abs, name))
		if e != nil {
			return nil, e
		}
	}

	return ret, nil
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
	self.parser.Reporter = reporter.NewPrefix(name)

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
