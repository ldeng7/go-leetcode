func fixedPoint(A []int) int {
	for i, a := range A {
		if i == a {
			return i
		}
	}
	return -1
}
