func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const M, N, INF = 100, 20, 0x7fffffff

var f, g [M + 1][N + 1]int

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	for i := 0; i <= M; i++ {
		for j := 0; j <= N; j++ {
			f[i][j], g[i][j] = INF, INF
		}
	}

	for i := 1; i <= n; i++ {
		f[0][i], g[0][i] = 0, 0
	}
	for i := 1; i <= m; i++ {
		for j := min(i, target); j >= 1; j-- {
			if h := houses[i-1]; h > 0 {
				f[j][h] = min(f[j][h], g[j-1][h])
				for k := 1; k <= n; k++ {
					if k != h {
						f[j][k] = INF
					}
				}
			} else {
				for k := 1; k <= n; k++ {
					f[j][k] = min(f[j][k], g[j-1][k]) + cost[i-1][k-1]
				}
			}

			m1, m2 := INF, INF
			for k := 1; k <= n; k++ {
				if t := f[j][k]; m1 > t {
					m1, m2 = t, m1
				} else if m2 > t {
					m2 = t
				}
			}
			if m1 != m2 {
				for k := 1; k <= n; k++ {
					if f[j][k] == m1 {
						g[j][k] = m2
					} else {
						g[j][k] = m1
					}
				}
			} else {
				for k := 1; k <= n; k++ {
					g[j][k] = m1
				}
			}
			for k := 1; k <= n; k++ {
				f[0][k], g[0][k] = INF, INF
			}
		}
	}

	o := INF
	for i := 1; i <= n; i++ {
		o = min(o, f[target][i])
	}
	if o == INF {
		return -1
	}
	return o
}
