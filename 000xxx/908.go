func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for _, a := range A[1:] {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	d := max - min - (K << 1)
	if d < 0 {
		d = 0
	}
	return d
}
