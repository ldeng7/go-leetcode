func romanToInt(s string) int {
	m := map[byte]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}
	n := 0
	for i, c := range []byte(s) {
		if i == len(s)-1 || m[c] >= m[s[i+1]] {
			n += m[c]
		} else {
			n -= m[c]
		}
	}
	return n
}