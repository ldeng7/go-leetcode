func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minDifficulty(jobDifficulty []int, d int) int {
	l := len(jobDifficulty)
	if l < d {
		return -1
	}
	t := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		ar := make([]int, d+1)
		for j := 0; j <= d; j++ {
			ar[j] = 10000
		}
		t[i] = ar
	}
	t[0][1] = 0
	for i := 1; i <= l; i++ {
		t[i][1] = max(t[i-1][1], jobDifficulty[i-1])
	}
	for i := 2; i <= d; i++ {
		for j := i; j <= l; j++ {
			m := jobDifficulty[j-1]
			for k := j - 1; k >= i-1; k-- {
				t[j][i] = min(t[j][i], t[k][i-1]+m)
				m = max(m, jobDifficulty[k-1])
			}
		}
	}
	return t[l][d]
}
