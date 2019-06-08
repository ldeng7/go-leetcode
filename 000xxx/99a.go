func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func mergeStones(stones []int, K int) int {
	l := len(stones)
	if l == 1 {
		return 0
	}
	if l < K || (l-1)%(K-1) != 0 {
		return -1
	}

	t := make([][][]int, l+1)
	for i := 0; i <= l; i++ {
		t[i] = make([][]int, l+1)
		for j := 0; j <= l; j++ {
			t[i][j] = make([]int, l+1)
			for k := 0; k <= K; k++ {
				t[i][j][k] = 100000
			}
		}
	}
	ss := make([]int, l+1)
	for i := 1; i <= l; i++ {
		ss[i] = ss[i-1] + stones[i-1]
		t[i][i][1] = 0
	}

	for i := 1; i < l; i++ {
		for j := 1; j <= l-i; j++ {
			j1 := j + i
			for k := 2; k <= K; k++ {
				for m := j; m < j1; m++ {
					t[j][j1][k] = min(t[j][j1][k], t[j][m][k-1]+t[m+1][j1][1])
				}
			}
			t[j][j1][1] = t[j][j1][K] + ss[j1] - ss[j-1]
		}
	}
	return t[1][l][1]
}
