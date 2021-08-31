func reinitializePermutation(n int) int {
	o := 0
	for t := 2; t != 2 || o == 0; o++ {
		if t <= n/2 {
			t = t*2 - 1
		} else {
			t = (t - n/2) * 2
		}
	}
	return o
}
