func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDays(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	nt, n1, b, bc, l := 0, 0, false, false, m*n
	t := make([]bool, l)
	ar, arl := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		ar[i], arl[i] = -1, -1
	}

	var cal func(int, int, int)
	cal = func(r, c, p int) {
		i, nc := r*n+c, 0
		t[i], ar[i], arl[i] = true, nt, nt
		nt, n1 = nt+1, n1+1

		cal1 := func(r, c int) {
			if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
				return
			}
			j := r*n + c
			if p == j {
				return
			}
			if t[j] {
				arl[i] = min(arl[i], ar[j])
			} else {
				nc++
				cal(r, c, i)
				arl[i] = min(arl[i], arl[j])
				if arl[j] >= ar[i] && p != -1 {
					b = true
				}
			}
		}

		cal1(r+1, c)
		cal1(r-1, c)
		cal1(r, c+1)
		cal1(r, c-1)
		if p == -1 && nc > 1 {
			b = true
		}
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 || t[r*n+c] {
				continue
			}
			if bc {
				return 0
			}
			bc = true
			cal(r, c, -1)
		}
	}
	if b {
		return 1
	} else if n1 >= 2 {
		return 2
	}
	return 1
}
