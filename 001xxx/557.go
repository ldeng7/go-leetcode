func findSmallestSetOfVertices(n int, edges [][]int) []int {
	o, t := make([]int, 0, n), make([]bool, n)
	for _, e := range edges {
		t[e[1]] = true
	}
	for i, b := range t {
		if !b {
			o = append(o, i)
		}
	}
	return o
}
