func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countGoodRectangles(rectangles [][]int) int {
	o, m := 0, 0
	for _, r := range rectangles {
		s := min(r[0], r[1])
		if s > m {
			m, o = s, 1
		} else if s == m {
			o++
		}
	}
	return o
}
