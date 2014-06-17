package parser

import (
	"io"

	"github.com/h8liu/e8/leaf/lexer"
	"github.com/h8liu/e8/prt"
)

type trackNode interface{}

type tracker struct {
	root  *trackLevel
	stack []*trackLevel
}

func (t *tracker) add(n trackNode) {
	_, isLevel := n.(*trackLevel)
	_, isToken := n.(*lexer.Token)
	assert(isLevel || isToken)

	nlevel := len(t.stack)
	assert(nlevel > 0)
	top := t.stack[nlevel-1]
	top.add(n)
}

func (t *tracker) push(s string) {
	level := new(trackLevel)
	level.name = s

	if len(t.stack) == 0 {
		assert(t.root == nil)
		t.root = level
		t.stack = append(t.stack, level)
	} else {
		t.add(level)
		t.stack = append(t.stack, level)
	}
}

func (t *tracker) pop() trackNode {
	nlevel := len(t.stack)
	assert(nlevel > 0)
	top := t.stack[nlevel-1]
	t.stack = t.stack[:nlevel-1]
	return top
}

func (t *tracker) PrintTrack(out io.Writer) {
	if t.root == nil {
		return
	}

	p := prt.New(out)
	p.Indent = "    "

	t.root.PrintTo(p)
}
