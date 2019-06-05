func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	o, m := 1, 1
	for {
		m %= K
		if m == 0 {
			return o
		} else {
			o++
			m = m*10 + 1
		}
	}
	return o
}
