package ir1

import (
	"github.com/h8liu/e8/printer"
)

type Stmt interface {
	PrintTo(p printer.Iface)
}

type NoopStmt struct {
	Comment string
}

var _ Stmt = new(NoopStmt)

func NewComment(c string) *NoopStmt {
	ret := new(NoopStmt)
	ret.Comment = c
	return ret
}

func Cm(c string) *NoopStmt { return NewComment(c) }

func (s *NoopStmt) PrintTo(p printer.Iface) {
	p.Printf("// %s", s.Comment)
}
