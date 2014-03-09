package tokens

// List all the keywords
func Keywords() []int { return keywordList }

// List all the operators
func Operators() []int { return operatorList }

var keywordList = func() []int {
	ret := make([]int, 0, keywordEnd-keywordBegin-1)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		ret = append(ret, i)
	}
	return ret
}()

var operatorList = func() []int {
	ret := make([]int, 0, operatorEnd-operatorBegin-1)
	for i := operatorBegin + 1; i < operatorEnd; i++ {
		ret = append(ret, i)
	}
	return ret
}()
