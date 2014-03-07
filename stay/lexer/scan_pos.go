package lexer

type scanPos struct {
	lineNo     int
	lineOffset int
}

func newScanPos() *scanPos {
	ret := new(scanPos)
	ret.lineNo = 1
	return ret
}

func (self *scanPos) Pos() uint32 {
	return uint32(self.lineNo)<<8 + uint32(self.lineOffset)
}

func (self *scanPos) NewLine() {
	self.lineNo++
	self.lineOffset = 0
}

func (self *scanPos) SyncTo(p *scanPos) {
	self.lineNo = p.lineNo
	self.lineOffset = p.lineOffset
}
