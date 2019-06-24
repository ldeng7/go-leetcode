import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func findRotateSteps(ring string, key string) int {
	m, n := len(key), len(ring)
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			t[i][j] = math.MaxInt64
			for k := 0; k < n; k++ {
				if ring[k] == key[i] {
					d := abs(j - k)
					t[i][j] = min(t[i][j], min(d, n-d)+t[i+1][k])
				}
			}
		}
	}
	return t[0][0] + m
}
