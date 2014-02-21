package vm

import (
	"testing"
)

func TestDataPage(t *testing.T) {
	p := NewDataPage()

	for i := uint32(0); i < PageSize; i++ {
		p.Write(i, uint8(i))
	}

	for i := uint32(0); i < PageSize; i++ {
		b := p.Read(i)
		if b != uint8(i) {
			t.Fail()
		}
	}
}

func TestAlign(t *testing.T) {
	p := NewDataPage()
	a := &Align{p}
	a.WriteU32(1024+3, 0x01020304)
	i := a.ReadU16(1024 + 1)
	if i != 0x0304 {
		t.Fail()
	}

	b := a.ReadU8(1024 + 1)
	if b != 0x03 {
		t.Fail()
	}
}

func TestMemory(t *testing.T) {
	m := NewMemory()
	p := NewDataPage()
	m.Map(4096, p)
	m.WriteU32(4096 + 1024 + 3, 0x01020304)
	i := m.ReadU16(4096 + 1024 + 3)
	if i != 0x0102 {
		t.Fail()
	}

	m.WriteU32(0x2, 0x13431c32)
	u32 := m.ReadU32(0x3)
	if u32 != 0 {
		t.Fail()
	}
}
