func mostSimilar(n int, roads [][]int, names []string, targetPath []string) []int {
	t, t1 := make([]int, n), make([]int, n)
	ar, g := make([][]int, n), make([][]int, n)
	for _, r := range roads {
		a, b := r[0], r[1]
		g[a], g[b] = append(g[a], b), append(g[b], a)
	}

	for _, p := range targetPath {
		for i := 0; i < n; i++ {
			t1[i] = 100
		}
		ar1 := make([][]int, n)
		for i := 0; i < n; i++ {
			for _, j := range g[i] {
				v := t[i]
				if p != names[j] {
					v++
				}
				if t1[j] > v {
					t1[j] = v
					ar1[j] = make([]int, len(ar[i])+1)
					copy(ar1[j], ar[i])
					ar1[j][len(ar[i])] = j
				}
			}
		}
		t, t1, ar = t1, t, ar1
	}

	i, v := 0, t[0]
	for j := 1; j < n; j++ {
		if t[j] < v {
			v, i = t[j], j
		}
	}
	return ar[i]
}
