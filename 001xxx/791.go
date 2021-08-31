func findCenter(edges [][]int) int {
	u, v := edges[0][0], edges[0][1]
	if v == edges[1][0] || v == edges[1][1] {
		return v
	}
	return u
}
