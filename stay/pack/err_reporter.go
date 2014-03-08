package pack

import (
	"fmt"
	"io"
)

type ErrReporter struct {
	p      *Package
	fid    uint32
	errors []*Error
}

func NewErrReporter(p *Package, fid uint8) *ErrReporter {
	ret := new(ErrReporter)
	ret.p = p
	ret.fid = uint32(fid) << 24
	ret.errors = make([]*Error, 0, 1024)
	return ret
}

func (self *ErrReporter) Report(lineno uint16, off uint8, e error) {
	pos := self.fid + (uint32(lineno) << 8) + uint32(off)
	err := &Error{pos, e}
	self.errors = append(self.errors, err)
}

func (self *ErrReporter) PrintTo(out io.Writer) error {
	filename := self.p.QueryFilename(uint8(self.fid >> 24))

	for _, e := range self.errors {
		_, err := fmt.Fprintln(out, e.Str(filename))
		if err != nil {
			return err
		}
	}

	return nil
}
