package ir1

// a memory structure
type Struct struct {
	vars []*Var
}

func NewStruct() *Struct {
	ret := new(Struct)
	ret.vars = make([]*Var, 0, 8)

	return ret
}
