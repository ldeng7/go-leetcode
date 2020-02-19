func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func minTaps(n int, ranges []int) int {
	l := len(ranges)
	ar := make([][2]int, l)
	for i, r := range ranges {
		ar[i] = [2]int{max(0, i-r), min(l-1, i+r)}
	}
	sort.Slice(ar, func(i, j int) bool {
		a, b := ar[i], ar[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	o, m, e, j := 0, 0, 0, -1
	for i := 0; i < l; i++ {
		p := ar[i]
		if p[0] <= e {
			if p[1] > m {
				m, j = p[1], i
			}
		} else {
			if -1 == j {
				return -1
			}
			o, m, e, i, j = o+1, 0, m, i-1, -1
		}
	}
	if e < l-1 {
		if -1 == j || m < l-1 {
			return -1
		}
		o++
	}
	return o
}
