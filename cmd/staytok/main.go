// Breaks a stay file into tokens.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h8liu/e8/leaf/lexer"
	"github.com/h8liu/e8/leaf/reporter"
)

func noError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "error:", e.Error())
		os.Exit(-1)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "requires exactly one input file")
		os.Exit(-1)
	}

	path := args[0]

	fin, e := os.Open(path)
	noError(e)

	lex := lexer.New(fin)
	lex.ReportTo(reporter.NewPrefix(path))

	for lex.Scan() {
		t := lex.Token()
		fmt.Printf("%s:%d:%d: %q - %s\n",
			path, t.Line, t.Col,
			t.Lit, t.Token,
		)
	}

	noError(fin.Close())
}
