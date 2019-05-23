func sortArrayByParity(A []int) []int {
	i1 := len(A) - 1
	for i0 := 0; i0 < i1; i0++ {
		if A[i0]&1 == 1 {
			for ; i1 > i0; i1-- {
				if A[i1]&1 == 0 {
					A[i0], A[i1] = A[i1], A[i0]
					i1--
					break
				}
			}
		}
	}
	return A
}
