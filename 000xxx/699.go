func fallingSquares(positions [][]int) []int {
	n := len(positions)
	heights := make([]int, n)
	for i, p := range positions {
		l, r := p[0], p[0]+p[1]
		heights[i] += p[1]
		for j := i + 1; j < n; j++ {
			l1, r1 := positions[j][0], positions[j][0]+positions[j][1]
			if l1 < r && r1 > l && heights[i] > heights[j] {
				heights[j] = heights[i]
			}
		}
	}
	out, i, h0 := make([]int, n), 0, 0
	for _, h := range heights {
		if h > h0 {
			h0 = h
		}
		out[i], i = h0, i+1
	}
	return out
}
