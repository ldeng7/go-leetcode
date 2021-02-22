func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		ar := make([]int, n)
		for j := 0; j < n-1; j++ {
			ar[preferences[i][j]] = n - j
		}
		t[i] = ar
	}

	l := len(pairs)
	ar := make([]bool, n)
	for i := 0; i < l; i++ {
		x, y := pairs[i][0], pairs[i][1]
		for j := i + 1; j < l; j++ {
			u, v := pairs[j][0], pairs[j][1]
			if t[x][u] > t[x][y] && t[u][x] > t[u][v] {
				ar[x], ar[u] = true, true
			}
			if t[x][v] > t[x][y] && t[v][x] > t[v][u] {
				ar[x], ar[v] = true, true
			}
			if t[y][u] > t[y][x] && t[u][y] > t[u][v] {
				ar[y], ar[u] = true, true
			}
			if t[y][v] > t[y][x] && t[v][y] > t[v][u] {
				ar[y], ar[v] = true, true
			}
		}
	}

	o := 0
	for _, b := range ar {
		if b {
			o++
		}
	}
	return o
}
