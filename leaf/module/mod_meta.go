package module

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type modMeta struct {
	files   []string
	modTime time.Time
}

func loadModMeta(libpath string) (*modMeta, error) {
	p := filepath.Join(libpath, "meta")
	bytes, e := ioutil.ReadFile(p)
	if os.IsNotExist(e) {
		return nil, nil
	}
	if e != nil {
		return nil, e
	}

	ret := new(modMeta)
	lines := strings.Split(string(bytes), "\n")
	if len(lines) == 0 {
		return nil, nil
	}

	ret.modTime, e = time.Parse(time.RFC3339Nano, lines[0])
	if e != nil {
		return nil, e
	}

	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ret.files = append(ret.files, line)
	}
	sort.Strings(ret.files)

	return ret, nil
}

func (self *modMeta) touch() {
	self.modTime = time.Now()
}

func (self *modMeta) fileList() string {
	return strings.Join(self.files, "\n")
}

func (self *modMeta) save(libpath string) error {
	buf := new(bytes.Buffer)
	fmt.Fprintln(buf, self.modTime.Format(time.RFC3339Nano))
	for _, f := range self.files {
		fmt.Fprintln(buf, f)
	}

	p := filepath.Join(libpath, "meta")
	e := os.MkdirAll(libpath, 0755)
	if e != nil {
		return e
	}

	return ioutil.WriteFile(p, buf.Bytes(), 0644)
}

func getListName(name string) string {
	if name == "imports" {
		return ".imports"
	}
	return name
}

func newModMeta(srcpath string) (*modMeta, error) {
	ret := new(modMeta)
	ret.files = make([]string, 0, MaxFile+1)

	files, e := ioutil.ReadDir(srcpath)
	if e != nil {
		return nil, e
	}

	for _, f := range files {
		name := f.Name()
		if name == "imports" || strings.HasSuffix(name, ".stay") {
			p := filepath.Join(srcpath, name)
			modTime, e := fileModTime(p)
			if e != nil {
				return nil, e
			}
			if modTime.After(ret.modTime) {
				ret.modTime = modTime
			}

			ret.files = append(ret.files, getListName(name))
		}
	}

	sort.Strings(ret.files)

	return ret, nil
}
