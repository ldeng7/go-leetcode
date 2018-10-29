func orderOfLargestPlusSign(n int, mines [][]int) int {
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		tt := make([]int, n)
		for i := 0; i < n; i++ {
			tt[i] = n
		}
		t[i] = tt
	}
	for _, m := range mines {
		t[m[0]][m[1]] = 0
	}
	for i := 0; i < n; i++ {
		l, r, u, d := 0, 0, 0, 0
		j, k := 0, n-1
		for j < n {
			if t[i][j] != 0 {
				l++
			} else {
				l = 0
			}
			if t[i][k] != 0 {
				r++
			} else {
				r = 0
			}
			if t[j][i] != 0 {
				u++
			} else {
				u = 0
			}
			if t[k][i] != 0 {
				d++
			} else {
				d = 0
			}
			if l < t[i][j] {
				t[i][j] = l
			}
			if r < t[i][k] {
				t[i][k] = r
			}
			if u < t[j][i] {
				t[j][i] = u
			}
			if d < t[k][i] {
				t[k][i] = d
			}
			j++
			k--
		}
	}
	out := 0
	for i := 0; i < n*n; i++ {
		o := t[i/n][i%n]
		if o > out {
			out = o
		}
	}
	return out
}
