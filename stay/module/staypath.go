package module

import (
	"os"
	"path/filepath"
)

var stayPath string

func parseStayPath() string {
	p := os.Getenv("STAYPATH")
	if p == "" {
		return ""
	}

	p, e := filepath.Abs(p)
	if e != nil {
		return ""
	}

	p, e = filepath.EvalSymlinks(p)
	if e != nil {
		return ""
	}

	return p
}

func init() {
	p := parseStayPath()
	if p == "" {
		var e error
		p, e = os.Getwd()
		if e != nil {
			panic(e)
		}
	}

	stayPath = p
}

func StayPath() string {
	return stayPath
}
