package parser_test

import (
	"os"

	"github.com/h8liu/e8/printer"
	"github.com/h8liu/e8/stay/parser"
)

func ExampleImports() {
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

	prog, e := parser.ParseString(testProg)
	if e != nil {
		panic(e)
	}

	pr := printer.New(os.Stdout)
	prog.PrintTo(pr)
	if pr.Error != nil {
		panic(pr.Error)
	}
	// Output:
	// import (
	//     "hello"
	//     fmt "hello"
	//     f "hello"
	//     . "hello"
	//     "ho"
	//     "hello/ho"
	//     "hello"
	//     tx "hello/ho"
	//     "hello"
	//     "ho"
	//     "ho"
	//     "ho"
	//     "ho"
	//     a "ho"
	//     b "ho"
	//     c "ho"
	// )
}
