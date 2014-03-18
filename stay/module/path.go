package module

import (
	"os"
	"path/filepath"
)

var (
	stayPath    string
	staySrcPath string
	stayLibPath string
)

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
	staySrcPath = filepath.Join(p, "src")
	stayLibPath = filepath.Join(p, "lib")
}

func SrcPath(m string) string { return filepath.Join(staySrcPath, m) }
func LibPath(m string) string { return filepath.Join(stayLibPath, m) }
func StayPath() string        { return stayPath }
