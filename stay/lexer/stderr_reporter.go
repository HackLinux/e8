package lexer

import (
	"fmt"
	"os"
)

type stderrReporter struct{}

var StderrReporter = new(stderrReporter)
var _ ErrReporter = new(stderrReporter)

func (self *stderrReporter) Report(lineno uint16, offset uint8, e error) {
	fmt.Fprintf(os.Stderr, "%d:%d: %v\n", lineno, offset, e)
}
