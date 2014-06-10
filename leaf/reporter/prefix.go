package reporter

import (
	"fmt"
	"os"
)

// A reporter that reports to stderr with a filename prefix
type Prefix struct {
	prefix string
}

func (self *Prefix) Report(line int, col int, e error) {
	fmt.Fprintf(os.Stderr, "%s:%d:%d: %v\n",
		self.prefix, line, col, e,
	)
}

func NewPrefix(prefix string) *Prefix {
	ret := new(Prefix)
	ret.prefix = prefix
	return ret
}
