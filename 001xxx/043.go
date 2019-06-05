func maxSumAfterPartitioning(A []int, K int) int {
	l := len(A)
	t := make([]int, l)
	for i, a := range A {
		for j := 0; j < K && i-j >= 0; j++ {
			a1 := A[i-j]
			if a1 > a {
				a = a1
			}
			t1 := (j + 1) * a
			if i-j >= 1 {
				t1 += t[i-j-1]
			}
			if t1 > t[i] {
				t[i] = t1
			}
		}
	}
	return t[l-1]
}
