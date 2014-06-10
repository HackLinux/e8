package ir1

func assert(cond bool) {
	if !cond {
		panic("assertion failed")
	}
}
