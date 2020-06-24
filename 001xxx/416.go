func numberOfArrays(s string, k int) int {
	l := len(s)
	t := make([]int, l+1)
	t[0] = 1
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			continue
		}
		for j, p := 1, 0; j <= l-i; j++ {
			p = p*10 + int(s[i+j-1]-'0')
			if p > k {
				break
			}
			t[i+j] = (t[i] + t[i+j]) % 1000000007
		}
	}
	return t[l]
}
