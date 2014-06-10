package parser

import (
	"os"
	"strings"

	"github.com/h8liu/e8/leaf/ast"
	"github.com/h8liu/e8/leaf/reporter"
)

func ParseFile(path string) (*ast.Program, error) {
	fin, e := os.Open(path)
	if e != nil {
		return nil, e
	}

	parser := New()
	parser.Reporter = reporter.NewPrefix(path)

	ret, e := parser.Parse(fin)
	if e != nil {
		return nil, e
	}

	e = fin.Close()
	if e != nil {
		return nil, e
	}

	return ret, nil
}

func ParseString(s string) (*ast.Program, error) {
	p := New()
	return p.Parse(strings.NewReader(s))
}
