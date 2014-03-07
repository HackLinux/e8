package align

func A16(offset uint32) uint32 { return offset >> 1 << 1 }
func A32(offset uint32) uint32 { return offset >> 2 << 2 }
func A64(offset uint32) uint32 { return offset >> 3 << 3 }
