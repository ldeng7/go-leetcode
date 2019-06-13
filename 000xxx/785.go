func isBipartite(graph [][]int) bool {
	cs := make([]int, len(graph))
	var cal func(int, int) bool
	cal = func(c, i int) bool {
		if cs[i] != 0 {
			return cs[i] == c
		}
		cs[i] = c
		for _, j := range graph[i] {
			if !cal(-c, j) {
				return false
			}
		}
		return true
	}
	for i, _ := range graph {
		if cs[i] == 0 && !cal(1, i) {
			return false
		}
	}
	return true
}
