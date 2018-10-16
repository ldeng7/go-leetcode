func isPalindrome(s string) bool {
	bs := []byte{}
	for _, c := range []byte(s) {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			bs = append(bs, c)
		}
		if c >= 'A' && c <= 'Z' {
			bs = append(bs, c+('a'-'A'))
		}
	}
	for i := 0; i < (len(bs) >> 1); i++ {
		if bs[i] != bs[len(bs)-1-i] {
			return false
		}
	}
	return true
}
