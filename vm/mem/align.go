package mem

import (
	"e8vm.net/e8/vm/align"
)

// A page wrapper that takes a page and perform aligned reads
// and writes in a page.
// Byte order is little endian.
// Read and writes are performed byte by byte,
// where lower bytes are written/read first.
//
// If the offset is not properly aligned, it will be aligned down
// automatically.
type Align struct {
	Page
}

func maskOffset(offset uint32) uint32 { return offset & PageMask }

func offset8(offset uint32) uint32 {
	return maskOffset(offset)
}

func offset16(offset uint32) uint32 {
	return align.A16(maskOffset(offset))
}

func offset32(offset uint32) uint32 {
	return align.A32(maskOffset(offset))
}

func offset64(offset uint32) uint32 {
	return align.A64(maskOffset(offset))
}

func (self *Align) WriteU8(offset uint32, value uint8) {
	self.writeU8(offset8(offset), value)
}

func (self *Align) WriteU16(offset uint32, value uint16) {
	self.writeU16(offset16(offset), value)
}

func (self *Align) WriteU32(offset uint32, value uint32) {
	self.writeU32(offset32(offset), value)
}

func (self *Align) WriteF64(offset uint32, value float64) {
	panic("todo")
}

func (self *Align) ReadU8(offset uint32) uint8 {
	return self.readU8(offset8(offset))
}

func (self *Align) ReadU16(offset uint32) uint16 {
	return self.readU16(offset16(offset))
}

func (self *Align) ReadU32(offset uint32) uint32 {
	return self.readU32(offset32(offset))
}

func (self *Align) ReadF64(offset uint32) float64 {
	panic("todo")
}

func (self *Align) writeU8(offset uint32, value uint8) {
	self.Page.Write(offset, value)
}

func (self *Align) writeU16(offset uint32, value uint16) {
	self.Page.Write(offset, uint8(value))
	self.Page.Write(offset+1, uint8(value>>8))
}

func (self *Align) writeU32(offset uint32, value uint32) {
	self.Page.Write(offset, uint8(value))
	self.Page.Write(offset+1, uint8(value>>8))
	self.Page.Write(offset+2, uint8(value>>16))
	self.Page.Write(offset+3, uint8(value>>24))
}

func (self *Align) readU8(offset uint32) uint8 {
	return self.Page.Read(offset)
}

func (self *Align) readU16(offset uint32) uint16 {
	ret := uint16(self.Page.Read(offset))
	ret |= uint16(self.Page.Read(offset+1)) << 8
	return ret
}

func (self *Align) readU32(offset uint32) uint32 {
	ret := uint32(self.Page.Read(offset))
	ret |= uint32(self.Page.Read(offset+1)) << 8
	ret |= uint32(self.Page.Read(offset+2)) << 16
	ret |= uint32(self.Page.Read(offset+3)) << 24
	return ret
}
