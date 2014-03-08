package vm

import (
	"github.com/h8liu/e8/vm/inst"
	"github.com/h8liu/e8/vm/mem"
	"github.com/h8liu/e8/vm/regs"
)

/*
A VM core, consists a set of 32-bit address memory, and a set of registers.  It
has two anonymous (but private) members of *Registers and *mem.Memory, so it
"inherits" all methods from *Registers and *mem.Memory
*/
type Core struct {
	*regs.Registers
	*mem.Memory
}

var _ inst.Core = new(Core)

// Creates a core without system page. Output to os.Stdout, no debug logging.
func NewCore() *Core {
	ret := new(Core)
	ret.Registers = regs.New(inst.Nreg, inst.Nreg)
	ret.Memory = mem.New()

	return ret
}
