import "math"

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

func getMoneyAmount(n int) int {
	t := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		t[i] = make([]int, n+1)
	}
	for i := 2; i <= n; i++ {
		for j := i - 1; j > 0; j-- {
			mi := math.MaxInt64
			for k := j + 1; k < i; k++ {
				ma := k + max(t[j][k-1], t[k+1][i])
				mi = min(mi, ma)
			}
			if j+1 == i {
				t[j][i] = j
			} else {
				t[j][i] = mi
			}
		}
	}
	return t[1][n]
}
