func validWordAbbreviation(word string, abbr string) bool {
	lw := len(word)
	j, c := 0, 0
	for _, b := range []byte(abbr) {
		if b >= '0' && b <= '9' {
			if b == '0' && c == 0 {
				return false
			}
			c = c*10 + int(b-'0')
		} else {
			j += c
			if j >= lw || b != word[j] {
				return false
			}
			j, c = j+1, 0
		}
	}
	return j+c == lw
}
