package ir1

type Func struct {
	name string  // the function name
	arg  *Struct // structure of func call arguments
	ret  *Struct // structure of return values
	code *Block  // code block
}

func NewFunc(n string) *Func {
	ret := new(Func)
	ret.name = n
	ret.arg = NewStruct()
	ret.ret = NewStruct()
	ret.code = NewBlock()

	return ret
}
