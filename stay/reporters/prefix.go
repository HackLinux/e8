package reporters

import (
	"fmt"
	"os"
)

// A reporter that reports to stderr with a filename prefix
type Prefix struct {
	prefix string
}

func (self *Prefix) Report(lineno uint16, off uint8, e error) {
	fmt.Fprintf(os.Stderr, "%s:%d:%d: %v\n",
		self.prefix, lineno, off, e,
	)
}

func NewPrefix(prefix string) *Prefix {
	ret := new(Prefix)
	ret.prefix = prefix
	return ret
}
