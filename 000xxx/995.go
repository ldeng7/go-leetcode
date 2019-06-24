func minKBitFlips(A []int, K int) int {
	o, j := 0, 0
	for i, a := range A {
		if i >= K && A[i-K] == 2 {
			j--
		}
		if j&1 == a {
			if i+K > len(A) {
				return -1
			}
			A[i], o, j = 2, o+1, j+1
		}
	}
	return o
}
