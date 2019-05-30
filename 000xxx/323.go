func countComponents(n int, edges [][]int) int {
	out := n
	roots := make([]int, n)
	for i := 0; i < n; i++ {
		roots[i] = i
	}
	getRoot := func(i int) int {
		for roots[i] != i {
			i = roots[i]
		}
		return i
	}
	for _, e := range edges {
		i1, i2 := getRoot(e[0]), getRoot(e[1])
		if i1 != i2 {
			out--
			roots[i2] = i1
		}
	}
	return out
}
