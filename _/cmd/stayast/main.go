// Parse a stay program file and prints the AST
package main

import (
	"flag"
	"fmt"
	"os"

	"e8vm.net/e8/leaf/parser"
	"e8vm.net/e8/printer"
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

	pr := printer.New(os.Stdout)
	prog.PrintTo(pr)
	noError(pr.Error)
}
