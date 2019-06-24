func sortArrayByParityII(A []int) []int {
	i, j, l := 0, 1, len(A)
	for i < l && j < l {
		if A[i]&1 == 0 {
			i += 2
		}
		if A[j]&1 == 1 {
			j += 2
		}
		if i < l && j < l && A[i]&1 == 1 && A[j]&1 == 0 {
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
