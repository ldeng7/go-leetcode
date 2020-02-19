func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		ar := make([]int, n)
		for k := 0; k < n; k++ {
			ar[k] = 100000
		}
		t[i] = ar
	}
	for _, e := range edges {
		a, b, d := e[0], e[1], e[2]
		t[a][b], t[b][a] = d, d
	}
	for i := 0; i < n; i++ {
		t[i][i] = 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if d := t[j][i] + t[i][k]; t[j][k] > d {
					t[j][k], t[k][j] = d, d
				}
			}
		}
	}
	o, m := -1, 100000
	for i := 0; i < n; i++ {
		c := 0
		for j := 0; j < n; j++ {
			if t[i][j] <= distanceThreshold {
				c++
			}
		}
		if c < m || (c == m && i > o) {
			o, m = i, c
		}
	}
	return o
}
