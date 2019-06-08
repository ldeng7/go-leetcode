func peakIndexInMountainArray(A []int) int {
	l, r := 0, len(A)-1
	for l < r {
		m := l + (r-l)>>1
		if A[m] < A[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
