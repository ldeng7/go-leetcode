func numKLenSubstrNoRepeats(S string, K int) int {
	o, l, c := 0, len(S), [26]int{}
	if K > l {
		return 0
	}
	for i := 0; i < K; i++ {
		c[S[i]-'a']++
	}

	cal := func() {
		for i := 0; i < 26; i++ {
			if c[i] > 1 {
				return
			}
		}
		o++
	}

	cal()
	for i := K; i < l; i++ {
		c[S[i]-'a']++
		c[S[i-K]-'a']--
		cal()
	}
	return o
}
