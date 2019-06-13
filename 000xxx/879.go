func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func profitableSchemes(G int, P int, group []int, profit []int) int {
	l := len(group)
	t := make([][][]int, l+1)
	for i := 0; i <= l; i++ {
		t[i] = make([][]int, G+1)
		for j := 0; j <= G; j++ {
			t[i][j] = make([]int, P+1)
		}
	}
	for i := 1; i <= l; i++ {
		g, p := group[i-1], profit[i-1]
		for j := 1; j <= G; j++ {
			for k := 0; k <= P; k++ {
				if k <= p && j >= g {
					t[i][j][k]++
				}
				t[i][j][k] += t[i-1][j][k]
				if j >= g {
					t[i][j][k] += t[i-1][j-g][max(k-p, 0)]
				}
				t[i][j][k] %= 1000000007
			}
		}
	}

	return t[l][G][P]
}
