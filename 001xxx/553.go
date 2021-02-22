var m = map[int]int{}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDays(n int) int {
	if n < 2 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	}
	o := 1 + min(n%2+minDays(n/2), n%3+minDays(n/3))
	m[n] = o
	return o
}
