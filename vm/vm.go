/*
Package vm maintains the basic CPU unit needed for instruction execution.
It defines the interface for a VM core, and also implements the registers in it.
*/
package vm

import (
	"fmt"
	"io"
	"os"

	"e8vm.net/e8/inst"
	"e8vm.net/e8/mem"
)

// VM has a core and also has the system page
type VM struct {
	Stdout io.Writer // standard output
	Log    io.Writer // debug logging

	core *Core
	sys  *SysPage
}

func New() *VM {
	ret := new(VM)
	ret.Stdout = os.Stdout

	ret.core = NewCore()
	ret.sys = NewSysPage()

	ret.MapPage(0, ret.sys)

	return ret
}

// Executes one instruction.
func (self *VM) Step() {
	self.sys.Reset()

	pc := self.core.IncPC()
	u32 := self.core.ReadU32(pc)
	in := inst.Inst(u32)
	if self.Log != nil {
		fmt.Fprintf(self.Log, "%08x: %08x   %v", pc, u32, in)
		if in.Op() != inst.OpJ {
			rs := in.Rs()
			rt := in.Rt()
			rsv := self.core.ReadReg(rs)
			rtv := self.core.ReadReg(rt)
			fmt.Fprintf(self.Log, "  ; $%d=%d(%08x) $%d=%d(%08x)",
				rs, rsv, rsv, rt, rtv, rtv)
		}
		fmt.Fprintf(self.Log, "\n")
		// self.Registers.PrintTo(self.Log)
	}
	inst.Exec(self.core, in)

	self.sys.FlushStdout(self.Stdout)
}

// Executes at most n instructions. Returns the number of instructions actually
// executed. A core may return early when the core halts.
func (self *VM) Run(n int) int {
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
func (self *VM) SetPC(pc uint32) {
	self.core.WriteReg(inst.RegPC, pc)
}

// If the core halted.
// Currently, a core can halt gracefully by writing a byte to address 0x4.
// Or it will halt because of writing to address 0x0 to 0x7, which will
// cause the core halts because of an address error.
func (self *VM) Halted() bool { return self.sys.Halted() }

// If the core halted because of an address error.
// Address error currently only occurs when visiting the word at address 0.
func (self *VM) AddrError() bool { return self.sys.AddrError }

// The value when the core halts. This the byte written to address 0x4.
func (self *VM) HaltValue() uint8 { return self.sys.HaltValue }

// Returns if the core rests in peace, which means it halt with a halt value of 0
// (writing a byte 0 to 0x4).
func (self *VM) RIP() bool {
	return self.Halted() && self.HaltValue() == 0 && !self.AddrError()
}

// Checks if a page is valid
func (self *VM) CheckPage(addr uint32) bool {
	return self.core.Check(addr)
}

// Maps a page at particular address
func (self *VM) MapPage(addr uint32, p mem.Page) {
	self.core.Map(addr, p)
}

// Prints registers to out
func (self *VM) DumpRegs(out io.Writer) {
	self.core.PrintTo(out)
}
