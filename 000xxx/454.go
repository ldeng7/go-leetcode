func fourSumCount(A []int, B []int, C []int, D []int) int {
	out := 0
	m1, m2 := map[int]int{}, map[int]int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			m1[A[i]+B[j]]++
			m2[C[i]+D[j]]++
		}
	}
	for i, v := range m1 {
		out += v * m2[-i]
	}
	return out
}
