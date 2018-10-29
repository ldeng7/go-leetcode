func longestPalindrome(s string) int {
	if 0 == len(s) {
		return 0
	}
	m := map[byte]int{}
	for _, b := range []byte(s) {
		m[b] = m[b] + 1
	}
	sum := 0
	odd := false
	for _, c := range m {
		if c&1 == 1 {
			odd = true
		}
		sum += c &^ 1
	}
	if odd {
		sum += 1
	}
	return sum
}
