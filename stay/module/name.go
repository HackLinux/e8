package module

func IsValidName(name string) bool {
	for i, r := range name {
		if r == '_' {
			continue
		}
		if 'a' <= r && r <= 'z' {
			continue
		}
		if '0' <= r && r <= '9' && i > 0 {
			continue
		}
		return false
	}
	return true
}
