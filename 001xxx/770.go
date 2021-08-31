func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, m+1)
	}
	for i := 0; i < m; i++ {
		mi := multipliers[i]
		t[0][i+1] = t[0][i] + nums[n-i-1]*mi
		for j := 1; j <= i+1; j++ {
			k := i + 1 - j
			if k > 0 {
				t[j][k] = max(t[j-1][k]+nums[j-1]*mi, t[j][k-1]+nums[n-k]*mi)
			} else {
				t[j][k] = t[j-1][k] + nums[j-1]*mi
			}
		}
	}

	o := t[0][m]
	for i := 1; i <= m; i++ {
		o = max(o, t[i][m-i])
	}
	return o
}
