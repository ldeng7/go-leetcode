func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxSizeSlices(slices []int) int {
	o, l := 0, len(slices)
	n := l / 3
	t := make([][][2]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([][2]int, n+1)
	}
	t[l-1][1][1] = slices[l-1]
	for i := l - 2; i >= 0; i-- {
		for j := 1; j <= n; j++ {
			ar := t[i+1]
			p := &(t[i][j])
			(*p)[0], (*p)[1] = max(ar[j][0], ar[j][1]), ar[j-1][0]+slices[i]
		}
	}
	o = t[0][n][0]

	for i := 0; i < l; i++ {
		t[i] = make([][2]int, n+1)
	}
	for i := l - 2; i >= 0; i-- {
		for j := 1; j <= n; j++ {
			ar := t[i+1]
			p := &(t[i][j])
			(*p)[0], (*p)[1] = max(ar[j][0], ar[j][1]), ar[j-1][0]+slices[i]
		}
	}
	p := t[0][n]
	return max(o, max(p[0], p[1]))
}
