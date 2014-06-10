package reporter

import (
	"fmt"
	"os"
)

type simpleReporter struct{}

var Simple = new(simpleReporter)
var _ Interface = new(simpleReporter)

func (self *simpleReporter) Report(line int, col int, e error) {
	fmt.Fprintf(os.Stderr, "%d:%d: %v\n", line, col, e)
}
