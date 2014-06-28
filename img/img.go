package img

import (
	"fmt"
	"io"
	"os"

	"e8vm.net/e8/mem"
	"e8vm.net/e8/vm"
)

func Make(in io.Reader) (*vm.VM, error) {
	v := vm.New()
	e := LoadInto(v, in)
	if e != nil {
		return nil, e
	}
	return v, nil
}

func Open(path string) (*vm.VM, error) {
	fin, e := os.Open(path)
	if e != nil {
		return nil, e
	}

	defer fin.Close()

	return Make(fin)
}

func LoadInto(c *vm.VM, in io.Reader) error {
	var p mem.Page
	cur := uint32(0)

	for {
		header, buf, e := Read(in)
		if e == io.EOF {
			return nil
		}
		if e != nil {
			return e
		}

		for i, b := range buf {
			addr := header.addr + uint32(i)
			id := mem.PageId(addr)
			if id == 0 {
				return fmt.Errorf("attempt to map system page")
			}

			if cur == 0 || cur != id {
				cur = id
				if !c.CheckPage(addr) {
					p = mem.NewPage()
					c.MapPage(addr, p)
				}
			}

			p.Write(addr&mem.PageMask, b)
		}
	}

	return nil
}

func Read(in io.Reader) (*Header, []byte, error) {
	header := new(Header)
	e := header.ReadIn(in)
	if e != nil {
		return nil, nil, e
	}

	buf := make([]byte, header.size)
	_, e = io.ReadFull(in, buf)
	if e == io.EOF {
		e = io.ErrUnexpectedEOF
	}
	return header, buf, e
}
