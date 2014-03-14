package module

import (
	"fmt"
	"io"
)

type Reporter struct {
	p      *Module
	fid    uint32
	errors []*Error
}

func NewReporter(p *Module, fid uint8) *Reporter {
	ret := new(Reporter)
	ret.p = p
	ret.fid = uint32(fid) << 24
	ret.errors = make([]*Error, 0, 1024)
	return ret
}

func (self *Reporter) Report(line uint16, col uint8, e error) {
	pos := self.fid + (uint32(line) << 8) + uint32(col)
	err := &Error{pos, e}
	self.errors = append(self.errors, err)
}

func (self *Reporter) PrintTo(out io.Writer) error {
	filename := self.p.QueryFilename(uint8(self.fid >> 24))

	for _, e := range self.errors {
		_, err := fmt.Fprintln(out, e.Str(filename))
		if err != nil {
			return err
		}
	}

	return nil
}
