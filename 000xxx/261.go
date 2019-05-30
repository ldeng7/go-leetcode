func validTree(n int, edges [][]int) bool {
	roots := make([]int, n)
	var cal func(i int) int
	cal = func(i int) int {
		for -1 != roots[i] {
			i = roots[i]
		}
		return i
	}

	for i := 0; i < n; i++ {
		roots[i] = -1
	}
	for _, e := range edges {
		a, b := cal(e[0]), cal(e[1])
		if a == b {
			return false
		}
		roots[a] = b
	}
	return len(edges) == n-1
}
