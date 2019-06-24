func findMaxForm(strs []string, m int, n int) int {
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
	}
	for _, s := range strs {
		n0, n1 := 0, 0
		for i := 0; i < len(s); i++ {
			if s[i] == '0' {
				n0++
			} else {
				n1++
			}
		}
		for i := m; i >= n0; i-- {
			for j := n; j >= n1; j-- {
				t1 := t[i-n0][j-n1] + 1
				if t1 > t[i][j] {
					t[i][j] = t1
				}
			}
		}
	}
	return t[m][n]
}
