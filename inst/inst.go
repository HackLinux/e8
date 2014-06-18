package inst

type Inst uint32

const (
	OpShift    = 26
	RsShift    = 21
	RtShift    = 16
	RdShift    = 11
	ShamtShift = 6

	OpMask    = 0x3f << OpShift
	RsMask    = 0x1f << RsShift
	RtMask    = 0x1f << RtShift
	RdMask    = 0x1f << RdShift
	ShamtMask = 0x1f << ShamtShift
	FunctMask = 0x3f
	ImMask    = 0xffff

	Nfunct = 64
	Nop    = 64
)

// Returns the uint32 representation
func (i Inst) U32() uint32 { return uint32(i) }

// Returns the op field
func (i Inst) Op() uint8 { return uint8(i >> 26) }

// Returns the rs field
func (i Inst) Rs() uint8 { return uint8(i>>21) & 0x1f }

// Returns the rt field
func (i Inst) Rt() uint8 { return uint8(i>>16) & 0x1f }

// Returns the rd field
func (i Inst) Rd() uint8 { return uint8(i>>11) & 0x1f }

// Returns the shamt field
func (i Inst) Sh() uint8 { return uint8(i>>6) & 0x1f }

// Returns the funct field
func (i Inst) Fn() uint8 { return uint8(i) & 0x3f }

// Returns the immediate (16-bit) field as an unsigned int
func (i Inst) Imu() uint16 { return uint16(i) }

// Returns the immediate (16-bit) field as an signed int
func (i Inst) Ims() int16 { return int16(uint16(i)) }

// Returns the address field
func (i Inst) Ad() int32 { return int32(i) << 6 >> 6 }

type instFunc func(c Core, i Inst)

func makeInstList(m map[uint8]instFunc, n uint8) []instFunc {
	ret := make([]instFunc, n)
	for i := range ret {
		ret[i] = opNoop
	}
	for i, inst := range m {
		ret[i] = inst
	}
	return ret
}

const (
	OpRinst = 0
	OpJ     = 0x02
	OpBeq   = 0x04
	OpBne   = 0x05

	OpAddi = 0x08
	OpSlti = 0x0A
	OpAndi = 0x0C
	OpOri  = 0x0D
	OpLui  = 0x0F

	OpLw  = 0x23
	OpLhs = 0x21
	OpLhu = 0x25
	OpLbs = 0x20
	OpLbu = 0x24
	OpSw  = 0x2B
	OpSh  = 0x29
	OpSb  = 0x28
)

var instList = makeInstList(
	map[uint8]instFunc{
		OpRinst: opRinst,
		OpJ:     opJ,
		OpBeq:   opBeq,
		OpBne:   opBne,

		OpAddi: opAddi,
		OpSlti: opSlti,
		OpAndi: opAndi,
		OpOri:  opOri,
		OpLui:  opLui,

		OpLw:  opLw,
		OpLhs: opLhs,
		OpLhu: opLhu,
		OpLbs: opLbs,
		OpLbu: opLbu,
		OpSw:  opSw,
		OpSh:  opSh,
		OpSb:  opSb,
	}, Nop,
)

const (
	FnAdd = 0x20
	FnSub = 0x22
	FnAnd = 0x24
	FnOr  = 0x25
	FnXor = 0x26
	FnNor = 0x27
	FnSlt = 0x2A

	FnMul  = 0x18
	FnMulu = 0x19
	FnDiv  = 0x1A
	FnDivu = 0x1B
	FnMod  = 0x1C
	FnModu = 0x1D

	FnSll  = 0x00
	FnSrl  = 0x02
	FnSra  = 0x03
	FnSllv = 0x04
	FnSrlv = 0x06
	FnSrav = 0x07
)

var rInstList = makeInstList(
	map[uint8]instFunc{
		FnAdd: opAdd,
		FnSub: opSub,
		FnAnd: opAnd,
		FnOr:  opOr,
		FnXor: opXor,
		FnNor: opNor,
		FnSlt: opSlt,

		FnMul:  opMul,
		FnMulu: opMulu,
		FnDiv:  opDiv,
		FnDivu: opDivu,
		FnMod:  opMod,
		FnModu: opModu,

		FnSll:  opSll,
		FnSrl:  opSrl,
		FnSra:  opSra,
		FnSllv: opSllv,
		FnSrlv: opSrlv,
		FnSrav: opSrav,
	}, Nfunct,
)

// Executes an instruction.
func Exec(c Core, i Inst) { instList[i.Op()](c, i) }

func opRinst(c Core, i Inst) { rInstList[i.Fn()](c, i) }

func opJ(c Core, i Inst) {
	pc := c.ReadReg(RegPC)
	c.WriteReg(RegPC, pc+uint32(int32(i<<6)>>4))
}

func opNoop(c Core, i Inst) {}
