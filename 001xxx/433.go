func checkIfCanBreak(s1 string, s2 string) bool {
	c1, c2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		c1[s1[i]-'a']++
	}
	for i := 0; i < len(s2); i++ {
		c2[s2[i]-'a']++
	}
	sum1, sum2, b1, b2 := 0, 0, true, true
	for i := 0; i < 26; i++ {
		sum1 += c1[i]
		sum2 += c2[i]
		if sum1 < sum2 {
			b1 = false
		}
		if sum2 < sum1 {
			b2 = false
		}
		if !(b1 || b2) {
			return false
		}
	}
	return true
}
