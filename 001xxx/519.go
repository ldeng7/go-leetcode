func countSubTrees(n int, edges [][]int, labels string) []int {
	f, t, o := [26]int{}, make([][]int, n), make([]int, n)
	for i := 0; i < 26; i++ {
		f[i] = 2
	}
	for _, e := range edges {
		a, b := e[0], e[1]
		t[a] = append(t[a], b)
		t[b] = append(t[b], a)
	}

	var cal func(int, int)
	cal = func(v, p int) {
		k := labels[v] - 'a'
		c := f[k]
		f[k]++
		for _, a := range t[v] {
			if a != p {
				cal(a, v)
			}
			o[v] = f[k] - c
		}
	}
	cal(0, -1)
	return o
}
