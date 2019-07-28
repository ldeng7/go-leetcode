func isPalindrome(s string) bool {
	l, bs := len(s), []byte{}
	for i := 0; i < l; i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			bs = append(bs, c)
		}
		if c >= 'A' && c <= 'Z' {
			bs = append(bs, c+('a'-'A'))
		}
	}
	l = len(bs)
	for i := 0; i < l>>1; i++ {
		if bs[i] != bs[l-1-i] {
			return false
		}
	}
	return true
}
