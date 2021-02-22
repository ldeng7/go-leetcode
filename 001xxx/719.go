import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func checkWays(pairs [][]int) int {
	m := 0
	for _, p := range pairs {
		m = max(m, max(p[0], p[1]))
	}
	g := make([][]bool, m+1)
	for i := 1; i <= m; i++ {
		g[i] = make([]bool, m+1)
	}
	c := make([]int, m+1)
	for _, p := range pairs {
		a, b := p[0], p[1]
		g[a][b], g[b][a] = true, true
		c[a]++
		c[b]++
	}

	ar := make([]int, 0, m)
	for i := 1; i <= m; i++ {
		if c[i] > 0 {
			ar = append(ar, i)
		}
	}
	sort.Slice(ar, func(i, j int) bool {
		return c[ar[i]] > c[ar[j]]
	})
	if c[ar[0]] != len(ar)-1 {
		return 0
	}

	t := make([][]bool, m+1)
	for i := 1; i <= m; i++ {
		t[i] = make([]bool, m+1)
	}
	p := make([]int, m+1)
	for i, a := range ar {
		for j := i - 1; j >= 0; j-- {
			if g[a][ar[j]] {
				p[a] = ar[j]
				for k := p[a]; k > 0; k = p[k] {
					t[a][k] = true
				}
				break
			}
		}
	}

	o := 1
	for i := 1; i <= m; i++ {
		for j := i + 1; j <= m; j++ {
			b := g[i][j]
			if b && c[i] == c[j] {
				o = 2
			}
			if b != (t[i][j] || t[j][i]) {
				return 0
			}
		}
	}
	return o
}
