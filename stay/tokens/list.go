package tokens

// List all the keywords
func ListKeywords() []int {
	ret := make([]int, 0, keywordEnd-keywordBegin-1)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		ret = append(ret, i)
	}
	return ret
}

// List all the operators
func ListOperators() []int {
	ret := make([]int, 0, operatorEnd-operatorBegin-1)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		ret = append(ret, i)
	}
	return ret
}
