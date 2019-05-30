import "sort"

func minTotalDistance(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	rs, cs := []int{}, []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if 1 == grid[i][j] {
				rs = append(rs, i)
				cs = append(cs, j)
			}
		}
	}
	sort.Ints(cs)
	o, i, j := 0, 0, len(rs)-1
	for i < j {
		o += rs[j] - rs[i] + cs[j] - cs[i]
		i++
		j--
	}
	return o
}
