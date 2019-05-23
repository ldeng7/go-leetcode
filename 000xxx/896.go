func isMonotonic(A []int) bool {
	if len(A) <= 2 {
		return true
	}
	i, d, l := 1, 0, len(A)
	for ; i < l; i++ {
		if d = A[i] - A[i-1]; d != 0 {
			break
		}
	}
	if d > 0 {
		for i = i + 1; i < l; i++ {
			if A[i]-A[i-1] < 0 {
				return false
			}
		}
	} else if d < 0 {
		for i = i + 1; i < l; i++ {
			if A[i]-A[i-1] > 0 {
				return false
			}
		}
	}
	return true
}
