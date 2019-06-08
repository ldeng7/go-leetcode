func lenLongestFibSubseq(A []int) int {
	o, l := 0, len(A)
	m := map[int]int{}
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
	}
	for i, a := range A {
		m[a] = i
		for j := 0; j < i; j++ {
			b := a - A[j]
			k, ok := m[b]
			if !ok {
				k = -1
			}
			if b < A[j] && k >= 0 {
				t[j][i] = t[k][j] + 1
			} else {
				t[j][i] = 2
			}
			if t[j][i] > o {
				o = t[j][i]
			}
		}
	}
	if o > 2 {
		return o
	}
	return 0
}
