package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/h8liu/e8/stay/packag"
)

func parseDir(ps []*packag.Package, p string) ([]*packag.Package, error) {
	pack, e := packag.LoadPackage(p)
	if e != nil {
		return ps, e
	}
	fmt.Println(pack.Name)
	ps = append(ps, pack)

	files, e := ioutil.ReadDir(p)
	if e != nil {
		return ps, e
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		ps, e = parseDir(ps, filepath.Join(p, f.Name()))
		if e != nil {
			return ps, e
		}
	}

	return ps, nil
}

func main() {
	p := os.Getenv("STAYPATH")
	if p == "" {
		panic("STAYPATH undefined")
	}

	ps := make([]*packag.Package, 0, 100)
	ps, e := parseDir(ps, p)
	if e != nil {
		panic(e)
	}
}
