package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h8liu/e8/stay/lexer"
	"github.com/h8liu/e8/stay/tokens"
)

func noError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "error:", e.Error())
		os.Exit(-1)
	}
}

func posString(p uint32) string {
	line, off := p>>8, p&0xff
	return fmt.Sprintf("%d:%d", line, off)
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
	for {
		to, pos, lit := lex.Scan()
		if to == tokens.EOF {
			break
		}

		fmt.Printf("%s:%s: %q - %s\n",
			path, posString(pos),
			lit, tokens.TokenStr(to),
		)
	}

	noError(lex.Err())

	es := lex.LexErrors()
	for _, e := range es {
		fmt.Fprintf(os.Stderr, "%s:%s: %s\n",
			path, posString(e.Pos), e,
		)
	}

	noError(fin.Close())
}
