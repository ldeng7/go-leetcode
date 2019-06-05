import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minWindow(S string, T string) string {
	m, n, b, l := len(S), len(T), -1, math.MaxInt64
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
		t[i][0] = i
		for j := 1; j <= n; j++ {
			t[i][j] = -1
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= min(i, n); j++ {
			if S[i-1] == T[j-1] {
				t[i][j] = t[i-1][j-1]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
		if t[i][n] != -1 {
			l1 := i - t[i][n]
			if l1 < l {
				l, b = l1, t[i][n]
			}
		}
	}
	if b == -1 {
		return ""
	}
	return S[b : b+l]
}
