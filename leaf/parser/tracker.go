package parser

type trackNode interface{}

type tracker struct {
	stack []*trackLevel
}

type trackLevel struct {
	name string
	subs []trackNode
}

func (level *trackLevel) add(n trackNode) {
	level.subs = append(level.subs, n)
}

func (t *tracker) add(n trackNode) {
	nlevel := len(t.stack)
	assert(nlevel > 0)
	top := t.stack[nlevel-1]
	top.add(n)
}

func (t *tracker) push(s string) {
	level := new(trackLevel)
	level.name = s

	t.add(level)
	t.stack = append(t.stack, level)
}

func (t *tracker) pop() trackNode {
	nlevel := len(t.stack)
	assert(nlevel > 0)
	top := t.stack[nlevel-1]
	t.stack = t.stack[:nlevel-1]
	return top
}
