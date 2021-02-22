func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func averageWaitingTime(customers [][]int) float64 {
	var s, m float64 = 0, 0
	for _, c := range customers {
		i := float64(c[0])
		m = max(m, i) + float64(c[1])
		s += m - i
	}
	return s / float64(len(customers))
}
