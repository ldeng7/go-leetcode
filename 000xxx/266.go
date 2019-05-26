func canPermutePalindrome(s string) bool {
	m := map[byte]int{}
	for _, b := range []byte(s) {
		m[b]++
	}
	n := 0
	for _, c := range m {
		if c&1 == 1 {
			n++
			if n > 1 {
				return false
			}
		}
	}
	return true
}
