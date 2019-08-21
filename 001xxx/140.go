func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func stoneGameII(piles []int) int {
	l := len(piles)
	switch l {
	case 0:
		return 0
	case 1:
		return piles[0]
	case 2:
		return piles[0] + piles[1]
	}
	sum, t := [110]int{}, [110][210]int{}
	for i := 0; i < 110; i++ {
		for j := 0; j < 210; j++ {
			t[i][j] = -1
		}
	}
	for i := 1; i <= l; i++ {
		sum[i] = sum[i-1] + piles[i-1]
	}

	var cal func(int, int) int
	cal = func(i, m int) int {
		if t[i][m] != -1 {
			return t[i][m]
		} else if i+m<<1 >= l {
			return sum[l] - sum[i]
		}
		for j := 1; j <= m<<1; j++ {
			t[i][m] = max(t[i][m], sum[l]-sum[i]-cal(i+j, max(m, j)))
		}
		return t[i][m]
	}
	return cal(0, 1)
}
