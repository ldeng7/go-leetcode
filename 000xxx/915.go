func partitionDisjoint(A []int) int {
	m, m1, j := A[0], A[0], 0
	for i := 1; i < len(A); i++ {
		a := A[i]
		if a > m {
			m = a
		}
		if a < m1 {
			m1, j = m, i
		}
	}
	return j + 1
}
