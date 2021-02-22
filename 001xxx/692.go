func waysToDistribute(n int, k int) int {
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, k)
	}
	for i := 0; i < n; i++ {
		t[i][0] = 1
	}
	for i := 1; i < k; i++ {
		for j := i; j < n; j++ {
			t[j][i] = (t[j-1][i]*(i+1) + t[j-1][i-1]) % 1000000007
		}
	}
	return t[n-1][k-1]
}
