package reporters

import (
	"fmt"
	"os"
)

type simpleReporter struct{}

var Simple = new(simpleReporter)
var _ ErrReporter = new(simpleReporter)

func (self *simpleReporter) Report(lineno uint16, offset uint8, e error) {
	fmt.Fprintf(os.Stderr, "%d:%d: %v\n", lineno, offset, e)
}
