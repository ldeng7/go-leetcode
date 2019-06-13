func prevPermOpt1(A []int) []int {
	l := len(A)
	i := l - 2
	for ; i >= 0 && A[i] <= A[i+1]; i-- {
	}
	if i == -1 {
		return A
	}
	j := l - 1
	l, r := A[i], A[i+1]
	for ; j > i+1 && (A[j] >= l || A[j] <= r); j-- {
	}
	A[i], A[j] = A[j], A[i]
	return A
}
