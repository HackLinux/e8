package main

import (
	"flag"
)

func mainParse(args []string) {
	fset := flag.NewFlagSet("leaf-parse", flag.ExitOnError)
	fset.Parse(args)

	// TODO:
}
