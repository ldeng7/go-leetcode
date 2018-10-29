func canTransform(start string, end string) bool {
	l, r := 0, 0
	for i := 0; i < len(start); i++ {
		if start[i] == 'L' {
			l--
		} else if start[i] == 'R' {
			r++
		}
		if end[i] == 'L' {
			l++
		} else if end[i] == 'R' {
			r--
		}
		if l < 0 || r < 0 || (l != 0 && r != 0) {
			return false
		}
	}
	return l == 0 && r == 0
}
