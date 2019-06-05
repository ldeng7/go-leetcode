func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}
	u0 := isUpper(word[0])
	if u0 {
		u1 := isUpper(word[1])
		for i := 2; i < len(word); i++ {
			if isUpper(word[i]) != u1 {
				return false
			}
		}
	} else {
		for i := 1; i < len(word); i++ {
			if isUpper(word[i]) {
				return false
			}
		}
	}
	return true
}
