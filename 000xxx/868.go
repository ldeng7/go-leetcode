func binaryGap(N int) int {
	o, pi := 0, 31
	for i := 0; i < 31; i++ {
		if ((N >> uint(i)) & 1) == 1 {
			if i-pi > o {
				o = i - pi
			}
			pi = i
		}
	}
	return o
}
