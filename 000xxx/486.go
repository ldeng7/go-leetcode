func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func PredictTheWinner(nums []int) bool {
	l := len(nums)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
		t[i][i] = nums[i]
	}
	for k := 1; k < l; k++ {
		for i, j := 0, k; j < l; i, j = i+1, j+1 {
			t[i][j] = max(nums[i]-t[i+1][j], nums[j]-t[i][j-1])
		}
	}
	return t[0][l-1] >= 0
}
