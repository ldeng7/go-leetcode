func constructDistancedSequence(n int) []int {
	o := make([]int, n*2-1)
	var cal func(int, int) bool
	cal = func(i, v int) bool {
		if i >= len(o) {
			return true
		} else if o[i] != 0 {
			return cal(i+1, v)
		}
		b := false
		for j := n; j >= 1; j-- {
			m := 1 << (j - 1)
			if v&m != 0 {
				continue
			}
			if j == 1 {
				v |= m
				o[i] = j
				if cal(i+1, v) {
					b = true
					break
				}
				v &^= m
				o[i] = 0
			} else if i+j < 2*n-1 && 0 == o[i+j] {
				v |= m
				o[i], o[i+j] = j, j
				if cal(i+1, v) {
					b = true
					break
				}
				v &^= m
				o[i], o[i+j] = 0, 0
			}
		}
		return b
	}

	cal(0, 0)
	return o
}
