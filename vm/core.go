/*
Package vm maintains the basic CPU unit needed for instruction execution.
It defines the interface for a VM core, and also implements the registers in it.
*/
package vm

import (
	"fmt"
	"io"
	"os"

	"github.com/h8liu/e8/vm/inst"
	"github.com/h8liu/e8/vm/mem"
)

type registers struct{ *Registers }
type memory struct{ *mem.Memory }

/*
A VM core, consists a set of 32-bit address memory, and a set of registers.  It
has two anonymous (but private) members of *Registers and *mem.Memory, so it
"inherits" all methods from *Registers and *mem.Memory
*/
type Core struct {
	registers
	memory

	Stdout io.Writer // standard output
	Log    io.Writer // debug logging

	sys *SysPage
	alu *inst.ALU
}

var _ inst.Core = new(Core)

// Creates a core without system page. Output to os.Stdout, no debug logging.
func NewCore() *Core {
	ret := new(Core)
	ret.Registers = NewRegisters()
	ret.Memory = mem.New()
	ret.Stdout = os.Stdout

	ret.alu = inst.NewALU()

	return ret
}

// Same as NewCore(), but with a system page.
func New() *Core {
	ret := NewCore()

	ret.sys = NewSysPage()
	ret.Memory.Map(0, ret.sys)

	return ret
}

// Run one instruction.
func (self *Core) Step() {
	self.sys.ClearError()

	pc := self.IncPC()
	u32 := self.Memory.ReadU32(pc)
	in := inst.Inst(u32)
	if self.Log != nil {
		fmt.Fprintf(self.Log, "%08x: %08x   %v", pc, u32, in)
		if in.Op() != inst.OpJ {
			rs := in.Rs()
			rt := in.Rt()
			rsv := self.ReadReg(rs)
			rtv := self.ReadReg(rt)
			fmt.Fprintf(self.Log, "  ; $%d=%d(%08x) $%d=%d(%08x)",
				rs, rsv, rsv, rt, rtv, rtv)
		}
		fmt.Fprintf(self.Log, "\n")
		// self.Registers.PrintTo(self.Log)
	}
	self.alu.Inst(self, in)

	self.sys.FlushStdout(self.Stdout)
}

// Run at most n instructions. Returns the number of instructions actually
// executed. A core may return early when the core pauses.
func (self *Core) Run(n int) int {
	i := 0
	for i < n {
		self.Step()
		i++

		if self.sys.Halted() {
			break
		}
	}

	return i
}

// Set the program counter. Note the last 2 bits are bind to 0, so
// the program counter will be automatically aligned.
func (self *Core) SetPC(pc uint32) {
	self.Registers.WriteReg(inst.RegPC, pc)
}

// If the core halted.
// Currently, a core can halt gracefully by writing a byte to address 0x4.
// Or it will halt because of writing to address 0x0 to 0x3, which will
// cause the core halts because of an address error.
func (self *Core) Halted() bool { return self.sys.Halted() }

// If the core halted because of an address error.
// Address error currently only occurs when visiting the word at address 0.
func (self *Core) AddrError() bool { return self.sys.AddrError }

// The value when the core halts. This the byte written to address 0x4.
func (self *Core) HaltValue() uint8 { return self.sys.HaltValue }

// Returns if the core rests in peace, which means it halt with a halt value of 0
// (writing a byte 0 to 0x4).
func (self *Core) RIP() bool {
	return self.Halted() && self.HaltValue() == 0 && !self.AddrError()
}
