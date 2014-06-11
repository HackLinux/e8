package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h8liu/e8/leaf/lexer"
)

func mainLex(args []string) {
	fset := flag.NewFlagSet("leaf-lex", flag.ExitOnError)
	fset.Parse(args)

	files := fset.Args()

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "no input file.")
		return
	}

	onError := func(e error) { fmt.Fprintln(os.Stderr, e) }

	for _, f := range files {
		fmt.Printf("[%s]\n", f)

		fin, e := os.Open(f)
		if e != nil {
			onError(e)
			continue
		}

		lex := lexer.New(fin, f)
		lex.ErrorFunc = onError

		for lex.Scan() {
			tok := lex.Token()
			leading := fmt.Sprintf("%s:%d:%d:", f, tok.Line, tok.Col)
			fmt.Printf("%-20s %-10s - %q\n",
				leading,
				tok.Token, tok.Lit,
			)
		}
	}
}
