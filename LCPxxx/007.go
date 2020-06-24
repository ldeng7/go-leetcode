func numWays(n int, relation [][]int, k int) int {
	t := [10][15]int{}
	t[0][0] = 1
	for i := 0; i < k; i++ {
		for _, r := range relation {
			t[i+1][r[1]] += t[i][r[0]]
		}
	}
	return t[k][n-1]
}
