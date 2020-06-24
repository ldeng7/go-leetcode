const MOD = 1000000007

var t [51 * 51 * 11]int

func ways(pizza []string, k int) int {
	m, n := len(pizza), len(pizza[0])
	var s [51][51]int
	for i := 0; i < 28611; i++ {
		t[i] = -1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			s[i][j] = s[i+1][j] + s[i][j+1] - s[i+1][j+1]
			if pizza[i][j] == 'A' {
				s[i][j]++
			}
		}
	}

	var cal func(int, int, int) int
	cal = func(i, j, h int) int {
		var o int
		p := &(t[i*51*11+j*11+h])
		if -1 != *p {
			o = *p
		} else if h == k {
			if s[i][j] > 0 {
				o = 1
			}
		} else {
			for r := i; r < m-1; r++ {
				if s[i][j] > s[r+1][j] {
					o = (o + cal(r+1, j, h+1)) % MOD
				}
			}
			for c := j; c < n-1; c++ {
				if s[i][j] > s[i][c+1] {
					o = (o + cal(i, c+1, h+1)) % MOD
				}
			}
		}
		*p = o
		return o
	}

	return cal(0, 0, 1)
}
