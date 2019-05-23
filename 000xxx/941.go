func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	i, j := 0, len(A)-1
	for ; i < len(A)-1; i++ {
		if A[i+1] == A[i] {
			return false
		} else if A[i+1] < A[i] {
			break
		}
	}
	for ; j > 0; j-- {
		if A[j-1] == A[j] {
			return false
		} else if A[j-1] < A[j] {
			break
		}
	}
	return i == j && i != 0 && j != len(A)-1
}
