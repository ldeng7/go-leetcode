func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
	}

	cal := func(k int) bool {
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if i < k || j < k {
					continue
				}
				if t[i][j]-t[i-k][j]-t[i][j-k]+t[i-k][j-k] <= threshold {
					return true
				}
			}
		}
		return false
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			t[i][j] = mat[i-1][j-1] + t[i-1][j] + t[i][j-1] - t[i-1][j-1]
		}
	}
	l, r := 0, min(m, n)
	for l <= r {
		if l == r || l == r-1 {
			break
		}
		mid := l + (r-l)>>1
		if cal(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if cal(r) {
		return r
	}
	return l
}
