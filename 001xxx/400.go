func canConstruct(s string, k int) bool {
	l := len(s)
	if l < k {
		return false
	}
	ar := [26]int{}
	for i := 0; i < l; i++ {
		ar[s[i]-'a']++
	}
	c := 0
	for _, a := range ar {
		if a&1 == 1 {
			c++
		}
	}
	return c <= k
}
