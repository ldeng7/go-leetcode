func findLonelyPixel(picture [][]byte) int {
	if 0 == len(picture) || 0 == len(picture[0]) {
		return 0
	}
	m, n := len(picture), len(picture[0])
	rcs, ccs := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				rcs[i]++
				ccs[j]++
			}
		}
	}
	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && rcs[i] == 1 && ccs[j] == 1 {
				out++
			}
		}
	}
	return out
}
