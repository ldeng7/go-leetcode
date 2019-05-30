func findBlackPixel(picture [][]byte, N int) int {
	if 0 == len(picture) || 0 == len(picture[0]) {
		return 0
	}
	m, n := len(picture), len(picture[0])
	cs, ms := make([]int, n), map[string]int{}
	for i := 0; i < m; i++ {
		c := 0
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				c++
				cs[j]++
			}
		}
		if c == N {
			ms[string(picture[i])]++
		}
	}
	out := 0
	for s, c := range ms {
		if c != N {
			continue
		}
		for i := 0; i < n; i++ {
			if s[i] == 'B' && cs[i] == N {
				out += N
			}
		}
	}
	return out
}
