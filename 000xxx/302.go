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

func minArea(image [][]byte, x int, y int) int {
	if 0 == len(image) || 0 == len(image[0]) {
		return 0
	}
	m, n := len(image), len(image[0])
	u, d, l, r := x, x, y, y
	var cal func(int, int)
	cal = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || image[i][j] != '1' {
			return
		}
		image[i][j] = 2
		u, d, l, r = min(u, i), max(d, i), min(l, j), max(r, j)
		cal(i-1, j)
		cal(i, j-1)
		cal(i+1, j)
		cal(i, j+1)
	}
	cal(x, y)
	return (d - u + 1) * (r - l + 1)
}
