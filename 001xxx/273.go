func deleteTreeNodes(nodes int, parent []int, value []int) int {
	t := make([]int, nodes)
	for i := nodes - 1; i > 0; i-- {
		p, v := parent[i], value[i]
		value[p] += v
		if v != 0 {
			t[p] += t[i] + 1
		}
	}
	if value[0] != 0 {
		return t[0] + 1
	}
	return 0
}
