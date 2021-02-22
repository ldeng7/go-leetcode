func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minimumTimeRequired(jobs []int, k int) int {
	l := len(jobs)
	ar := make([]int, l)
	var o, m int = 1e9, 0

	var cal func(int, int)
	cal = func(i, c int) {
		if c >= l {
			o = min(o, m)
			return
		}
		for j := min(i+1, k-1); j >= 0; j-- {
			p := m
			ar[j] += jobs[c]
			m = max(ar[j], m)
			if m < o {
				cal(max(j, i), c+1)
			}
			ar[j] -= jobs[c]
			m = p
		}
	}
	cal(-1, 0)
	return o
}
