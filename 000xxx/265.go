func minCostII(costs [][]int) int {
	if 0 == len(costs) || 0 == len(costs[0]) {
		return 0
	}
	m1, m2 := -1, -1
	for i := 0; i < len(costs); i++ {
		pm1, pm2 := m1, m2
		m1, m2 = -1, -1
		for j := 0; j < len(costs[i]); j++ {
			if j != pm1 {
				if pm1 >= 0 {
					costs[i][j] += costs[i-1][pm1]
				}
			} else {
				if pm2 >= 0 {
					costs[i][j] += costs[i-1][pm2]
				}
			}
			if m1 < 0 || costs[i][j] < costs[i][m1] {
				m2, m1 = m1, j
			} else if m2 < 0 || costs[i][j] < costs[i][m2] {
				m2 = j
			}
		}
	}
	return costs[len(costs)-1][m1]
}
