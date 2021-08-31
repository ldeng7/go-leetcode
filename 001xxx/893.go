func isCovered(ranges [][]int, left int, right int) bool {
	t := [51]bool{}
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			t[i] = true
		}
	}
	for i := left; i <= right; i++ {
		if !t[i] {
			return false
		}
	}
	return true
}
