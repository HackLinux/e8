package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("todo: intro")
		return
	}

	cmd := args[1]
	args = args[2:]

	switch cmd {
	case "dasm":
		mainDasm(args)
	case "run":
		mainRun(args)
	default:
		fmt.Fprintf(os.Stderr, "e8: unknown subcommand %q.\n", cmd)
		fmt.Fprintf(os.Stderr, "Run 'e8 help' for usage.\n")
	}

	// we need several sub commands here
	/*
		we need a bunch of sub command here
		- help
		- asm
		- dasm
		- run
	*/
}

func printError(e error) {
	fmt.Fprintf(os.Stderr, "error: %s", e)
}
