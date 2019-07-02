func shortestBridge(A [][]int) int {
	if 0 == len(A) || 0 == len(A[0]) {
		return 0
	}
	m, n := len(A), len(A[0])
	i, j := 0, 0
ot:
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if A[i][j] == 1 {
				break ot
			}
		}
	}
	a := A[i][j]

	var cal func(int, int) bool
	cal = func(i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n {
			return true
		}
		b := A[i][j]
		if b == a+1 {
			return true
		} else if b == 0 {
			A[i][j] = a + 1
			return true
		} else if b < a {
			return false
		}
		A[i][j] = a + 1
		b1, b2, b3, b4 := cal(i-1, j), cal(i+1, j), cal(i, j-1), cal(i, j+1)
		return b1 && b2 && b3 && b4
	}

	o := 0
	for cal(i, j) {
		a = A[i][j]
		o++
	}
	return o
}
