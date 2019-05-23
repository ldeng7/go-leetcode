func reverseString(s []byte) {
	for i := 0; i < len(s)>>1; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
