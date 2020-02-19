func breakPalindrome(palindrome string) string {
	l := len(palindrome)
	if l < 2 {
		return ""
	}
	bs := []byte(palindrome)
	for i := 0; i < l>>1; i++ {
		if bs[i] != 'a' {
			bs[i] = 'a'
			return string(bs)
		}
	}
	bs[l-1] = 'b'
	return string(bs)
}
