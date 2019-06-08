func allPathsSourceTarget(graph [][]int) [][]int {
	out := [][]int{}
	var cal func(int, []int)
	cal = func(i int, path []int) {
		path = append(path, i)
		if i == len(graph)-1 {
			p := make([]int, len(path))
			copy(p, path)
			out = append(out, p)
		} else {
			for _, j := range graph[i] {
				cal(j, path)
			}
		}
	}
	cal(0, []int{})
	return out
}
