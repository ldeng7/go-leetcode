func numEnclaves(A [][]int) int {
	if 0 == len(A) || 0 == len(A[0]) {
		return 0
	}
	m, n := len(A), len(A[0])
	var cal func(int, int)
	cal = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || A[i][j] == 0 {
			return
		}
		A[i][j] = 0
		cal(i-1, j)
		cal(i, j-1)
		cal(i+1, j)
		cal(i, j+1)
	}
	for j := 0; j < n; j++ {
		cal(0, j)
		cal(m-1, j)
	}
	for i := 1; i < m-1; i++ {
		cal(i, 0)
		cal(i, n-1)
	}
	cnt := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if A[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}
