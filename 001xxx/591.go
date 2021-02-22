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

type Stat struct {
	Cnt  int
	RMin int
	RMax int
	CMin int
	CMax int
}

func isPrintable(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	stats := [61]Stat{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s := &(stats[grid[i][j]])
			s.Cnt++
			if s.Cnt == 1 {
				s.RMin, s.RMax, s.CMin, s.CMax = i, i, j, j
			} else {
				s.RMin, s.RMax, s.CMin, s.CMax = min(s.RMin, i), max(s.RMax, i), min(s.CMin, j), max(s.CMax, j)
			}
		}
	}

	var cal func(int) bool
	cal = func(c int) bool {
		s := &(stats[c])
		for i := s.RMin; i <= s.RMax; i++ {
			for j := s.CMin; j <= s.CMax; j++ {
				c1 := grid[i][j]
				if c1 == c {
					grid[i][j] = -1
				} else if c1 == -1 {
					return true
				} else if (c1 != 0) && (cal(c1)) {
					return true
				}
			}
		}
		for i := s.RMin; i <= s.RMax; i++ {
			for j := s.CMin; j <= s.CMax; j++ {
				grid[i][j] = 0
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if c := grid[i][j]; c > 0 && cal(c) {
				return false
			}
		}
	}
	return true
}
