func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxVacationDays(flights [][]int, days [][]int) int {
	m, n := len(flights), len(days[0])
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
	}
	out := 0
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			t[j][i] = days[j][i]
			for k := 0; k < m; k++ {
				if (j == k || flights[j][k] != 0) && i < n-1 {
					t[j][i] = max(t[j][i], t[k][i+1]+days[j][i])
				}
				if i == 0 && (j == 0 || flights[0][j] != 0) {
					out = max(out, t[j][0])
				}
			}
		}
	}
	return out
}
