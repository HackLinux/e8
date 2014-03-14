package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/h8liu/e8/stay/module"
)

func parseDir(ms []*module.Module, p string) ([]*module.Module, error) {
	mod, e := module.LoadModule(p)
	if e != nil {
		return ms, e
	}
	fmt.Println(mod.Name)
	ms = append(ms, mod)

	files, e := ioutil.ReadDir(p)
	if e != nil {
		return ms, e
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		ms, e = parseDir(ms, filepath.Join(p, f.Name()))
		if e != nil {
			return ms, e
		}
	}

	return ms, nil
}

func main() {
	p := os.Getenv("STAYPATH")
	if p == "" {
		panic("STAYPATH undefined")
	}

	ms := make([]*module.Module, 0, 100)
	ms, e := parseDir(ms, p)
	if e != nil {
		panic(e)
	}
}
