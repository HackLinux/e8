package tokens

func IsOperator(t int) bool {
	return operatorBegin < t && t < operatorEnd
}

func IsKeyword(t int) bool {
	return keywordBegin < t && t < keywordEnd
}

func IsLiteral(t int) bool {
	return literalBegin < t && t < literalEnd
}
