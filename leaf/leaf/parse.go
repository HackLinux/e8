package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h8liu/e8/leaf/parser"
)

func mainParse(args []string) {
	fset := flag.NewFlagSet("leaf-parse", flag.ExitOnError)
	fset.Parse(args)

	files := fset.Args()

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "no input file.")
		return
	}

	for _, f := range files {
		fmt.Printf("[%s]\n", f)

		tree, errs := parser.ParseTree(f)

		if len(errs) > 0 {
			for _, e := range errs {
				fmt.Fprintln(os.Stderr, e)
			}
		}

		if tree != nil { // might be nil when the file does not exist
			tree.PrintTree(os.Stdout)
		}
	}
}
