func countServers(grid [][]int) int {
	o := 0
	for _, node := range grid {
		r, c, i := 0, 0, 0
		for j, b := range node {
			if b == 1 {
				r, i = r+1, j
			}
		}
		if r > 1 {
			o += r
		} else if r == 1 {
			for _, node := range grid {
				c += node[i]
			}
		}
		if c > 1 {
			o++
		}
	}
	return o
}
