// Parse a stay program file and prints the AST
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h8liu/e8/stay/parser"
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
	prog, e := parser.ParseFile(path)
	noError(e)

	for _, imp := range prog.Imports {
		if imp.As == "" {
			fmt.Printf("import %q\n", imp.Path)
		} else {
			fmt.Printf("import %s %q\n", imp.As, imp.Path)
		}
	}
}
