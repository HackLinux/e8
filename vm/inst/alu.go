package inst

// ALU is an instruction execution engine. It fulfills the definition of e8
// RISC definition.  Its unexported fields contains buffer for decomposing
// an instruction.
type ALU struct {
	fields *fields
}

// Creates an ALU
func NewALU() *ALU {
	ret := new(ALU)
	ret.fields = new(fields)
	return ret
}

// Executes an instruction, using register and memory interface
// provided by c.
func (self *ALU) Inst(c Core, inst Inst) {
	self.fields.inst = inst
	opInst(c, self.fields)
}
