func peakIndexInMountainArray(A []int) int {
	l, h := 0, len(A)-1
	for l < h {
		m := int((uint(l) + uint(h)) >> 1)
		if A[m] < A[m+1] {
			l = m + 1
		} else {
			h = m
		}
	}
	return l
}
