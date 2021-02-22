func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func connectTwoGroups(cost [][]int) int {
	m, n := len(cost), len(cost[0])
	o, l := 0, max(m, n)
	mii, mij := make([]int, m), make([]int, n)
	for j := 0; j < n; j++ {
		mij[j] = 100
	}
	for i := 0; i < m; i++ {
		mii[i] = 100
		for j := 0; j < n; j++ {
			c := cost[i][j]
			mii[i], mij[j] = min(mii[i], c), min(mij[j], c)
		}
		o += mii[i]
	}
	for j := 0; j < n; j++ {
		o += mij[j]
	}
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			t[i][j] = max(0, mii[i]+mij[j]-cost[i][j])
		}
	}

	ls, lx, ly := make([]int, l), make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		ls[i], lx[i] = -1, -1
		for j := 0; j < l; j++ {
			lx[i] = max(lx[i], t[i][j])
		}
	}

	var vx, vy []bool
	var cal func(int) bool
	cal = func(i int) bool {
		vx[i] = true
		for j := 0; j < l; j++ {
			if !vy[j] && lx[i]+ly[j] == t[i][j] {
				vy[j] = true
				if ls[j] == -1 || cal(ls[j]) {
					ls[j] = i
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < l; i++ {
		for {
			vx, vy = make([]bool, l), make([]bool, l)
			if cal(i) {
				break
			}
			var d int = 1e10
			for j := 0; j < l; j++ {
				if vx[j] {
					for k := 0; k < l; k++ {
						if !vy[k] {
							d = min(d, lx[j]+ly[k]-t[j][k])
						}
					}
				}
			}
			if d == 1e10 {
				return o + 1
			}
			for j := 0; j < l; j++ {
				if vx[j] {
					lx[j] -= d
				}
				if vy[j] {
					ly[j] += d
				}
			}
		}
	}
	for i := 0; i < l; i++ {
		if ls[i] != -1 {
			o -= t[ls[i]][i]
		}
	}
	return o
}
