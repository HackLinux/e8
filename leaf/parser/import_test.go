package parser_test

import (
	"bytes"
	"testing"

	"github.com/h8liu/e8/printer"
	"github.com/h8liu/e8/stay/parser"
)

const testProg = `
// prints hello

import "hello"

import fmt "hello"
import f "hello"
import . "hello"

/** dafe */

import (
	"ho"
	"hello/ho"
)

import ()

import (



)

import "hello"

import (
	// t2 "ho"
	tx "hello/ho"
)

import "hello"
import "ho"

import ( "ho"; "ho"; "ho" )
import ( a "ho"; b "ho"; c "ho"; )

//so
`

const expectOutput = `import (
	"hello"
	fmt "hello"
	f "hello"
	. "hello"
	"ho"
	"hello/ho"
	"hello"
	tx "hello/ho"
	"hello"
	"ho"
	"ho"
	"ho"
	"ho"
	a "ho"
	b "ho"
	c "ho"
)
`

func TestImports(t *testing.T) {

	prog, e := parser.ParseString(testProg)
	if e != nil {
		t.Fatal(e)
		return
	}

	buf := new(bytes.Buffer)
	pr := printer.New(buf)
	prog.PrintTo(pr)
	if pr.Error != nil {
		t.Fatal(pr.Error)
	}

	s := buf.String()
	if s != expectOutput {
		t.Fatal(s)
	}
}