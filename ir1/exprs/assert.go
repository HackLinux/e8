package exprs

func assert(cond bool) {
	if !cond {
		panic("assertion failed")
	}
}
