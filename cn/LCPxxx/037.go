import "sort"

func min(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func maxi(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minRecSize(lines [][]int) float64 {
	sort.Slice(lines, func(i, j int) bool {
		x1, y1, x2, y2 := lines[i][0], lines[i][1], lines[j][0], lines[j][1]
		return x1 < x2 || (x1 == x2 && y1 < y2)
	})
	l := len(lines)
	ar := make([][]int, 0, l)
	for i, p := range lines {
		if i == 0 || i == l-1 || p[0] != lines[i-1][0] || p[0] != lines[i+1][0] {
			ar = append(ar, p)
		}
	}

	xmin, ymin, xmax, ymax, c := 1e20, 1e20, -1e20, -1e20, 0
	for i, p := range ar {
		xi, yi := float64(p[0]), float64(p[1])
		for j := maxi(i-3, 0); j < i; j++ {
			xj, yj := float64(ar[j][0]), float64(ar[j][1])
			if xi == xj {
				continue
			}
			x, y := (yi-yj)/(xj-xi), (xj*yi-xi*yj)/(xj-xi)
			xmin, ymin, xmax, ymax, c = min(xmin, x), min(ymin, y), max(xmax, x), max(ymax, y), c+1
		}
	}
	if c == 0 {
		return 0
	}
	return (xmax - xmin) * (ymax - ymin)
}
