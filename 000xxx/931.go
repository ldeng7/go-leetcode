import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minFallingPathSum(A [][]int) int {
	m, n := len(A), len(A[0])
	t := make([][]int, m)
	t[0] = make([]int, n)
	for j := 0; j < m; j++ {
		t[0][j] = A[0][j]
	}
	for i := 1; i < m; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			k := math.MaxInt64
			if j >= 1 {
				k = min(k, t[i-1][j-1])
			}
			if j < n-1 {
				k = min(k, t[i-1][j+1])
			}
			t[i][j] = min(k, t[i-1][j]) + A[i][j]
		}
	}

	o := math.MaxInt64
	ar := t[m-1]
	for j := 0; j < n; j++ {
		o = min(o, ar[j])
	}
	return o
}
