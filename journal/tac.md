// tac design

// global heap data
data {
    a u32 = u32 0
    something u32[2] = {0, 2}
    pointer u32[4]
    _ u32 // padding
}

func foo (
	ret u32
    x u32
    y u32
    f f64
) {
	t u32 = 0 
	f i32 = -1
	p ptr = 0
	_1 = x + y // type auto infered
	_2 = x - y
	_3 = x * y
	_3 = x / y
	_3 = x % y
	_4 = &x // load the pointer
	_5 u32 = *_4 // interpret the pointed value as a uint32
	_6 ptr = *_4 // interpret the pointed value as pointer
	
    // if this is constant, then use j
    // otherwise, use 
	push x // push the arg
	push _4 // push another arg
	push _ // padding
    call bar

	ret = x / y
}
