func strangePrinter(s string) int {
	l := len(s)
	if 0 == l {
		return 0
	}
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
	}
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if i != j {
				t[i][j] = t[i+1][j] + 1
			} else {
				t[i][j] = 1
			}
			for k := i + 1; k <= j; k++ {
				if s[k] == s[i] {
					t1 := t[i+1][k-1] + t[k][j]
					if t1 < t[i][j] {
						t[i][j] = t1
					}
				}
			}
		}
	}
	return t[0][l-1]
}
