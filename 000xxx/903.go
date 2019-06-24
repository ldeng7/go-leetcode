func numPermsDISequence(S string) int {
	l := len(S)
	t := make([]int, l+1)
	for i := 0; i <= l; i++ {
		t[i] = 1
	}
	for i := 0; i < l; i++ {
		if S[i] == 'I' {
			t = t[:len(t)-1]
			for j := 1; j < len(t); j++ {
				t[j] = (t[j] + t[j-1]) % 1000000007
			}
		} else {
			t = t[1:]
			for j := len(t) - 2; j >= 0; j-- {
				t[j] = (t[j] + t[j+1]) % 1000000007
			}
		}
	}
	return t[0]
}
