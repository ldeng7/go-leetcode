func repeatedNTimes(A []int) int {
	l := len(A)
	for i := 0; i < l-2; i++ {
		a := A[i]
		if a == A[i+1] || a == A[i+2] {
			return a
		}
	}
	return A[l-1]
}
