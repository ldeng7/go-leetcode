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

func splitArray(nums []int, m int) int {
	n := len(nums)
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			t[i][j] = math.MaxInt64
		}
	}
	t[0][0] = 0
	ss := make([]int, n+1)
	for j := 1; j <= n; j++ {
		ss[j] = ss[j-1] + nums[j-1]
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := i - 1; k < j; k++ {
				t[i][j] = min(t[i][j], max(t[i-1][k], ss[j]-ss[k]))
			}
		}
	}
	return t[m][n]
}
