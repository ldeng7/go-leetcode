func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func kConcatenationMaxSum(arr []int, k int) int {
	s, s1, m := 0, 0, 0
	for _, a := range arr {
		s += a
		s1 += a
		if s < 0 {
			s = 0
		}
		m = max(s, m)
	}
	m0 := m
	for _, a := range arr {
		s += a
		if s < 0 {
			s = 0
		}
		m = max(s, m)
	}
	m1, m2 := m, 0
	if k > 2 {
		m2 = m1 + (k-2)*s1
	}
	return max(max(m0, m1), m2) % 1000000007
}
