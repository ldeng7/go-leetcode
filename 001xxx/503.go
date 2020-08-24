func getLastMoment(n int, left []int, right []int) int {
	o := 0
	for _, a := range left {
		if a > o {
			o = a
		}
	}
	if 0 != len(right) {
		m := right[0]
		for i := 1; i < len(right); i++ {
			if a := right[i]; a < m {
				m = a
			}
		}
		if t := n - m; t > o {
			o = t
		}
	}
	return o
}
