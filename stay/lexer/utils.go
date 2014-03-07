package lexer

func isLetter(r rune) bool {
	if 'a' <= r && r <= 'z' {
		return true
	}
	if 'A' <= r && r <= 'Z' {
		return true
	}
	if r == '_' {
		return true
	}
	return false
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isWhite(r rune) bool {
	return r == ' ' || r == '\r' || r == '\t'
}
