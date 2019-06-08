func eventualSafeNodes(graph [][]int) []int {
	l := len(graph)
	c := make([]int, l)
	o := []int{}

	var cal func(i int) bool
	cal = func(i int) bool {
		if c[i] > 0 {
			return c[i] == 2
		}
		c[i] = 1
		for _, j := range graph[i] {
			if c[j] == 2 {
				continue
			}
			if c[j] == 1 || !cal(j) {
				return false
			}
		}
		c[i] = 2
		return true
	}

	for i := 0; i < l; i++ {
		if cal(i) {
			o = append(o, i)
		}
	}
	return o
}
