func C(m, n int) int {
	if n-m < m {
		m = n - m
	}
	o := 1
	for i := n; i > n-m; i-- {
		o *= i
	}
	d := 1
	for i := m; i > 1; i-- {
		d *= i
	}
	return o / d
}

func keyboard(k int, n int) int {
	t := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		t[i] = make([]int, 27)
	}
	for i := 0; i <= 26; i++ {
		t[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= 26; j++ {
			for m := 0; m <= k; m++ {
				if i-m >= 0 {
					t[i][j] += t[i-m][j-1] * C(m, i)
				}
			}
			t[i][j] %= 1000000007
		}
	}
	return t[n][26]
}
