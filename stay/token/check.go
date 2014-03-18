package token

func IsOperator(t Token) bool {
	return operatorBegin < t && t < operatorEnd
}

func IsKeyword(t Token) bool {
	return keywordBegin < t && t < keywordEnd
}

func IsLiteral(t Token) bool {
	return literalBegin < t && t < literalEnd
}
