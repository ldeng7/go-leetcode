func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	q := []int{}
	t := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		t[i] = map[int]bool{}
	}
	for _, e := range edges {
		t[e[0]][e[1]], t[e[1]][e[0]] = true, true
	}
	for i := 0; i < n; i++ {
		if len(t[i]) == 1 {
			q = append(q, i)
		}
	}
	for n > 2 {
		l := len(q)
		n -= l
		for i := 0; i < l; i++ {
			j := q[0]
			q = q[1:]
			for k, _ := range t[j] {
				delete(t[k], j)
				if len(t[k]) == 1 {
					q = append(q, k)
				}
			}
		}
	}
	return q
}
