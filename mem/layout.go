package mem

const SegSize = (1 << 32) / 4

const (
	SegIO uint32 = SegSize * iota
	SegCode
	SegHeap
	SegStack
)
