package inst

// Defines the operations that a CPU box for instruction needs to implement
type Core interface {
	// integer register operations
	WriteReg(a uint8, v uint32)
	ReadReg(a uint8) uint32

	// floating point register operations
	WriteFloatReg(a uint8, v float64)
	ReadFloatReg(a uint8) float64

	// memory operations
	WriteU8(addr uint32, v uint8)
	WriteU16(addr uint32, v uint16)
	WriteU32(addr uint32, v uint32)

	ReadU8(addr uint32) uint8
	ReadU16(addr uint32) uint16
	ReadU32(addr uint32) uint32
}

const (
	NintReg   = 32          // nunber of integer registers
	NfloatReg = NintReg     // number of floating point registers
	RegPC     = NintReg - 1 // the index of program counter
)
