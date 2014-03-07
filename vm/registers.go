package vm

import (
	"fmt"
	"io"

	"github.com/h8liu/e8/vm/align"
	"github.com/h8liu/e8/vm/inst"
)

// Registers container for e8 ALU. Based on the instruction design, it has 32
// 32-bit integer registers arnd 32 64-bit floating point registers.
// For integer register, $0 is bind to 0. $31 is program counter, and its
// last 2 bits are bind to 0. All other ones are general purpose registers.
type Registers struct {
	regs  []uint32
	fregs []float64
}

// Create new register containers.
func NewRegisters() *Registers {
	ret := new(Registers)

	ret.regs = make([]uint32, inst.Nreg)
	ret.fregs = make([]float64, inst.Nfreg)

	return ret
}

// Read integer register a
func (self *Registers) ReadReg(a uint8) uint32   { return self.regs[a] }

// Read floating point register a
func (self *Registers) ReadFreg(a uint8) float64 { return self.fregs[a] }

// Write integer register a with value v.  Writing to $0 will have no effect,
// writing to $31, the program counter will be automatically aligned.
func (self *Registers) WriteReg(a uint8, v uint32) {
	if a == 0 {
		// do nothing
	} else if a == inst.RegPC {
		self.regs[inst.RegPC] = align.U32(v)
	} else {
		self.regs[a] = v
	}
}

// Write floating point register a with value v
func (self *Registers) WriteFreg(a uint8, v float64) {
	self.fregs[a] = v
}

// Increase $31, program counter by 4.
func (self *Registers) IncPC() uint32 {
	ret := self.regs[inst.RegPC]
	self.regs[inst.RegPC] += 4
	return ret
}

// Print the (integer) register values to an output stream. Useful for debugging.
func (self *Registers) PrintTo(w io.Writer) {
	for i := uint8(0); i < inst.Nreg; i++ {
		fmt.Fprintf(w, "$%02d:%08x", i, self.ReadReg(i))
		if (i+1)%4 == 0 {
			fmt.Fprintln(w)
		} else {
			fmt.Fprint(w, " ")
		}
	}
}
