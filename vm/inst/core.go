package inst

// Defines the operations that a CPU box for instruction needs to implement
type Core interface {
	// Integer register operations
	WriteReg(a uint8, v uint32)
	ReadReg(a uint8) uint32

	// Floating point register operations
	WriteFloatReg(a uint8, v float64)
	ReadFloatReg(a uint8) float64

	// Memory operations
	WriteU8(addr uint32, v uint8)
	WriteU16(addr uint32, v uint16)
	WriteU32(addr uint32, v uint32)
	WriteF64(addr uint32, v float64)

	ReadU8(addr uint32) uint8
	ReadU16(addr uint32) uint16
	ReadU32(addr uint32) uint32
	ReadF64(addr uint32) float64
}

const (
	Nreg  = 32       // nunber of integer registers
	RegPC = Nreg - 1 // the index of program counter
)
